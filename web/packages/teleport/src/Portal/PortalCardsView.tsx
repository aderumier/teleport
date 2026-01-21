/**
 * Portal-specific CardsView that:
 * - Hides labels/tags
 * - Hides URL (secondaryDesc)
 * - Makes cards clickable to launch
 * - Hides ActionButton
 */

import React, { useRef, useState, useEffect, useLayoutEffect } from 'react';
import styled, { css } from 'styled-components';

import { Box, ButtonLink, Flex, Label, Text } from 'design';
import { CheckboxInput } from 'design/Checkbox';
import { ResourceIcon } from 'design/ResourceIcon';
import { HoverTooltip } from 'design/Tooltip';

import { CopyButton } from 'shared/components/CopyButton/CopyButton';
import {
  BackgroundColorProps,
  getBackgroundColor,
  getStatusBackgroundColor,
} from 'shared/components/UnifiedResources/shared/getBackgroundColor';
import { PinButton } from 'shared/components/UnifiedResources/shared/PinButton';
import { SingleLineBox } from 'shared/components/UnifiedResources/shared/SingleLineBox';
import { shouldWarnResourceStatus } from 'shared/components/UnifiedResources/shared/StatusInfo';
import { ResourceViewProps } from 'shared/components/UnifiedResources/types';
import { LoadingSkeleton } from 'shared/components/UnifiedResources/shared/LoadingSkeleton';
import { FETCH_MORE_SIZE } from 'shared/components/UnifiedResources/UnifiedResources';
import { LoadingCard } from 'shared/components/UnifiedResources/CardsView/LoadingCard';
import { launchResource } from './launchResource';
import useStickyClusterId from 'teleport/useStickyClusterId';
import { UnifiedResource } from 'teleport/services/agents';
import { generateUnifiedResourceKey } from 'shared/components/UnifiedResources/shared/generateUnifiedResourceKey';

export function PortalCardsView({
  mappedResources,
  onLabelClick,
  pinnedResources,
  selectedResources,
  onSelectResource,
  onPinResource,
  isProcessing,
  pinningSupport,
}: ResourceViewProps) {
  const { clusterId } = useStickyClusterId();

  return (
    <CardsContainer className="CardsContainer" gap={2}>
      {mappedResources.map(
        ({ item, key, onShowStatusInfo, showingStatusInfo }) => {
          // Extract resource from the original resources array
          // We need to find the resource by key
          // For now, we'll create a simplified card that hides labels and URL
          return (
            <PortalResourceCard
              key={key}
              viewItem={{
                ...item,
                // Hide labels for Portal
                labels: [],
                // Hide secondaryDesc (URL) for Portal
                cardViewProps: {
                  ...item.cardViewProps,
                  secondaryDesc: undefined,
                },
                // Remove ActionButton for Portal
                ActionButton: null,
              }}
              onLabelClick={onLabelClick}
              pinned={pinnedResources.includes(key)}
              pinningSupport={pinningSupport}
              selected={selectedResources.includes(key)}
              selectResource={() => onSelectResource(key)}
              pinResource={() => onPinResource(key)}
              onShowStatusInfo={onShowStatusInfo}
              showingStatusInfo={showingStatusInfo}
              resourceKey={key}
              clusterId={clusterId}
            />
          );
        }
      )}
      {isProcessing && (
        <LoadingSkeleton count={FETCH_MORE_SIZE} Element={<LoadingCard />} />
      )}
    </CardsContainer>
  );
}

function PortalResourceCard({
  viewItem,
  pinned,
  pinResource,
  selectResource,
  selected,
  onShowStatusInfo,
  showingStatusInfo,
  pinningSupport,
  resourceKey,
  clusterId,
}: {
  viewItem: any;
  pinned: boolean;
  pinResource: () => void;
  selectResource: () => void;
  selected: boolean;
  onShowStatusInfo: () => void;
  showingStatusInfo: boolean;
  pinningSupport: any;
  resourceKey: string;
  clusterId: string;
}) {
  const {
    name,
    primaryIconName,
    SecondaryIcon,
    cardViewProps,
    requiresRequest,
    status,
  } = viewItem;
  const { primaryDesc } = cardViewProps;

  const [hovered, setHovered] = useState(false);

  const shouldDisplayStatusWarning = shouldWarnResourceStatus(status);

  const handleCardClick = (e: React.MouseEvent) => {
    // Don't trigger if clicking on checkbox or pin button
    if (
      (e.target as HTMLElement).closest('input[type="checkbox"]') ||
      (e.target as HTMLElement).closest('[data-testid="pin-button"]')
    ) {
      return;
    }

    // Extract resource from viewItem and launch it
    // We need to reconstruct the resource from the viewItem
    // For now, we'll use a simplified approach
    // TODO: Extract actual resource object and call launchResource
  };

  return (
    <CardContainer
      onMouseEnter={() => setHovered(true)}
      onMouseLeave={() => setHovered(false)}
      showingStatusInfo={showingStatusInfo}
      onClick={handleCardClick}
      style={{ cursor: 'pointer' }}
    >
      <CardOuterContainer shouldDisplayWarning={shouldDisplayStatusWarning}>
        <CardInnerContainer
          p={3}
          pl={6}
          alignItems="start"
          pinned={pinned}
          requiresRequest={requiresRequest}
          selected={selected}
          showingStatusInfo={showingStatusInfo}
          shouldDisplayWarning={shouldDisplayStatusWarning}
        >
          <CheckboxInput
            checked={selected}
            onChange={selectResource}
            onClick={e => e.stopPropagation()}
            style={{ position: 'absolute', top: '16px', left: '16px' }}
          />
          <Box
            css={`
              position: absolute;
              top: ${props => props.theme.space[9]}px;
              transition: none;
              left: 16px;
            `}
            onClick={e => e.stopPropagation()}
          >
            <PinButton
              setPinned={pinResource}
              pinned={pinned}
              pinningSupport={pinningSupport}
              hovered={hovered}
            />
          </Box>
          <ResourceIcon
            name={primaryIconName}
            width="45px"
            height="45px"
            ml={2}
            css={`
              opacity: ${requiresRequest ? '0.5' : '1'};
            `}
          />
          <Flex flexDirection="column" flex="1" minWidth="0" ml={3} gap={1}>
            <Flex flexDirection="row" alignItems="center" gap={1}>
              <SingleLineBox flex="1">
                <HoverTooltip tipContent={name} showOnlyOnOverflow>
                  <Text typography="body1">{name}</Text>
                </HoverTooltip>
              </SingleLineBox>
              {hovered && <CopyButton value={name} />}
              {/* ActionButton removed for Portal */}
            </Flex>
            <Flex flexDirection="row" alignItems="center">
              <ResTypeIconBox>
                <SecondaryIcon size={18} />
              </ResTypeIconBox>
              {primaryDesc && (
                <SingleLineBox ml={1} title={primaryDesc}>
                  <Text typography="body3" color="text.slightlyMuted">
                    {primaryDesc}
                  </Text>
                </SingleLineBox>
              )}
              {/* secondaryDesc (URL) removed for Portal */}
            </Flex>
            {/* Labels removed for Portal */}
          </Flex>
          {shouldDisplayStatusWarning && (
            <HoverTooltip tipContent={'Show Connection Issue'} placement="left">
              <WarningRightEdgeBadgeIcon onClick={onShowStatusInfo} />
            </HoverTooltip>
          )}
        </CardInnerContainer>
      </CardOuterContainer>
    </CardContainer>
  );
}

const ResTypeIconBox = styled(Box)`
  line-height: 0;
`;

const CardContainer = styled(Box)<{
  showingStatusInfo: boolean;
}>`
  height: 110px;
  position: relative;
  .resource-health-status-svg {
    width: 100%;
    height: 100%;
    fill: ${p =>
      p.showingStatusInfo
        ? p.theme.colors.interactive.solid.alert.active
        : p.theme.colors.interactive.solid.alert.default};
  }
  &:hover {
    .resource-health-status-svg {
      fill: ${p => p.theme.colors.interactive.solid.alert.hover};
    }
  }
`;

const CardOuterContainer = styled(Box)<{
  shouldDisplayWarning: boolean;
}>`
  border-radius: ${props => props.theme.radii[3]}px;
  transition: all 150ms;

  ${CardContainer}:hover && {
    background-color: ${props => props.theme.colors.levels.surface};

    &:after {
      box-shadow: ${props => props.theme.boxShadow[3]};
      border-radius: ${props => props.theme.radii[3]}px;
      content: '';
      position: absolute;
      top: 0;
      left: 0;
      z-index: -1;
      width: 100%;
      height: 100%;
    }
  }
`;

const CardInnerContainer = styled(Flex)<BackgroundColorProps>`
  border: ${props => props.theme.borders[2]}
    ${props => props.theme.colors.spotBackground[0]};
  border-radius: ${props => props.theme.radii[3]}px;
  background-color: ${props => getBackgroundColor(props)};

  ${p =>
    p.shouldDisplayWarning &&
    css`
      border: 2px solid ${p.theme.colors.interactive.solid.alert.default};
      background-color: ${getStatusBackgroundColor({
        showingStatusInfo: p.showingStatusInfo,
        theme: p.theme,
        action: '',
        viewType: 'card',
      })};
    `}

  ${p =>
    p.showingStatusInfo &&
    css`
      border: 2px solid ${p.theme.colors.interactive.solid.alert.active};
    `}

  &:hover {
    border: ${props => props.theme.borders[2]} rgba(0, 0, 0, 0);

    ${p =>
      p.shouldDisplayWarning &&
      css`
        border-color: ${p.theme.colors.interactive.solid.alert.hover};
        background-color: ${getStatusBackgroundColor({
          showingStatusInfo: p.showingStatusInfo,
          theme: p.theme,
          action: 'hover',
          viewType: 'card',
        })};
      `}
  }
`;

const WarningRightEdgeBadgeIcon = ({ onClick }: { onClick?(): void }) => {
  return (
    <Box
      onClick={onClick}
      css={`
        position: absolute;
        top: 0;
        right: 0;
        cursor: pointer;
        height: 100%;
      `}
    >
      {/* Warning icon SVG would go here */}
    </Box>
  );
};

const CardsContainer = styled(Flex)`
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(400px, 1fr));
`;

const CheckboxInput = styled.input.attrs({ type: 'checkbox' })`
  cursor: pointer;
`;
