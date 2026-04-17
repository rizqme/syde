// Re-export lucide-react icons under the names the sidebar and other
// components already import. Keeps the API surface unchanged.

import { ComponentType, SVGProps } from 'react';
import {
  Home,
  FolderOpen,
  FolderClosed,
  Network,
  Box,
  LayoutGrid,
  Zap,
  Diamond,
  GitBranch,
  Scale,
  FileCheck,
  ClipboardList,
  CheckSquare,
  Lightbulb,
  Layers,
  Search,
  File,
  ChevronRight,
  ChevronDown,
  Eye,
  EyeOff,
} from 'lucide-react';

type IconProps = SVGProps<SVGSVGElement>;

// Direct re-exports (lucide components accept className, width, height, etc.)
export const HomeIcon = Home;
export const FolderTreeIcon = FolderOpen;
export const FolderClosedIcon = FolderClosed;
export const FolderOpenIcon = FolderOpen;
export const GraphIcon = Network;
export const BoxIcon = Box;
export const ComponentIcon = LayoutGrid;
export const PlugIcon = Zap;
export const DiamondIcon = Diamond;
export const GitBranchIcon = GitBranch;
export const ScaleIcon = Scale;
export const RequirementIcon = FileCheck;
export const ClipboardIcon = ClipboardList;
export const CheckSquareIcon = CheckSquare;
export const LightbulbIcon = Lightbulb;
export const LayersIcon = Layers;
export const SearchIcon = Search;
export const FileIcon = File;
export const ChevronRightIcon = ChevronRight;
export const ChevronDownIcon = ChevronDown;
export const EyeIcon = Eye;
export const EyeOffIcon = EyeOff;

// Map an entity kind to its sidebar icon component.
const KIND_ICON_MAP: Record<string, ComponentType<IconProps>> = {
  system: Box,
  component: LayoutGrid,
  contract: Zap,
  concept: Diamond,
  flow: GitBranch,
  decision: Scale,
  requirement: FileCheck,
  plan: ClipboardList,
  task: CheckSquare,
};

export function iconForKind(kind: string | undefined): ComponentType<IconProps> | null {
  if (!kind) return null;
  return KIND_ICON_MAP[kind] || null;
}
