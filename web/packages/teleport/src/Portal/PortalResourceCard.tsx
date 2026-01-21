/**
 * Portal-specific ResourceCard that:
 * - Hides tags/labels
 * - Hides URL (secondaryDesc)
 * - Makes the card clickable to launch
 * - Hides the ActionButton
 */

import React, { useRef, useState } from 'react';
import styled from 'styled-components';

import { Box, Flex, Text } from 'design';
import { ResourceIcon } from 'design/ResourceIcon';
import { HoverTooltip } from 'design/Tooltip';

import { CopyButton } from 'shared/components/CopyButton/CopyButton';
import {
  BackgroundColorProps,
  getBackgroundColor,
} from 'shared/components/UnifiedResources/shared/getBackgroundColor';
import { PinButton } from 'shared/components/UnifiedResources/shared/PinButton';
import { SingleLineBox } from 'shared/components/UnifiedResources/shared/SingleLineBox';
import { ResourceItemProps } from 'shared/components/UnifiedResources/types';
import { UnifiedResource } from 'teleport/services/agents';
import { ResourceActionButton } from 'teleport/src/UnifiedResources/ResourceActionButton';

// Extract launch action from ResourceActionButton
function getLaunchAction(resource: UnifiedResource): (() => void) | null {
  // This is a simplified version - we'll need to handle different resource types
  // For now, we'll trigger the action button's default behavior
  return null;
}

export function PortalResourceCard({
  onLabelClick,
  pinningSupport,
  pinned,
  pinResource,
  selectResource,
  selected,
  onShowStatusInfo,
  showingStatusInfo,
  viewItem,
  resource,
}: Omit<ResourceItemProps, 'expandAllLabels'> & {
  resource: UnifiedResource;
}) {
  const {
    name,
    primaryIconName,
    SecondaryIcon,
    cardViewProps,
    requiresRequest,
  } = viewItem;
  const { primaryDesc } = cardViewProps;

  const [hovered, setHovered] = useState(false);

  const handleCardClick = (e: React.MouseEvent) => {
    // Don't trigger if clicking on checkbox or pin button
    if (
      (e.target as HTMLElement).closest('input[type="checkbox"]') ||
      (e.target as HTMLElement).closest('[data-testid="pin-button"]')
    ) {
      return;
    }

    // Trigger the launch action
    // We'll need to extract the launch logic from ResourceActionButton
    // For now, we'll create a hidden button and click it
    const actionButton = document.createElement('button');
    actionButton.style.display = 'none';
    document.body.appendChild(actionButton);
    // This is a placeholder - we need to properly extract the launch action
  };

  return (
    <CardContainer
      onMouseEnter={() => setHovered(true)}
      onMouseLeave={() => setHovered(false)}
      showingStatusInfo={showingStatusInfo}
      onClick={handleCardClick}
      style={{ cursor: 'pointer' }}
    >
      <CardOuterContainer>
        <CardInnerContainer
          p={3}
          pl={6}
          alignItems="start"
          pinned={pinned}
          requiresRequest={requiresRequest}
          selected={selected}
          showingStatusInfo={showingStatusInfo}
        >
          <CheckboxInput
            checked={selected}
            onChange={selectResource}
            style={{ position: 'absolute', top: '16px', left: '16px' }}
            onClick={e => e.stopPropagation()}
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
`;

const CardOuterContainer = styled(Box)`
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

  &:hover {
    border: ${props => props.theme.borders[2]} rgba(0, 0, 0, 0);
  }
`;

const CheckboxInput = styled.input.attrs({ type: 'checkbox' })`
  cursor: pointer;
`;
