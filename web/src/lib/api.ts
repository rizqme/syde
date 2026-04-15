const BASE = '/api';

let projectSlug: string | null = null;

export function setProject(slug: string) {
  projectSlug = slug;
}

export function getProject() {
  return projectSlug;
}

async function fetchJSON<T>(path: string): Promise<T> {
  const res = await fetch(`${BASE}/${projectSlug}/${path}`);
  if (!res.ok) throw new Error(`API error: ${res.status}`);
  return res.json();
}

export async function fetchProjects(): Promise<{ projects: Project[] }> {
  const res = await fetch(`${BASE}/projects`);
  return res.json();
}

export const api = {
  status: () => fetchJSON<StatusResponse>('status'),
  entities: (kind?: string) => fetchJSON<EntitiesResponse>(kind ? `entities?kind=${kind}` : 'entities'),
  entity: (slug: string) => fetchJSON<EntityDetailResponse>(`entity/${slug}`),
  graph: () => fetchJSON<GraphResponse>('graph'),
  plans: () => fetchJSON<PlansResponse>('plans'),
  learnings: () => fetchJSON<LearningsResponse>('learnings'),
  tasks: () => fetchJSON<TasksResponse>('tasks'),
  designs: () => fetchJSON<DesignsResponse>('designs'),
  search: (q: string) => fetchJSON<SearchResponse>(`search?q=${encodeURIComponent(q)}`),
  constraints: () => fetchJSON<ConstraintsResponse>('constraints'),
  tree: () => fetchJSON<TreeResponse>('tree'),
  treeNode: (path: string) => fetchJSON<TreeContextBundle>(`tree/${encodeURIComponent(path)}`),
};

export interface TreeNode {
  type: 'file' | 'dir';
  parent: string;
  children?: string[];
  size?: number;
  hash?: string;
  mtime?: string;
  binary?: boolean;
  summary?: string;
  summary_stale?: boolean;
  updated_at?: string;
  ignored?: boolean;
}

export interface TreeResponse {
  scanned_at: string;
  root: string;
  nodes: Record<string, TreeNode>;
}

export interface TreeContextBundle {
  path: string;
  type: 'file' | 'dir';
  breadcrumb: { path: string; summary: string; stale: boolean }[];
  summary: string;
  stale: boolean;
  size?: number;
  content?: string;
  total_bytes?: number;
  truncated?: boolean;
  children?: { path: string; type: string; summary: string; stale: boolean }[];
}

// Types
export interface Project {
  slug: string;
  path: string;
  name: string;
  last_opened: string;
}

export interface StatusResponse {
  counts: Record<string, number>;
  total: number;
}

export interface EntitySummary {
  id: string;
  kind: string;
  name: string;
  slug: string;
  description: string;
  file: string;
  relationship_count: number;
  learning_count: number;
  tags?: string[];
  files?: string[];
  relationships?: { type: string; target: string; label?: string }[];
  attributes?: { name: string; description?: string; refs?: string[] }[];
  actions?: { name: string; description?: string }[];
  meaning?: string;
  contract_kind?: string;
  interaction_pattern?: string;
}

export interface EntitiesResponse {
  entities: EntitySummary[];
  count: number;
}

export interface FileRef {
  path: string;
  summary: string;
  stale: boolean;
  in_tree: boolean;
}

export interface EntityDetailResponse {
  entity: Record<string, any>;
  file_refs?: FileRef[];
  relationships: Relationship[];
  learnings: any[];
  tasks: any[];
  suggested_queries: string[];
}

export interface Relationship {
  target_id: string;
  target_slug: string;
  target_name: string;
  target_kind: string;
  direction?: 'outbound' | 'inbound';
  type: string;
}

export interface GraphNode {
  id: string;
  name: string;
  kind: string;
}

export interface GraphEdge {
  source: string;
  target: string;
  type: string;
  label?: string;
}

export interface GraphResponse {
  nodes: GraphNode[];
  edges: GraphEdge[];
}

export interface PlanPhase {
  id: string;
  name?: string;
  parent_phase?: string;
  action: string;
  entity_kind: string;
  entity_name: string;
  entities?: { kind: string; name: string; data: Record<string, any> }[];
  status: string;
  description: string;
  details?: string;
  notes?: string;
  tasks?: string[];
}

export interface Plan {
  name: string;
  status: string;
  progress: number;
  phases: PlanPhase[];
  created: string;
}

export interface PlansResponse {
  plans: Plan[];
}

export interface Learning {
  name: string;
  category: string;
  confidence: string;
  description: string;
  entity_refs: string[];
  file: string;
}

export interface LearningsResponse {
  learnings: Learning[];
}

export interface Task {
  name: string;
  status: string;
  priority: string;
  plan_ref: string;
}

export interface TasksResponse {
  tasks: Task[];
}

export interface DesignsResponse {
  designs: { name: string; design_type: string }[];
}

export interface SearchHit {
  id: string;
  kind: string;
  name: string;
  slug: string;
  file: string;
  field: string;
}

export interface SearchResponse {
  query: string;
  results: SearchHit[];
  count: number;
}

export interface ConstraintsResponse {
  decisions: { name: string; statement: string; category: string }[];
  learnings: { name: string; category: string; description: string }[];
}
