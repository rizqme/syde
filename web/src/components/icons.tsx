// Flat outline icons (lucide-style, currentColor, 24x24 viewBox).
// Hand-rolled instead of pulling in a dep — only a dozen icons in use.

import { SVGProps } from 'react';

type IconProps = SVGProps<SVGSVGElement>;

function Svg({ children, ...props }: IconProps & { children: React.ReactNode }) {
  return (
    <svg
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      strokeWidth="1.75"
      strokeLinecap="round"
      strokeLinejoin="round"
      aria-hidden="true"
      {...props}
    >
      {children}
    </svg>
  );
}

export function HomeIcon(props: IconProps) {
  return (
    <Svg {...props}>
      <path d="M3 11.5 12 4l9 7.5" />
      <path d="M5 10.5V20h14v-9.5" />
      <path d="M10 20v-5h4v5" />
    </Svg>
  );
}

export function FolderTreeIcon(props: IconProps) {
  return (
    <Svg {...props}>
      <path d="M3 7a2 2 0 0 1 2-2h4l2 2h8a2 2 0 0 1 2 2v9a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V7Z" />
    </Svg>
  );
}

export function GraphIcon(props: IconProps) {
  // 4 nodes connected — visually evokes a network graph.
  return (
    <Svg {...props}>
      <circle cx="6" cy="6" r="2" />
      <circle cx="18" cy="6" r="2" />
      <circle cx="6" cy="18" r="2" />
      <circle cx="18" cy="18" r="2" />
      <circle cx="12" cy="12" r="2" />
      <path d="M7.5 7.5 10.5 10.5" />
      <path d="M16.5 7.5 13.5 10.5" />
      <path d="M7.5 16.5 10.5 13.5" />
      <path d="M16.5 16.5 13.5 13.5" />
    </Svg>
  );
}

export function FolderClosedIcon(props: IconProps) {
  return (
    <Svg {...props}>
      <path d="M3 7a2 2 0 0 1 2-2h4l2 2h8a2 2 0 0 1 2 2v9a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V7Z" />
    </Svg>
  );
}

export function FolderOpenIcon(props: IconProps) {
  return (
    <Svg {...props}>
      <path d="M3 7a2 2 0 0 1 2-2h4l2 2h8a2 2 0 0 1 2 2v1H3V7Z" />
      <path d="M3 9h18l-2.4 8.4a2 2 0 0 1-1.92 1.6H5.32a2 2 0 0 1-1.92-1.6L3 9Z" />
    </Svg>
  );
}

export function FileIcon(props: IconProps) {
  return (
    <Svg {...props}>
      <path d="M14 3H7a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V8l-5-5Z" />
      <path d="M14 3v5h5" />
    </Svg>
  );
}

export function ChevronRightIcon(props: IconProps) {
  return (
    <Svg {...props}>
      <path d="m9 6 6 6-6 6" />
    </Svg>
  );
}

export function ChevronDownIcon(props: IconProps) {
  return (
    <Svg {...props}>
      <path d="m6 9 6 6 6-6" />
    </Svg>
  );
}

export function BoxIcon(props: IconProps) {
  return (
    <Svg {...props}>
      <path d="M21 8 12 3 3 8v8l9 5 9-5V8Z" />
      <path d="m3.3 8.3 8.7 5 8.7-5" />
      <path d="M12 13.3V21" />
    </Svg>
  );
}

export function ComponentIcon(props: IconProps) {
  return (
    <Svg {...props}>
      <rect x="3" y="3" width="7" height="7" rx="1" />
      <rect x="14" y="3" width="7" height="7" rx="1" />
      <rect x="3" y="14" width="7" height="7" rx="1" />
      <rect x="14" y="14" width="7" height="7" rx="1" />
    </Svg>
  );
}

export function PlugIcon(props: IconProps) {
  // Stylised contract / boundary icon — a zap inside a rounded square.
  return (
    <Svg {...props}>
      <rect x="3" y="3" width="18" height="18" rx="3" />
      <path d="M13 7l-4 7h3l-1 4 4-7h-3l1-4Z" />
    </Svg>
  );
}

export function DiamondIcon(props: IconProps) {
  return (
    <Svg {...props}>
      <path d="M12 3 21 12l-9 9-9-9 9-9Z" />
    </Svg>
  );
}

export function GitBranchIcon(props: IconProps) {
  return (
    <Svg {...props}>
      <circle cx="6" cy="5" r="2" />
      <circle cx="6" cy="19" r="2" />
      <circle cx="18" cy="12" r="2" />
      <path d="M6 7v10" />
      <path d="M6 12h6a4 4 0 0 0 4-4" />
    </Svg>
  );
}

export function ScaleIcon(props: IconProps) {
  return (
    <Svg {...props}>
      <path d="M12 4v16" />
      <path d="M6 20h12" />
      <path d="M5 8h14" />
      <path d="M5 8 2 14a3 3 0 0 0 6 0L5 8Z" />
      <path d="M19 8l-3 6a3 3 0 0 0 6 0l-3-6Z" />
    </Svg>
  );
}

export function RequirementIcon(props: IconProps) {
  return (
    <Svg {...props}>
      <path d="M8 5h10a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H8a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2Z" />
      <path d="M10 3v4" />
      <path d="M16 3v4" />
      <path d="M10 12h6" />
      <path d="M10 16h4" />
      <path d="M4 9h4" />
      <path d="M4 13h4" />
      <path d="M4 17h4" />
    </Svg>
  );
}

export function ClipboardIcon(props: IconProps) {
  return (
    <Svg {...props}>
      <rect x="5" y="5" width="14" height="16" rx="2" />
      <rect x="9" y="3" width="6" height="4" rx="1" />
      <path d="M9 12h6" />
      <path d="M9 16h6" />
    </Svg>
  );
}

export function CheckSquareIcon(props: IconProps) {
  return (
    <Svg {...props}>
      <rect x="3" y="3" width="18" height="18" rx="3" />
      <path d="m8 12 3 3 5-6" />
    </Svg>
  );
}

export function LightbulbIcon(props: IconProps) {
  return (
    <Svg {...props}>
      <path d="M9 18h6" />
      <path d="M10 21h4" />
      <path d="M12 3a6 6 0 0 0-4 10.5c.7.8 1 1.6 1 2.5v1h6v-1c0-.9.3-1.7 1-2.5A6 6 0 0 0 12 3Z" />
    </Svg>
  );
}

export function LayersIcon(props: IconProps) {
  return (
    <Svg {...props}>
      <path d="M12 3 2 8l10 5 10-5-10-5Z" />
      <path d="m2 13 10 5 10-5" />
      <path d="m2 18 10 5 10-5" />
    </Svg>
  );
}

export function SearchIcon(props: IconProps) {
  return (
    <Svg {...props}>
      <circle cx="11" cy="11" r="7" />
      <path d="m20 20-3.5-3.5" />
    </Svg>
  );
}

// Map an entity kind to its sidebar icon component. Used by RelationshipChip
// so a chip pointing at a system shows BoxIcon, a chip pointing at a
// contract shows PlugIcon, etc. Same iconography as the sidebar nav.
import { ComponentType } from 'react';
const KIND_ICON_MAP: Record<string, ComponentType<IconProps>> = {
  system: BoxIcon,
  component: ComponentIcon,
  contract: PlugIcon,
  concept: DiamondIcon,
  flow: GitBranchIcon,
  decision: ScaleIcon,
  requirement: RequirementIcon,
  plan: ClipboardIcon,
  task: CheckSquareIcon,
};
export function iconForKind(kind: string | undefined): ComponentType<IconProps> | null {
  if (!kind) return null;
  return KIND_ICON_MAP[kind] || null;
}
