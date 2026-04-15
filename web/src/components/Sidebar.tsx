import { StatusResponse } from '../lib/api';
import {
  HomeIcon,
  FolderTreeIcon,
  GraphIcon,
  BoxIcon,
  ComponentIcon,
  PlugIcon,
  DiamondIcon,
  GitBranchIcon,
  ScaleIcon,
  ClipboardIcon,
  CheckSquareIcon,
  LightbulbIcon,
  LayersIcon,
  SearchIcon,
} from './icons';

const KIND_GROUPS = [
  {
    label: 'Architecture',
    items: [
      { kind: 'system', label: 'Systems', Icon: BoxIcon },
      { kind: 'component', label: 'Components', Icon: ComponentIcon },
      { kind: 'contract', label: 'Contracts', Icon: PlugIcon },
      { kind: 'concept', label: 'Concepts', Icon: DiamondIcon },
    ],
  },
  {
    label: 'Behavior',
    items: [
      { kind: 'flow', label: 'Flows', Icon: GitBranchIcon },
      { kind: 'decision', label: 'Decisions', Icon: ScaleIcon },
    ],
  },
  {
    label: 'Work',
    items: [
      { kind: 'plan', label: 'Plans', Icon: ClipboardIcon },
      { kind: 'task', label: 'Tasks', Icon: CheckSquareIcon },
      { kind: 'learning', label: 'Learnings', Icon: LightbulbIcon },
    ],
  },
];

interface SidebarProps {
  projectName: string;
  status: StatusResponse | null;
  activeKind: string | null;
  onNavigate: (kind: string | null) => void;
  onSearch: () => void;
}

export function Sidebar({ projectName, status, activeKind, onNavigate, onSearch }: SidebarProps) {
  const counts = status?.counts || {};

  return (
    <aside className="w-56 border-r border-border flex flex-col py-4 px-3 shrink-0 bg-background">
      <div className="flex items-center gap-2 px-3 mb-6">
        <LayersIcon className="w-[18px] h-[18px]" />
        <span className="text-sm font-semibold tracking-tight">syde</span>
      </div>

      <nav className="flex flex-col gap-0.5">
        <button
          onClick={() => onNavigate(null)}
          className={`sidebar-item flex items-center w-full ${activeKind === null ? 'active' : ''}`}
        >
          <HomeIcon className="w-4 h-4 mr-2.5 shrink-0" />
          <span>Overview</span>
        </button>
        <button
          onClick={() => onNavigate('__tree__')}
          className={`sidebar-item flex items-center w-full ${activeKind === '__tree__' ? 'active' : ''}`}
        >
          <FolderTreeIcon className="w-4 h-4 mr-2.5 shrink-0" />
          <span>File Tree</span>
        </button>
        <button
          onClick={() => onNavigate('__graph__')}
          className={`sidebar-item flex items-center w-full ${activeKind === '__graph__' ? 'active' : ''}`}
        >
          <GraphIcon className="w-4 h-4 mr-2.5 shrink-0" />
          <span>Graph</span>
        </button>

        {KIND_GROUPS.map((group) => (
          <div key={group.label} className="mt-3">
            <div className="px-3 text-[10px] font-medium text-muted-foreground uppercase tracking-widest mb-1">
              {group.label}
            </div>
            {group.items.map((item) => {
              const Icon = item.Icon;
              return (
                <button
                  key={item.kind}
                  onClick={() => onNavigate(item.kind)}
                  className={`sidebar-item flex items-center justify-between w-full ${activeKind === item.kind ? 'active' : ''}`}
                >
                  <span className="flex items-center min-w-0">
                    <Icon className="w-4 h-4 mr-2.5 shrink-0" />
                    <span className="truncate">{item.label}</span>
                  </span>
                  {counts[item.kind] ? (
                    <span className="text-[10px] text-muted-foreground bg-muted px-1.5 py-0.5 rounded-full">
                      {counts[item.kind]}
                    </span>
                  ) : null}
                </button>
              );
            })}
          </div>
        ))}
      </nav>

      <div className="mt-auto pt-4 border-t border-border">
        <button
          onClick={onSearch}
          className="sidebar-item w-full flex items-center justify-between text-muted-foreground"
        >
          <span className="flex items-center">
            <SearchIcon className="w-4 h-4 mr-2.5 shrink-0" />
            Search
          </span>
          <kbd className="text-[10px] bg-muted px-1.5 py-0.5 rounded border border-border">⌘K</kbd>
        </button>
      </div>
    </aside>
  );
}
