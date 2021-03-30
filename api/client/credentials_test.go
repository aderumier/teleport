/*
Copyright 2021 Gravitational, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package client

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/gravitational/teleport/api/utils/sshutils"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/ssh"
)

func TestLoadTLS(t *testing.T) {
	t.Parallel()

	// Load expected tls.Config.
	expectedTLSConfig := getExpectedTLSConfig(t)
	// Load TLSConfigCreds.
	creds := LoadTLS(expectedTLSConfig)
	// Build tls.Config and compare to expected tls.Config.
	tlsConfig, err := creds.TLSConfig()
	require.NoError(t, err)
	requireEqualTLSConfig(t, expectedTLSConfig, tlsConfig)

	// Load invalid tls.Config.
	invalidTLSCreds := LoadTLS(nil)
	_, err = invalidTLSCreds.TLSConfig()
	require.Error(t, err)
	_, err = invalidTLSCreds.SSHClientConfig()
	require.Error(t, err)
}

func TestLoadIdentityFile(t *testing.T) {
	t.Parallel()

	// Load expected tls.Config and ssh.ClientConfig.
	expectedTLSConfig := getExpectedTLSConfig(t)
	expectedSSHConfig := getExpectedSSHConfig(t)

	// Write identity file to disk.
	path := filepath.Join(t.TempDir(), "file")
	idFile := &IdentityFile{
		PrivateKey: keyPEM,
		Certs: Certs{
			TLS: tlsCert,
			SSH: sshCert,
		},
		CACerts: CACerts{
			TLS: [][]byte{tlsCACert},
			SSH: [][]byte{sshCACert},
		},
	}
	err := WriteIdentityFile(idFile, path)
	require.NoError(t, err)

	// Load identity file from disk.
	creds := LoadIdentityFile(path)
	// Build tls.Config and compare to expected tls.Config.
	tlsConfig, err := creds.TLSConfig()
	require.NoError(t, err)
	requireEqualTLSConfig(t, expectedTLSConfig, tlsConfig)

	// Build ssh.ClientConfig and compare to expected ssh.ClientConfig.
	sshConfig, err := creds.SSHClientConfig()
	require.NoError(t, err)
	requireEqualSSHConfig(t, expectedSSHConfig, sshConfig)

	// Load invalid identity.
	creds = LoadIdentityFile("invalid_path")
	_, err = creds.TLSConfig()
	require.Error(t, err)
	_, err = creds.SSHClientConfig()
	require.Error(t, err)
}

func TestLoadKeyPair(t *testing.T) {
	t.Parallel()

	// Load expected tls.Config.
	expectedTLSConfig := getExpectedTLSConfig(t)

	// Write key pair and CAs files from bytes.
	path := t.TempDir() + "username"
	certPath, keyPath, caPath := path+".crt", path+".key", path+".cas"
	err := ioutil.WriteFile(certPath, tlsCert, 0600)
	require.NoError(t, err)
	err = ioutil.WriteFile(keyPath, keyPEM, 0600)
	require.NoError(t, err)
	err = ioutil.WriteFile(caPath, tlsCACert, 0600)
	require.NoError(t, err)

	// Load key pair from disk.
	creds := LoadKeyPair(certPath, keyPath, caPath)
	// Build tls.Config and compare to expected tls.Config.
	tlsConfig, err := creds.TLSConfig()
	require.NoError(t, err)
	requireEqualTLSConfig(t, expectedTLSConfig, tlsConfig)

	// Load invalid keypairs.
	invalidIdentityCreds := LoadKeyPair("invalid_path", "invalid_path", "invalid_path")
	_, err = invalidIdentityCreds.TLSConfig()
	require.Error(t, err)
}

func TestLoadProfile(t *testing.T) {
	t.Parallel()

	// Load expected tls.Config and ssh.ClientConfig.
	expectedTLSConfig := getExpectedTLSConfig(t)
	expectedSSHConfig := getExpectedSSHConfig(t)

	// Write identity file to disk.
	dir := t.TempDir()
	name := "proxy.example.com"
	p := &Profile{
		WebProxyAddr: "proxy.example.com:3080",
		SiteName:     "example.com",
		Username:     "testUser",
		Dir:          dir,
	}

	// Save profile and keys to disk.
	require.NoError(t, p.SaveToDir(dir, true))
	require.NoError(t, os.MkdirAll(p.keyDir(), 0700))
	require.NoError(t, os.MkdirAll(p.userKeyDir(), 0700))
	require.NoError(t, os.MkdirAll(p.sshDir(), 0700))
	require.NoError(t, ioutil.WriteFile(p.keyPath(), keyPEM, 0600))
	require.NoError(t, ioutil.WriteFile(p.tlsCertPath(), tlsCert, 0600))
	require.NoError(t, ioutil.WriteFile(p.tlsCasPath(), tlsCACert, 0600))
	require.NoError(t, ioutil.WriteFile(p.sshCertPath(), sshCert, 0600))
	require.NoError(t, ioutil.WriteFile(p.sshCasPath(), sshCACert, 0600))

	// Load profile from disk.
	creds := LoadProfile(dir, name)
	// Build tls.Config and compare to expected tls.Config.
	tlsConfig, err := creds.TLSConfig()
	require.NoError(t, err)
	requireEqualTLSConfig(t, expectedTLSConfig, tlsConfig)
	// Build ssh.ClientConfig and compare to expected ssh.ClientConfig.
	sshConfig, err := creds.SSHClientConfig()
	require.NoError(t, err)
	requireEqualSSHConfig(t, expectedSSHConfig, sshConfig)
	// Build Dialer
	_, err = creds.Dialer(Config{})
	require.NoError(t, err)

	// Load invalid profile.
	creds = LoadProfile("invalid_dir", "invalid_name")
	_, err = creds.TLSConfig()
	require.Error(t, err)
	_, err = creds.SSHClientConfig()
	require.Error(t, err)
	_, err = creds.Dialer(Config{})
	require.Error(t, err)
}

func getExpectedTLSConfig(t *testing.T) *tls.Config {
	cert, err := tls.X509KeyPair(tlsCert, keyPEM)
	require.NoError(t, err)

	pool := x509.NewCertPool()
	require.True(t, pool.AppendCertsFromPEM(tlsCACert))

	return configure(&tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      pool,
	})
}

func getExpectedSSHConfig(t *testing.T) *ssh.ClientConfig {
	config, err := sshutils.ProxyClientSSHConfig(sshCert, keyPEM, [][]byte{sshCACert})
	require.NoError(t, err)

	return config
}

func requireEqualTLSConfig(t *testing.T, expected *tls.Config, actual *tls.Config) {
	require.Empty(t, cmp.Diff(expected, actual,
		cmpopts.IgnoreFields(tls.Config{}, "GetClientCertificate"),
		cmpopts.IgnoreUnexported(tls.Config{}, x509.CertPool{}),
	))
}

func requireEqualSSHConfig(t *testing.T, expected *ssh.ClientConfig, actual *ssh.ClientConfig) {
	require.Empty(t, cmp.Diff(expected, actual,
		cmpopts.IgnoreFields(ssh.ClientConfig{}, "Auth", "HostKeyCallback"),
	))
}

var (
	tlsCert = []byte(`-----BEGIN CERTIFICATE-----
MIIDyzCCArOgAwIBAgIQD3MiJ2Au8PicJpCNFbvcETANBgkqhkiG9w0BAQsFADBe
MRQwEgYDVQQKEwtleGFtcGxlLmNvbTEUMBIGA1UEAxMLZXhhbXBsZS5jb20xMDAu
BgNVBAUTJzIwNTIxNzE3NzMzMTIxNzQ2ODMyNjA5NjAxODEwODc0NTAzMjg1ODAe
Fw0yMTAyMTcyMDI3MjFaFw0yMTAyMTgwODI4MjFaMIGCMRUwEwYDVQQHEwxhY2Nl
c3MtYWRtaW4xCTAHBgNVBAkTADEYMBYGA1UEEQwPeyJsb2dpbnMiOm51bGx9MRUw
EwYDVQQKEwxhY2Nlc3MtYWRtaW4xFTATBgNVBAMTDGFjY2Vzcy1hZG1pbjEWMBQG
BSvODwEHEwtleGFtcGxlLmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoC
ggEBAM5FFaCeK59lwIthyXgSCMZbHTDxsy66Cbm/XhwFbKQLngyS0oKkHbh06INN
UfTAAEaFlMG0CzdAyGyRSu9FK8BE127kRHBs6hb1pTgy2f6TFkFo/h4WTWW4GQSi
O8Al7A2tuRjc3mAnk71q+kvpQYS7tnkhmFCYE8jKxMtlYG39x4kQ6btll7P9zI6X
Zv5RRrlzqADuwZpEcLYVi0TjITqPbx3rDZT4l+EmslhaoG+xE5Vu+GYXLlvwB9E/
amfN1Z9Kps4Ob6Jxxse9kjeMir9mwiNkBWVyhH/LETDA9Xa6sTQ2e75MYM7yXJLY
OmBKV4g176Qf1T1ye7a/Ggn4t2UCAwEAAaNgMF4wDgYDVR0PAQH/BAQDAgWgMB0G
A1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAMBgNVHRMBAf8EAjAAMB8GA1Ud
IwQYMBaAFJWqMooE05nf263F341pOO+mPMSqMA0GCSqGSIb3DQEBCwUAA4IBAQCK
s0yPzkSuCY/LFeHJoJeNJ1SR+EKbk4zoAnD0nbbIsd2quyYIiojshlfehhuZE+8P
bzpUNG2aYKq+8lb0NO+OdZW7kBEDWq7ZwC8OG8oMDrX385fLcicm7GfbGCmZ6286
m1gfG9yqEte7pxv3yWM+7X2bzEjCBds4feahuKPNxOAOSfLUZiTpmOVlRzrpRIhu
2XxiuH+E8n4AP8jf/9bGvKd8PyHohtHVf8HWuKLZxWznQhoKkcfmUmlz5q8ci4Bq
WQdM2NXAMABGAofGrVklPIiraUoHzr0Xxpia4vQwRewYXv8bCPHW+8g8vGBGvoG2
gtLit9DL5DR5ac/CRGJt
-----END CERTIFICATE-----`)

	keyPEM = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAzkUVoJ4rn2XAi2HJeBIIxlsdMPGzLroJub9eHAVspAueDJLS
gqQduHTog01R9MAARoWUwbQLN0DIbJFK70UrwETXbuREcGzqFvWlODLZ/pMWQWj+
HhZNZbgZBKI7wCXsDa25GNzeYCeTvWr6S+lBhLu2eSGYUJgTyMrEy2Vgbf3HiRDp
u2WXs/3Mjpdm/lFGuXOoAO7BmkRwthWLROMhOo9vHesNlPiX4SayWFqgb7ETlW74
ZhcuW/AH0T9qZ83Vn0qmzg5vonHGx72SN4yKv2bCI2QFZXKEf8sRMMD1drqxNDZ7
vkxgzvJcktg6YEpXiDXvpB/VPXJ7tr8aCfi3ZQIDAQABAoIBAE1Vk207wAksAgt/
5yQwRr/vizs9czuSnnDYsbT5x6idfm0iYvB+DXKJyl7oD1Ee5zuJe6NAGHBnxn0F
4D1jBqs4ZDj8NjicbQucn4w5bIfIp7BwZ83p+KypYB/fn11EGoNqXZpXvLv6Oqbq
w9rQIjNcmWZC1TNqQQioFS5Y3NV/gw5uYCRXZlSLMsRCvcX2+LN2EP76ZbkpIVpT
CidC2TxwFPPbyMsG774Olfz4U2IDgX1mO+milF7RIa/vPADSeHAX6tJHmZ13GsyP
0GAdPbFa0Ls/uykeGi1uGPFkdkNEqbWlDf1Z9IG0dr/ck2eh8G2X8E+VFgzsKp4k
WtH9nGECgYEA53lFodLiKQjQR7IoUmGp+P6qnrDwOdU1RfT9jse35xOb9tYvZs3X
kUXU+MEGAMW1Pvmo1v9xOjZbdFYB9I/tIYTSyjYQNaFjgJMPMLSx2qjMzhFXAY5f
8t20/CBt2V1q46aa8tR2ll//QvY4mqvJUaaB0pkuasFbKMXJcGKdvdkCgYEA5CAo
UI8NVA9GqAJfs7hkGHQwpX1X1+JpFhF4dZKsV40NReqaK0vd/mWTYjlMOPO6oolr
PoCDUlQYU6poIDtEnfJ6KkYuLMgxZKnS2OlDthKoZJe9aUTCP1RhTVHyyABRXbGg
tNMKFYkZ38C9+JM+X5T0eKZTHeK+wjiZd55+sm0CgYAmyp0PxI6gP9jf2wyE2dcp
YkxnsdFgb8mwwqDnl7LLJ+8gS76/5Mk2kFRjp72AzaFVP3O7LC3miouDEJLdUG12
C5NjzfGjezt4payLBg00Tsub0S4alaigw+T7x9eA8PXj1tzqyw5gnw/hQfA0g4uG
gngJOiCcRXEogRUEH5K96QKBgFUnB8ViUHhTJ22pTS3Zo0tZe5saWYLVGaLKLKu+
byRTG2RAuQF2VUwTgFtGxgPwPndTUjvHXr2JdHcugaWeWfOXQjCrd6rxozZPCcw7
7jF1b3P1DBfSOavIBHYHI9ex/q05k6JLsFTvkz/pQ0AZPkwRXtv2QcpDDC+VTvvO
pr5VAoGBAJBhNjs9wAu+ZoPcMZcjIXT/BAj2tQYiHoRnNpvQjDYbQueUBeI0Ry8d
5QnKS2k9D278P6BiDBz1c+fS8UErOxY6CS0pi4x3fjMliPwXj/w7AzjlXgDBhRcp
90Ns/9SamlBo9j8ETm9g9D3EVir9zF5XvoR13OdN9gabGy1GuubT
-----END RSA PRIVATE KEY-----`)

	tlsCACert = []byte(`-----BEGIN CERTIFICATE-----
MIIDiTCCAnGgAwIBAgIRAJlp/39yg8U604bjsxgcoC0wDQYJKoZIhvcNAQELBQAw
XjEUMBIGA1UEChMLZXhhbXBsZS5jb20xFDASBgNVBAMTC2V4YW1wbGUuY29tMTAw
LgYDVQQFEycyMDM5MjIyNTY2MzcxMDQ0NDc3MzYxNjA0MTk0NjU2MTgzMDA5NzMw
HhcNMjEwMjAzMDAyOTQ2WhcNMzEwMjAxMDAyOTQ2WjBeMRQwEgYDVQQKEwtleGFt
cGxlLmNvbTEUMBIGA1UEAxMLZXhhbXBsZS5jb20xMDAuBgNVBAUTJzIwMzkyMjI1
NjYzNzEwNDQ0NzczNjE2MDQxOTQ2NTYxODMwMDk3MzCCASIwDQYJKoZIhvcNAQEB
BQADggEPADCCAQoCggEBAKnIJmcKgzj/FbvF6/OYkw3owsS3XU6AcJZ7HmTfYpZF
ozqTDVJdHMFQVfu6cp/6hkzoZ/t7hKT6Nd/O2mlIZdBCfT5ZKESRvTGAeCUANKA5
/D4+6PDdW6AutOFUGbHQ1nYLB7HRgaXF/aZmzFPsPNwX8Wm8EByL+Dws61EmSBBv
Soado5rPG78mAnRpFvyYbzBDkxzsgLIfv0EPw9jhSjrT3OVjCXnBv53u2S+UbJfR
jmI7MutjNbJ/rIBp7JpRHJASmW7oj65WPH0SE0+67XwXYKbs0b7CcSuYW+1S+l9R
uGswW4hqwMloP9sTZoWzgT+nCXQSYUavQF+UJZ/dklMCAwEAAaNCMEAwDgYDVR0P
AQH/BAQDAgKkMA8GA1UdEwEB/wQFMAMBAf8wHQYDVR0OBBYEFC4555Otcq4GAYcc
QQDJgh1TKrFvMA0GCSqGSIb3DQEBCwUAA4IBAQBYsEMJYmSD6Dc1suUEnkWo7kOw
va/aaOu0Phy9SK3hCjg+tatHVVDHO2dZdVCAvCe36BcLiZL1ovFZAXzEzOovwLx/
AVjXpMXTJj52RSMOAtRVSkk3/WOHrGOGIBW2bCKxF4ORXJfWJrdtaObwPPV5sbDC
ACdlNMujdBfUM8EDNmvREI/sVmqL6FK9l6elO/bWLJoiaRTxI+CMixpfIYq8pAwJ
UpgZGjcwco4eqXm7rgbQ4wLaMU6hyk8OE5Glk5E6qpnbVzlrL/jl2iE6EqvI6GJn
Na6B0YR7mdrrL+lyzymnOr6UOrT5nUWRAB1QeY7dhBNnsvoZwaS3VLSc1KCk
-----END CERTIFICATE-----`)

	sshCert = []byte("ssh-rsa-cert-v01@openssh.com AAAAHHNzaC1yc2EtY2VydC12MDFAb3BlbnNzaC5jb20AAAAg8C10PShw+GxCadSlC4nFURIAyvDtgWRvHPabpL5wzDQAAAADAQABAAABAQDORRWgniufZcCLYcl4EgjGWx0w8bMuugm5v14cBWykC54MktKCpB24dOiDTVH0wABGhZTBtAs3QMhskUrvRSvARNdu5ERwbOoW9aU4Mtn+kxZBaP4eFk1luBkEojvAJewNrbkY3N5gJ5O9avpL6UGEu7Z5IZhQmBPIysTLZWBt/ceJEOm7ZZez/cyOl2b+UUa5c6gA7sGaRHC2FYtE4yE6j28d6w2U+JfhJrJYWqBvsROVbvhmFy5b8AfRP2pnzdWfSqbODm+iccbHvZI3jIq/ZsIjZAVlcoR/yxEwwPV2urE0Nnu+TGDO8lyS2DpgSleINe+kH9U9cnu2vxoJ+LdlAAAAAAAAAAAAAAABAAAADGFjY2Vzcy1hZG1pbgAAABAAAAAMYWNjZXNzLWFkbWluAAAAAGAtfCkAAAAAYC4lJQAAAAAAAACdAAAAFnBlcm1pdC1wb3J0LWZvcndhcmRpbmcAAAAAAAAACnBlcm1pdC1wdHkAAAAAAAAADnRlbGVwb3J0LXJvbGVzAAAALQAAACl7InZlcnNpb24iOiJ2MSIsInJvbGVzIjpbImFjY2Vzcy1hZG1pbiJdfQAAAA90ZWxlcG9ydC10cmFpdHMAAAATAAAAD3sibG9naW5zIjpudWxsfQAAAAAAAAEXAAAAB3NzaC1yc2EAAAADAQABAAABAQD3VbuNmR0h3tjYIkTVG+HfNByigp6tuNl8XVylIWx7a7ojRA1nJVzAtNs9QQMut8XY+7jxf4Ue83eIaE0e0QKA0GZlRdbSG0zaYzK8CDAcPVN6Ywt8jnGKuuMhBAckGkN/9nyuJHgTAKeHYgdgQgijPuW/D59s3Sk3vCRHryZzJfZDQ52i40B1q2zLvCcQa6UBvPblHAF3usRa08DnsNkgLey1EkkyvBazqt1amH2Epl3uJRHHUtRVSp2a+0597leT58RZNFfFfB9pccPJfD7cn+iiDmN62T/8YslLYl/O6xCJ43Or7wIRHwJ1tY5hq/Bw7LYn29zeBrIkxIvsH8WtAAABFAAAAAxyc2Etc2hhMi01MTIAAAEAhIz0X+wgA0B8Bi67ALpTEA3kHVWaQY3aT+Ig8obof9upq51H0YlySPJph8h6pVzfSJzQYtuGbmzQ/XAGRMn541mnSUGoy0WCHzscyCowaj9VgjFyVpct7Nz98dB3PnRocNTajGGla+AteZEU3d6KXv/CaA4NGwO3k0rYB+UfX0AAaatAwwxnzYehpCvwSqPdrq/OIyb0aljZHADoNRrcnmYDbB1V76WWY6eTCxYGXx1QyU4A8kH9U8pIZ1fVif/i8dSTbBTftTtv5bmO4WUbVscRw/xIqgZ8v6StNLGHPTt/+Zn+iUoiIrwcnpy+yQp2SRTv7+Lg2SSvJO818x3NNg==")

	sshCACert = []byte("@cert-authority *.example.com ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDMIgxZpT5362npj0x6NQA76IB73bcK85K8cEyKURuHtFC83RjBzvzqtUz6X02+6ohVZiR2MdmsXkCLznzwEIZ0NtoxgnLTZLmduPLeAuYW2vIFpd0G17y6Yog9vxhQ0BLdlhU5Y3JYjRYjmQMfe1iD/RXWD6rEvgWlz+c3HMQR33JqkVIEFH34upfkC2RQG3TXjMe5t14l3yCTtyF5YGzN7+6z/4+/EDto/F3zVtSEp+k8XE/m0ddTGo7usa8ErAom31RwrgkNRmgJmPleDwEflybEsgGKApJXkfFxmG2wu20JoEt/CFjY3fIIa/5aqIGJPpMH4aEdLcj/iyNCog8D type=host")
)
