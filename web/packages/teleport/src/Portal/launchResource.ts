/**
 * Utility function to launch a resource programmatically
 * This extracts the launch logic from ResourceActionButton
 */

import cfg from 'teleport/config';
import { openNewTab } from 'teleport/lib/util';
import { UnifiedResource } from 'teleport/services/agents';
import { App, AppSubKind } from 'teleport/services/apps';
import { Database } from 'teleport/services/databases';
import { Desktop } from 'teleport/services/desktops';
import { Kube } from 'teleport/services/kube';
import { Node } from 'teleport/services/nodes';
import { GitServer } from 'teleport/services/gitServers';
import useStickyClusterId from 'teleport/useStickyClusterId';

export function launchResource(
  resource: UnifiedResource,
  clusterId: string
): void {
  switch (resource.kind) {
    case 'app':
      launchApp(resource as App, clusterId);
      break;
    case 'db':
      // Database requires a dialog, so we can't launch directly
      // We'll need to handle this differently
      break;
    case 'kube_cluster':
      // Kubernetes requires a dialog, so we can't launch directly
      break;
    case 'windows_desktop':
      launchDesktop(resource as Desktop, clusterId);
      break;
    case 'node':
      // Node requires selecting a login, so we can't launch directly
      break;
    case 'git_server':
      // Git server requires a dialog, so we can't launch directly
      break;
    default:
      break;
  }
}

function launchApp(app: App, clusterId: string): void {
  const {
    launchUrl,
    awsConsole,
    awsRoles,
    fqdn,
    publicAddr,
    isCloud,
    isTcp,
    samlApp,
    samlAppSsoUrl,
    samlAppLaunchUrls,
    subKind,
  } = app;

  if (isCloud) {
    // Cloud apps cannot be launched
    return;
  }

  if (isTcp || subKind === AppSubKind.MCP) {
    // TCP and MCP apps require dialogs, so we can't launch directly
    return;
  }

  if (samlApp) {
    // For SAML apps, use the SSO URL or first launch URL
    const url = samlAppLaunchUrls?.[0]?.url || samlAppSsoUrl;
    if (url) {
      openNewTab(url);
    }
    return;
  }

  if (awsConsole || subKind === AppSubKind.AwsIcAccount) {
    // AWS apps require role selection, so we can't launch directly
    return;
  }

  // Simple web app - launch directly
  if (launchUrl) {
    openNewTab(launchUrl);
  }
}

function launchDesktop(desktop: Desktop, clusterId: string): void {
  const url = cfg.getDesktopRoute({
    clusterId,
    desktopName: desktop.name,
    username: desktop.osUser,
  });
  openNewTab(url);
}
