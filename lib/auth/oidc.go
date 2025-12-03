/*
 * Teleport
 * Copyright (C) 2023  Gravitational, Inc.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/gravitational/trace"
	"github.com/zitadel/oidc/v3/pkg/client"
	"github.com/zitadel/oidc/v3/pkg/client/rp"
	"github.com/zitadel/oidc/v3/pkg/oidc"
	"golang.org/x/oauth2"

	"github.com/gravitational/teleport"
	"github.com/gravitational/teleport/api/constants"
	apidefaults "github.com/gravitational/teleport/api/defaults"
	"github.com/gravitational/teleport/api/types"
	apievents "github.com/gravitational/teleport/api/types/events"
	"github.com/gravitational/teleport/lib/auth/authclient"
	"github.com/gravitational/teleport/lib/authz"
	apiutils "github.com/gravitational/teleport/api/utils"
	"github.com/gravitational/teleport/api/utils/keys/hardwarekey"
	"github.com/gravitational/teleport/lib/defaults"
	"github.com/gravitational/teleport/lib/events"
	"github.com/gravitational/teleport/lib/services"
	"github.com/gravitational/teleport/lib/utils"
)

type OIDCService interface {
	CreateOIDCAuthRequest(ctx context.Context, req types.OIDCAuthRequest) (*types.OIDCAuthRequest, error)
	CreateOIDCAuthRequestForMFA(ctx context.Context, req types.OIDCAuthRequest) (*types.OIDCAuthRequest, error)
	ValidateOIDCAuthCallback(ctx context.Context, q url.Values) (*authclient.OIDCAuthResponse, error)
}

var errOIDCNotImplemented = &trace.AccessDeniedError{Message: "OIDC is only available in enterprise subscriptions"}

// UpsertOIDCConnector creates or updates an OIDC connector.
func (a *Server) UpsertOIDCConnector(ctx context.Context, connector types.OIDCConnector) (types.OIDCConnector, error) {
	upserted, err := a.Services.UpsertOIDCConnector(ctx, connector)
	if err != nil {
		return nil, trace.Wrap(err)
	}
	if err := a.emitter.EmitAuditEvent(ctx, &apievents.OIDCConnectorCreate{
		Metadata: apievents.Metadata{
			Type: events.OIDCConnectorCreatedEvent,
			Code: events.OIDCConnectorCreatedCode,
		},
		UserMetadata: authz.ClientUserMetadata(ctx),
		ResourceMetadata: apievents.ResourceMetadata{
			Name: connector.GetName(),
		},
	}); err != nil {
		a.logger.WarnContext(ctx, "Failed to emit OIDC connector create event", "error", err)
	}

	return upserted, nil
}

// UpdateOIDCConnector updates an existing OIDC connector.
func (a *Server) UpdateOIDCConnector(ctx context.Context, connector types.OIDCConnector) (types.OIDCConnector, error) {
	updated, err := a.Services.UpdateOIDCConnector(ctx, connector)
	if err != nil {
		return nil, trace.Wrap(err)
	}
	if err := a.emitter.EmitAuditEvent(ctx, &apievents.OIDCConnectorUpdate{
		Metadata: apievents.Metadata{
			Type: events.OIDCConnectorUpdatedEvent,
			Code: events.OIDCConnectorUpdatedCode,
		},
		UserMetadata: authz.ClientUserMetadata(ctx),
		ResourceMetadata: apievents.ResourceMetadata{
			Name: connector.GetName(),
		},
	}); err != nil {
		a.logger.WarnContext(ctx, "Failed to emit OIDC connector update event", "error", err)
	}

	return updated, nil
}

// CreateOIDCConnector creates a new OIDC connector.
func (a *Server) CreateOIDCConnector(ctx context.Context, connector types.OIDCConnector) (types.OIDCConnector, error) {
	created, err := a.Services.CreateOIDCConnector(ctx, connector)
	if err != nil {
		return nil, trace.Wrap(err)
	}
	if err := a.emitter.EmitAuditEvent(ctx, &apievents.OIDCConnectorCreate{
		Metadata: apievents.Metadata{
			Type: events.OIDCConnectorCreatedEvent,
			Code: events.OIDCConnectorCreatedCode,
		},
		UserMetadata: authz.ClientUserMetadata(ctx),
		ResourceMetadata: apievents.ResourceMetadata{
			Name: connector.GetName(),
		},
	}); err != nil {
		a.logger.WarnContext(ctx, "Failed to emit OIDC connector create event", "error", err)
	}

	return created, nil
}

// DeleteOIDCConnector deletes an OIDC connector by name.
func (a *Server) DeleteOIDCConnector(ctx context.Context, connectorName string) error {
	if err := a.Services.DeleteOIDCConnector(ctx, connectorName); err != nil {
		return trace.Wrap(err)
	}
	if err := a.emitter.EmitAuditEvent(ctx, &apievents.OIDCConnectorDelete{
		Metadata: apievents.Metadata{
			Type: events.OIDCConnectorDeletedEvent,
			Code: events.OIDCConnectorDeletedCode,
		},
		UserMetadata: authz.ClientUserMetadata(ctx),
		ResourceMetadata: apievents.ResourceMetadata{
			Name: connectorName,
		},
	}); err != nil {
		a.logger.WarnContext(ctx, "Failed to emit OIDC connector delete event", "error", err)
	}
	return nil
}

// CreateOIDCAuthRequest delegates the method call to the oidcAuthService if present,
// or returns a NotImplemented error if not present.
func (a *Server) CreateOIDCAuthRequest(ctx context.Context, req types.OIDCAuthRequest) (*types.OIDCAuthRequest, error) {
	if a.oidcAuthService == nil {
		return nil, errOIDCNotImplemented
	}

	rq, err := a.oidcAuthService.CreateOIDCAuthRequest(ctx, req)
	return rq, trace.Wrap(err)
}

// CreateOIDCAuthRequestForMFA delegates the method call to the oidcAuthService if present,
// or returns a NotImplemented error if not present.
func (a *Server) CreateOIDCAuthRequestForMFA(ctx context.Context, req types.OIDCAuthRequest) (*types.OIDCAuthRequest, error) {
	if a.oidcAuthService == nil {
		return nil, errOIDCNotImplemented
	}

	rq, err := a.oidcAuthService.CreateOIDCAuthRequestForMFA(ctx, req)
	return rq, trace.Wrap(err)
}

// ValidateOIDCAuthCallback delegates the method call to the oidcAuthService if present,
// or returns a NotImplemented error if not present.
func (a *Server) ValidateOIDCAuthCallback(ctx context.Context, q url.Values) (*authclient.OIDCAuthResponse, error) {
	if a.oidcAuthService == nil {
		return nil, errOIDCNotImplemented
	}

	resp, err := a.oidcAuthService.ValidateOIDCAuthCallback(ctx, q)
	return resp, trace.Wrap(err)
}

// OIDCAuthService implements OIDC authentication using zitadel/oidc/v3
type OIDCAuthService struct {
	auth         *Server
	emitter      apievents.Emitter
	clients      map[string]*oidcClient
	lock         sync.Mutex
	getClaimsFun func(ctx context.Context, rpClient rp.RelyingParty, connector types.OIDCConnector, code string) (oidc.Claims, error)
}

type OIDCAuthServiceConfig struct {
	Auth    *Server
	Emitter apievents.Emitter
}

func (cfg *OIDCAuthServiceConfig) CheckAndSetDefaults() error {
	if cfg.Auth == nil {
		return trace.BadParameter("auth.Server not provided")
	}
	if cfg.Emitter == nil {
		cfg.Emitter = events.NewDiscardEmitter()
	}
	return nil
}

func NewOIDCAuthService(cfg *OIDCAuthServiceConfig) (*OIDCAuthService, error) {
	if err := cfg.CheckAndSetDefaults(); err != nil {
		return nil, err
	}

	return &OIDCAuthService{
		auth:         cfg.Auth,
		emitter:      cfg.Emitter,
		clients:      make(map[string]*oidcClient),
		getClaimsFun: getClaims,
	}, nil
}

// oidcClient is internal structure that stores OIDC relying party and its config
type oidcClient struct {
	rp        rp.RelyingParty
	connector types.OIDCConnector
	// syncCtx controls the provider sync goroutine.
	syncCtx    context.Context
	syncCancel context.CancelFunc
	// firstSync will be closed once the first provider sync succeeds
	firstSync chan struct{}
}

// ErrOIDCNoRoles results from not mapping any roles from OIDC claims.
var ErrOIDCNoRoles = trace.AccessDenied("No roles mapped from claims. The mappings may contain typos.")

// getOIDCConnectorAndClient returns the associated oidc connector
// and relying party for the given oidc auth request.
func (oas *OIDCAuthService) getOIDCConnectorAndClient(ctx context.Context, request types.OIDCAuthRequest) (types.OIDCConnector, rp.RelyingParty, error) {
	// stateless test flow
	if request.SSOTestFlow {
		if request.ConnectorSpec == nil {
			return nil, nil, trace.BadParameter("ConnectorSpec cannot be nil when SSOTestFlow is true")
		}

		if request.ConnectorID == "" {
			return nil, nil, trace.BadParameter("ConnectorID cannot be empty")
		}

		connector, err := types.NewOIDCConnector(request.ConnectorID, *request.ConnectorSpec)
		if err != nil {
			return nil, nil, trace.Wrap(err)
		}

		// we don't want to cache the client. construct it directly.
		client, err := newOIDCClient(ctx, connector, request.ProxyAddress)
		if err != nil {
			return nil, nil, trace.Wrap(err)
		}
		if err := client.waitFirstSync(defaults.WebHeadersTimeout); err != nil {
			return nil, nil, trace.Wrap(err)
		}

		// close this request-scoped oidc client after 10 minutes
		go func() {
			ticker := oas.auth.GetClock().NewTicker(defaults.OIDCAuthRequestTTL)
			defer ticker.Stop()
			select {
			case <-ticker.Chan():
				client.syncCancel()
			case <-client.syncCtx.Done():
			}
		}()

		return connector, client.rp, nil
	}

	// regular execution flow
	connector, err := oas.auth.GetOIDCConnector(ctx, request.ConnectorID, true)
	if err != nil {
		return nil, nil, trace.Wrap(err)
	}

	client, err := oas.getCachedOIDCClient(ctx, connector, request.ProxyAddress)
	if err != nil {
		return nil, nil, trace.Wrap(err)
	}

	// Wait for the client to successfully sync after getting it from the cache.
	// We do this after caching the client to prevent locking the server during
	// the initial sync period.
	if err := client.waitFirstSync(defaults.WebHeadersTimeout); err != nil {
		return nil, nil, trace.Wrap(err)
	}
	return connector, client.rp, nil
}

// getCachedOIDCClient gets a cached oidc client for
// the given OIDC connector and redirectURL preference.
func (oas *OIDCAuthService) getCachedOIDCClient(ctx context.Context, conn types.OIDCConnector, proxyAddr string) (*oidcClient, error) {
	oas.lock.Lock()
	defer oas.lock.Unlock()

	// Each connector and proxy combination has a distinct client,
	// so we use a composite key to capture all combinations.
	clientMapKey := conn.GetName() + "_" + proxyAddr

	cachedClient, ok := oas.clients[clientMapKey]
	if ok {
		if !cachedClient.needsRefresh(conn) && cachedClient.syncCtx.Err() == nil {
			return cachedClient, nil
		}
		// Cached client needs to be refreshed or is no longer syncing.
		cachedClient.syncCancel()
		delete(oas.clients, clientMapKey)
	}

	// Create a new oidc client and add it to the cache.
	client, err := newOIDCClient(ctx, conn, proxyAddr)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	oas.clients[clientMapKey] = client
	return client, nil
}

func newOIDCClient(ctx context.Context, conn types.OIDCConnector, proxyAddr string) (*oidcClient, error) {
	redirectURL, err := services.GetRedirectURL(conn, proxyAddr)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	scopes := apiutils.Deduplicate(append([]string{"openid", "email"}, conn.GetScope()...))

	// Create relying party using zitadel library
	rpClient, err := rp.NewRelyingPartyOIDC(ctx, conn.GetIssuerURL(), conn.GetClientID(), conn.GetClientSecret(), redirectURL, scopes)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	oidcClient := &oidcClient{rp: rpClient, connector: conn, firstSync: make(chan struct{})}
	oidcClient.startSync(ctx)
	return oidcClient, nil
}

// needsRefresh returns whether the client's connector and the
// given connector have the same values for fields relevant to
// generating and syncing an oidc client.
func (c *oidcClient) needsRefresh(conn types.OIDCConnector) bool {
	return !cmp.Equal(conn.GetRedirectURLs(), c.connector.GetRedirectURLs()) ||
		conn.GetClientID() != c.connector.GetClientID() ||
		conn.GetClientSecret() != c.connector.GetClientSecret() ||
		!cmp.Equal(conn.GetScope(), c.connector.GetScope()) ||
		conn.GetIssuerURL() != c.connector.GetIssuerURL()
}

// startSync starts a goroutine to sync the client with its provider
// config until the given ctx is closed or the sync is canceled.
func (c *oidcClient) startSync(ctx context.Context) {
	c.syncCtx, c.syncCancel = context.WithCancel(ctx)
	go func() {
		// With zitadel, the provider config is fetched automatically when needed
		// We just mark the sync as complete immediately
		close(c.firstSync)
		<-c.syncCtx.Done()
	}()
}

// waitFirstSync waits for the client to start syncing successfully, or
// returns an error if syncing fails or fails to succeed within 10 seconds.
func (c *oidcClient) waitFirstSync(timeout time.Duration) error {
	timeoutTimer := time.NewTimer(timeout)

	select {
	case <-c.firstSync:
	case <-c.syncCtx.Done():
	case <-timeoutTimer.C:
		c.syncCancel()
		return trace.ConnectionProblem(nil, "timed out syncing oidc connector %v, ensure URL %q is valid and accessible and check configuration",
			c.connector.GetName(), c.connector.GetIssuerURL())
	}

	if !timeoutTimer.Stop() {
		<-timeoutTimer.C
	}

	return trace.Wrap(c.syncCtx.Err())
}

func (oas *OIDCAuthService) CreateOIDCAuthRequest(ctx context.Context, req types.OIDCAuthRequest) (*types.OIDCAuthRequest, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	connector, rpClient, err := oas.getOIDCConnectorAndClient(ctx, req)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	stateToken, err := utils.CryptoRandomHex(defaults.TokenLenBytes)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	req.StateToken = stateToken

	// Build authorization URL
	authURL := rp.AuthURL(stateToken, rpClient, rp.WithPrompt(oidc.SplitPrompt(connector.GetPrompt()...))...)
	req.RedirectURL = authURL

	// if the connector has an Authentication Context Class Reference (ACR) value set,
	// update redirect url and add it as a query value.
	acrValue := connector.GetACR()
	if acrValue != "" {
		u, err := url.Parse(req.RedirectURL)
		if err != nil {
			return nil, trace.Wrap(err)
		}
		q := u.Query()
		q.Set("acr_values", acrValue)
		u.RawQuery = q.Encode()
		req.RedirectURL = u.String()
	}

	log.Debugf("OIDC redirect URL: %v.", req.RedirectURL)

	err = oas.auth.Services.CreateOIDCAuthRequest(ctx, req, defaults.OIDCAuthRequestTTL)
	if err != nil {
		return nil, trace.Wrap(err)
	}
	return &req, nil
}

func (oas *OIDCAuthService) CreateOIDCAuthRequestForMFA(ctx context.Context, req types.OIDCAuthRequest) (*types.OIDCAuthRequest, error) {
	return oas.CreateOIDCAuthRequest(ctx, req)
}

// ValidateOIDCAuthCallback is called by the proxy to check OIDC query parameters
// returned by OIDC Provider, if everything checks out, auth server
// will respond with OIDCAuthResponse, otherwise it will return error
func (oas *OIDCAuthService) ValidateOIDCAuthCallback(ctx context.Context, q url.Values) (*authclient.OIDCAuthResponse, error) {
	event := &apievents.UserLogin{
		Metadata: apievents.Metadata{
			Type: events.UserLoginEvent,
		},
		Method: events.LoginMethodOIDC,
	}

	diagCtx := NewSSODiagContext(types.KindOIDC, oas.auth)

	auth, err := oas.validateOIDCAuthCallback(ctx, diagCtx, q)
	diagCtx.Info.Error = trace.UserMessage(err)

	diagCtx.WriteToBackend(ctx)

	claims := diagCtx.Info.OIDCClaims
	if claims != nil {
		attributes, err := apievents.EncodeMap(claims)
		if err != nil {
			event.Status.UserMessage = fmt.Sprintf("Failed to encode identity attributes: %v", err.Error())
			log.WithError(err).Debug("Failed to encode identity attributes.")
		} else {
			event.IdentityAttributes = attributes
		}
	}

	if err != nil {
		event.Code = events.UserSSOLoginFailureCode
		if diagCtx.Info.TestFlow {
			event.Code = events.UserSSOTestFlowLoginFailureCode
		}
		event.Status.Success = false
		event.Status.Error = trace.Unwrap(err).Error()
		event.Status.UserMessage = err.Error()

		if err := oas.emitter.EmitAuditEvent(ctx, event); err != nil {
			log.WithError(err).Warn("Failed to emit OIDC login failed event.")
		}

		return nil, trace.Wrap(err)
	}

	event.Code = events.UserSSOLoginCode
	if diagCtx.Info.TestFlow {
		event.Code = events.UserSSOTestFlowLoginCode
	}
	event.User = auth.Username
	event.Status.Success = true

	if err := oas.emitter.EmitAuditEvent(ctx, event); err != nil {
		log.WithError(err).Warn("Failed to emit OIDC login event.")
	}

	return auth, nil
}

func checkEmailVerifiedClaim(claims oidc.Claims) error {
	claimName := "email_verified"
	unverifiedErr := trace.AccessDenied("email not verified by OIDC provider")

	// Try to get email_verified from claims
	if claimsMap, ok := claims.(map[string]interface{}); ok {
		if emailVerified, ok := claimsMap[claimName]; ok {
			switch v := emailVerified.(type) {
			case string:
				if v == "false" {
					return unverifiedErr
				}
				if v == "true" {
					return nil
				}
				return trace.BadParameter("unable to parse oidc claim: %q, must be either 'true' or 'false', got '%s'", claimName, v)
			case bool:
				if !v {
					return unverifiedErr
				}
				return nil
			default:
				return trace.BadParameter("unable to parse oidc claim: %q, must be a string or bool", claimName)
			}
		}
	}

	return nil
}

func (oas *OIDCAuthService) validateOIDCAuthCallback(ctx context.Context, diagCtx *SSODiagContext, q url.Values) (*authclient.OIDCAuthResponse, error) {
	if errParam := q.Get("error"); errParam != "" {
		state := q.Get("state")
		if state != "" {
			diagCtx.RequestID = state
			req, err := oas.auth.GetOIDCAuthRequest(ctx, state)
			if err == nil {
				diagCtx.Info.TestFlow = req.SSOTestFlow
			}
		}

		errDesc := q.Get("error_description")
		oauthErr := trace.OAuth2("invalid_request", errParam, q)
		return nil, trace.WithUserMessage(oauthErr, "OIDC provider returned error: %v [%v]", errDesc, errParam)
	}

	code := q.Get("code")
	if code == "" {
		oauthErr := trace.OAuth2("invalid_request", "code query param must be set", q)
		return nil, trace.WithUserMessage(oauthErr, "Invalid parameters received from OIDC provider.")
	}

	stateToken := q.Get("state")
	if stateToken == "" {
		oauthErr := trace.OAuth2("invalid_request", "missing state query param", q)
		return nil, trace.WithUserMessage(oauthErr, "Invalid parameters received from OIDC provider.")
	}
	diagCtx.RequestID = stateToken

	req, err := oas.auth.GetOIDCAuthRequest(ctx, stateToken)
	if err != nil {
		return nil, trace.Wrap(err, "Failed to get OIDC Auth Request.")
	}
	diagCtx.Info.TestFlow = req.SSOTestFlow

	ctxC, cancel := context.WithCancel(ctx)
	defer cancel()

	connector, rpClient, err := oas.getOIDCConnectorAndClient(ctxC, *req)
	if err != nil {
		return nil, trace.Wrap(err, "Failed to get OIDC connector and client.")
	}

	// extract claims from both the id token and the userinfo endpoint and merge them
	claims, err := oas.getClaims(ctx, rpClient, connector, code)
	if err != nil {
		return nil, trace.Wrap(err, "Failed to extract OIDC claims. This may indicate need to set 'provider' flag in connector definition.")
	}
	diagCtx.Info.OIDCClaims = types.OIDCClaims(claimsToMap(claims))

	log.Debugf("OIDC claims: %v.", claims)
	if !connector.GetAllowUnverifiedEmail() {
		if err := checkEmailVerifiedClaim(claims); err != nil {
			return nil, trace.Wrap(err, "OIDC provider did not verify email.")
		}
	}

	acrValue := connector.GetACR()
	if acrValue != "" {
		err := validateACRValues(acrValue, connector.GetProvider(), claims)
		if err != nil {
			return nil, trace.Wrap(err, "OIDC ACR validation failure.")
		}
		log.Debugf("OIDC ACR values %q successfully validated.", acrValue)
	}

	// Extract identity from claims
	ident, err := extractIdentityFromClaims(claims)
	if err != nil {
		return nil, trace.OAuth2("unsupported_response_type", "unable to convert claims to identity", q)
	}
	diagCtx.Info.OIDCIdentity = &types.OIDCIdentity{
		ID:        ident.ID,
		Name:      ident.Name,
		Email:     ident.Email,
		ExpiresAt: ident.ExpiresAt,
	}
	log.Debugf("OIDC user %q expires at: %v.", ident.Email, ident.ExpiresAt)

	if len(connector.GetClaimsToRoles()) == 0 {
		oauthErr := trace.BadParameter("no claims to roles mapping, check connector documentation")
		return nil, trace.WithUserMessage(oauthErr, "Claims-to-roles mapping is empty, SSO user will never have any roles.")
	}
	log.Debugf("Applying %v OIDC claims to roles mappings.", len(connector.GetClaimsToRoles()))
	diagCtx.Info.OIDCClaimsToRoles = connector.GetClaimsToRoles()

	params, err := oas.calculateOIDCUser(diagCtx, connector, claims, ident, req)
	if err != nil {
		return nil, trace.Wrap(err, "Failed to calculate user attributes.")
	}

	diagCtx.Info.CreateUserParams = &types.CreateUserParams{
		ConnectorName: params.ConnectorName,
		Username:      params.Username,
		KubeGroups:    params.KubeGroups,
		KubeUsers:     params.KubeUsers,
		Roles:         params.Roles,
		Traits:        params.Traits,
		SessionTTL:    types.Duration(params.SessionTTL),
	}

	user, err := oas.createOIDCUser(ctx, params, req.SSOTestFlow)
	if err != nil {
		return nil, trace.Wrap(err, "Failed to create user from provided parameters.")
	}

	resp := &authclient.OIDCAuthResponse{
		Req: OIDCAuthRequestFromProto(req),
		Identity: types.ExternalIdentity{
			ConnectorID: params.ConnectorName,
			Username:    params.Username,
		},
		Username: user.GetName(),
	}

	if req.SSOTestFlow {
		diagCtx.Info.Success = true
		return resp, nil
	}

	if !req.CheckUser {
		return resp, nil
	}

	if req.CreateWebSession {
		session, err := oas.auth.CreateWebSessionFromReq(ctx, NewWebSessionRequest{
			User:             user.GetName(),
			Roles:            user.GetRoles(),
			Traits:           user.GetTraits(),
			SessionTTL:       params.SessionTTL,
			LoginTime:        oas.auth.GetClock().Now().UTC(),
			LoginIP:          req.ClientLoginIP,
			LoginUserAgent:   req.ClientUserAgent,
			AttestWebSession: true,
		})
		if err != nil {
			return nil, trace.Wrap(err, "Failed to create web session.")
		}
		resp.Session = session
	}

	sshPublicKey, tlsPublicKey, err := authclient.UserPublicKeys(
		req.PublicKey,
		req.SshPublicKey,
		req.TlsPublicKey,
	)
	if err != nil {
		return nil, trace.Wrap(err)
	}
	sshAttestationStatement, tlsAttestationStatement := authclient.UserAttestationStatements(
		hardwarekey.AttestationStatementFromProto(req.AttestationStatement),
		hardwarekey.AttestationStatementFromProto(req.SshAttestationStatement),
		hardwarekey.AttestationStatementFromProto(req.TlsAttestationStatement),
	)
	if len(sshPublicKey)+len(tlsPublicKey) > 0 {
		sshCert, tlsCert, err := oas.auth.CreateSessionCerts(ctx, &SessionCertsRequest{
			UserState:               user,
			SessionTTL:              params.SessionTTL,
			SSHPubKey:               sshPublicKey,
			TLSPubKey:               tlsPublicKey,
			SSHAttestationStatement: sshAttestationStatement,
			TLSAttestationStatement: tlsAttestationStatement,
			Compatibility:           req.Compatibility,
			RouteToCluster:          req.RouteToCluster,
			KubernetesCluster:       req.KubernetesCluster,
			LoginIP:                 req.ClientLoginIP,
		})
		if err != nil {
			return nil, trace.Wrap(err, "Failed to create session certificate.")
		}

		clusterName, err := oas.auth.GetClusterName()
		if err != nil {
			return nil, trace.Wrap(err, "Failed to obtain cluster name.")
		}

		resp.Cert = sshCert
		resp.TLSCert = tlsCert

		authority, err := oas.auth.GetCertAuthority(ctx, types.CertAuthID{
			Type:       types.HostCA,
			DomainName: clusterName.GetClusterName(),
		}, false)
		if err != nil {
			return nil, trace.Wrap(err, "Failed to obtain cluster's host CA.")
		}
		resp.HostSigners = append(resp.HostSigners, authority)
	}

	return resp, nil
}

// OIDCAuthRequestFromProto converts the types.OIDCAuthRequest to OIDCAuthRequest.
func OIDCAuthRequestFromProto(req *types.OIDCAuthRequest) authclient.OIDCAuthRequest {
	return authclient.OIDCAuthRequest{
		ConnectorID:       req.ConnectorID,
		PublicKey:         req.PublicKey,
		SSHPubKey:         req.SshPublicKey,
		TLSPubKey:         req.TlsPublicKey,
		CSRFToken:         req.CSRFToken,
		CreateWebSession:  req.CreateWebSession,
		ClientRedirectURL: req.ClientRedirectURL,
	}
}

type oidcIdentity struct {
	ID        string
	Name      string
	Email     string
	ExpiresAt time.Time
}

func extractIdentityFromClaims(claims oidc.Claims) (*oidcIdentity, error) {
	claimsMap := claimsToMap(claims)

	ident := &oidcIdentity{}

	if sub, ok := claimsMap["sub"].(string); ok {
		ident.ID = sub
	}

	if email, ok := claimsMap["email"].(string); ok {
		ident.Email = email
	}

	if name, ok := claimsMap["name"].(string); ok {
		ident.Name = name
	}

	if exp, ok := claimsMap["exp"].(float64); ok {
		ident.ExpiresAt = time.Unix(int64(exp), 0)
	}

	if ident.Email == "" {
		return nil, trace.BadParameter("email claim not found")
	}

	return ident, nil
}

func claimsToMap(claims oidc.Claims) map[string]interface{} {
	if claimsMap, ok := claims.(map[string]interface{}); ok {
		return claimsMap
	}
	// Try to marshal and unmarshal if it's a struct
	data, err := json.Marshal(claims)
	if err != nil {
		return make(map[string]interface{})
	}
	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return make(map[string]interface{})
	}
	return result
}

func (oas *OIDCAuthService) calculateOIDCUser(diagCtx *SSODiagContext, connector types.OIDCConnector, claims oidc.Claims, ident *oidcIdentity, request *types.OIDCAuthRequest) (*CreateUserParams, error) {
	claimsMap := claimsToMap(claims)

	username, err := usernameFromClaims(connector, claimsMap, ident)
	if err != nil {
		return nil, err
	}

	p := CreateUserParams{
		ConnectorName: connector.GetName(),
		Username:      username,
	}

	p.Traits = services.OIDCClaimsToTraits(claims)

	diagCtx.Info.OIDCTraitsFromClaims = p.Traits
	diagCtx.Info.OIDCConnectorTraitMapping = connector.GetTraitMappings()

	var warnings []string
	warnings, p.Roles = services.TraitsToRoles(connector.GetTraitMappings(), p.Traits)
	if len(p.Roles) == 0 {
		if len(warnings) != 0 {
			log.WithField("connector", connector).Warnf("No roles mapped from claims. Warnings: %q", warnings)
			diagCtx.Info.OIDCClaimsToRolesWarnings = &types.SSOWarnings{
				Message:  "No roles mapped for the user",
				Warnings: warnings,
			}
		} else {
			log.WithField("connector", connector).Warnf("No roles mapped from claims.")
			diagCtx.Info.OIDCClaimsToRolesWarnings = &types.SSOWarnings{
				Message: "No roles mapped for the user. The mappings may contain typos.",
			}
		}
		return nil, trace.Wrap(ErrOIDCNoRoles)
	}

	roles, err := services.FetchRoles(p.Roles, oas.auth, p.Traits)
	if err != nil {
		return nil, trace.Wrap(err)
	}
	roleTTL := roles.AdjustSessionTTL(apidefaults.MaxCertDuration)
	p.SessionTTL = utils.MinTTL(roleTTL, request.CertTTL)

	return &p, nil
}

func (oas *OIDCAuthService) createOIDCUser(ctx context.Context, p *CreateUserParams, dryRun bool) (types.User, error) {
	expires := oas.auth.GetClock().Now().UTC().Add(p.SessionTTL)

	log.Debugf("Generating dynamic OIDC identity %v/%v with roles: %v. Dry run: %v.", p.ConnectorName, p.Username, p.Roles, dryRun)
	user := &types.UserV2{
		Kind:    types.KindUser,
		Version: types.V2,
		Metadata: types.Metadata{
			Name:      p.Username,
			Namespace: apidefaults.Namespace,
			Expires:   &expires,
		},
		Spec: types.UserSpecV2{
			Roles:  p.Roles,
			Traits: p.Traits,
			OIDCIdentities: []types.ExternalIdentity{
				{
					ConnectorID: p.ConnectorName,
					Username:    p.Username,
				},
			},
			CreatedBy: types.CreatedBy{
				User: types.UserRef{Name: teleport.UserSystem},
				Time: oas.auth.GetClock().Now().UTC(),
				Connector: &types.ConnectorRef{
					Type:     constants.OIDC,
					ID:       p.ConnectorName,
					Identity: p.Username,
				},
			},
		},
	}

	if dryRun {
		return user, nil
	}

	existingUser, err := oas.auth.Services.GetUser(ctx, p.Username, false)
	if err != nil && !trace.IsNotFound(err) {
		return nil, trace.Wrap(err)
	}

	ctxtodo := context.TODO()

	if existingUser != nil {
		connectorRef := existingUser.GetCreatedBy().Connector

		if connectorRef == nil {
			return nil, trace.AlreadyExists("local user with name %q already exists. Either change "+
				"email in OIDC identity or remove local user and try again.", existingUser.GetName())
		}

		log.Debugf("Overwriting existing user %q created with %v connector %v.",
			existingUser.GetName(), connectorRef.Type, connectorRef.ID)

		if _, err := oas.auth.UpsertUser(ctxtodo, user); err != nil {
			return nil, trace.Wrap(err)
		}
	} else {
		if _, err := oas.auth.CreateUser(ctxtodo, user); err != nil {
			return nil, trace.Wrap(err)
		}
	}

	return user, nil
}

func usernameFromClaims(connector types.OIDCConnector, claims map[string]interface{}, ident *oidcIdentity) (string, error) {
	usernameClaim := connector.GetUsernameClaim()
	if usernameClaim == "" {
		return ident.Email, nil
	}

	if username, ok := claims[usernameClaim].(string); ok {
		return username, nil
	}

	return "", trace.BadParameter("The configured username_claim of %q was not received from the IdP. Please update the username_claim in connector %q.", usernameClaim, connector.GetName())
}

func (oas *OIDCAuthService) getClaims(ctx context.Context, rpClient rp.RelyingParty, connector types.OIDCConnector, code string) (oidc.Claims, error) {
	return oas.getClaimsFun(ctx, rpClient, connector, code)
}

func getClaims(ctx context.Context, rpClient rp.RelyingParty, connector types.OIDCConnector, code string) (oidc.Claims, error) {
	// Exchange authorization code for tokens
	tokens, err := rp.CodeExchange(ctx, code, rpClient)
	if err != nil {
		return nil, trace.Wrap(err, "failed to exchange authorization code")
	}

	// Verify and extract ID token claims
	idTokenClaims, err := rp.VerifyIDToken[*oidc.IDTokenClaims](ctx, tokens.IDToken(), rpClient)
	if err != nil {
		return nil, trace.Wrap(err, "unable to verify ID token")
	}

	log.Debugf("OIDC ID Token claims: %v.", idTokenClaims)

	// Try to get UserInfo claims
	userInfoClaims, err := getUserInfoClaims(ctx, rpClient, tokens.AccessToken())
	if err != nil {
		if trace.IsNotFound(err) || trace.IsAccessDenied(err) {
			log.Debugf("OIDC provider doesn't offer valid UserInfo endpoint. Returning token claims: %v.", idTokenClaims)
			return idTokenClaims, nil
		}
		return nil, trace.Wrap(err, "unable to fetch UserInfo claims")
	}
	log.Debugf("UserInfo claims: %v.", userInfoClaims)

	// Verify subject matches
	if idTokenClaims.Subject != userInfoClaims.Subject {
		return nil, trace.BadParameter("OIDC claim subjects in UserInfo does not match")
	}

	// Merge claims
	mergedClaims := mergeOIDCClaims(idTokenClaims, userInfoClaims)
	return mergedClaims, nil
}

func getUserInfoClaims(ctx context.Context, rpClient rp.RelyingParty, accessToken string) (oidc.UserInfo, error) {
	issuerURL := rpClient.OIDCConfig().Issuer

	err := isHTTPS(issuerURL)
	if err != nil {
		return oidc.UserInfo{}, trace.NotFound(err.Error())
	}

	// Get UserInfo endpoint from provider config
	dc, err := client.Discover(ctx, issuerURL, http.DefaultClient)
	if err != nil {
		return oidc.UserInfo{}, trace.Wrap(err)
	}

	if dc.UserinfoEndpoint == nil {
		return oidc.UserInfo{}, trace.NotFound("UserInfo endpoint not found")
	}

	endpoint := dc.UserinfoEndpoint.String()

	err = isHTTPS(endpoint)
	if err != nil {
		return oidc.UserInfo{}, trace.NotFound(err.Error())
	}

	log.Debugf("Fetching OIDC claims from UserInfo endpoint: %q.", endpoint)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return oidc.UserInfo{}, trace.Wrap(err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return oidc.UserInfo{}, trace.Wrap(err)
	}
	defer resp.Body.Close()

	code := resp.StatusCode
	if code < 200 || code > 299 {
		if code == http.StatusBadRequest || code == http.StatusUnauthorized ||
			code == http.StatusForbidden || code == http.StatusMethodNotAllowed {
			return oidc.UserInfo{}, trace.AccessDenied("bad status code: %v", code)
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return oidc.UserInfo{}, trace.Wrap(err)
		}
		return oidc.UserInfo{}, trace.ReadError(code, body)
	}

	var userInfo oidc.UserInfo
	err = json.NewDecoder(resp.Body).Decode(&userInfo)
	if err != nil {
		return oidc.UserInfo{}, trace.Wrap(err)
	}

	return userInfo, nil
}

func mergeOIDCClaims(idToken *oidc.IDTokenClaims, userInfo oidc.UserInfo) oidc.Claims {
	// Convert both to maps and merge
	idMap := claimsToMap(idToken)
	uiMap := claimsToMap(userInfo)

	for k, v := range uiMap {
		if _, exists := idMap[k]; !exists {
			idMap[k] = v
		}
	}

	return idMap
}

func validateACRValues(acrValue string, identityProvider string, claims oidc.Claims) error {
	claimsMap := claimsToMap(claims)

	switch identityProvider {
	case teleport.NetIQ:
		log.Debugf("Validating OIDC ACR values with '%v' rules.", identityProvider)

		tokenAcr, ok := claimsMap["acr"]
		if !ok {
			return trace.BadParameter("acr not found in claims")
		}
		tokenAcrMap, ok := tokenAcr.(map[string]interface{})
		if !ok {
			return trace.BadParameter("acr unexpected type: %T", tokenAcr)
		}
		tokenAcrValues, ok := tokenAcrMap["values"]
		if !ok {
			return trace.BadParameter("acr.values not found in claims")
		}
		tokenAcrValuesSlice, ok := tokenAcrValues.([]interface{})
		if !ok {
			return trace.BadParameter("acr.values unexpected type: %T", tokenAcr)
		}

		acrValueMatched := false
		for _, v := range tokenAcrValuesSlice {
			vv, ok := v.(string)
			if !ok {
				continue
			}
			if acrValue == vv {
				acrValueMatched = true
				break
			}
		}
		if !acrValueMatched {
			log.Debugf("No OIDC ACR match found for '%v' in '%v'.", acrValue, tokenAcrValues)
			return trace.BadParameter("acr claim does not match")
		}
	default:
		log.Debugf("Validating OIDC ACR values with default rules.")

		claimValue, ok := claimsMap["acr"].(string)
		if !ok {
			return trace.BadParameter("acr claim does not exist")
		}
		if claimValue != acrValue {
			log.Debugf("No OIDC ACR match found '%v' != '%v'.", acrValue, claimValue)
			return trace.BadParameter("acr claim does not match")
		}
	}

	return nil
}

func isHTTPS(u string) error {
	earl, err := url.Parse(u)
	if err != nil {
		return trace.Wrap(err)
	}
	if earl.Scheme != "https" {
		return trace.BadParameter("expected scheme https, got %q", earl.Scheme)
	}
	return nil
}
