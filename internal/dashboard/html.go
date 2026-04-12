package dashboard

var indexHTML = `<!DOCTYPE html>
<html lang="en" class="dark">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>syde</title>
<script src="https://cdn.tailwindcss.com"></script>
<script>
tailwind.config = {
  darkMode: 'class',
  theme: {
    extend: {
      colors: {
        background: '#09090b',
        foreground: '#fafafa',
        card: '#18181b',
        'card-foreground': '#fafafa',
        border: '#27272a',
        muted: '#27272a',
        'muted-foreground': '#a1a1aa',
        accent: '#27272a',
        'accent-foreground': '#fafafa',
      },
      fontFamily: {
        sans: ['Inter', 'system-ui', '-apple-system', 'sans-serif'],
        mono: ['JetBrains Mono', 'Menlo', 'monospace'],
      },
    }
  }
}
</script>
<style>
  body{background:#09090b;color:#fafafa}
  .badge{display:inline-flex;align-items:center;padding:2px 8px;border-radius:9999px;font-size:.7rem;font-weight:500;border:1px solid #27272a}
  .badge-active{color:#22c55e;border-color:#166534}
  .badge-draft{color:#a1a1aa;border-color:#3f3f46}
  .badge-gotcha,.badge-constraint{color:#f59e0b;border-color:#78350f}
  .card{background:#18181b;border:1px solid #27272a;border-radius:.5rem}
  .card:hover{border-color:#3f3f46}
  .sidebar-item{padding:.5rem .75rem;border-radius:.375rem;cursor:pointer;color:#a1a1aa;font-size:.875rem;transition:all .15s}
  .sidebar-item:hover{background:#27272a;color:#fafafa}
  .sidebar-item.active{background:#27272a;color:#fafafa}
  .table-row{border-bottom:1px solid #27272a}
  .table-row:hover{background:#1a1a1e}
  .progress-bar{height:6px;background:#27272a;border-radius:3px;overflow:hidden}
  .progress-fill{height:100%;background:#fafafa;border-radius:3px;transition:width .3s}
</style>
</head>
<body class="font-sans antialiased">
<div id="app" class="flex h-screen">
  <aside class="w-56 border-r border-border flex flex-col py-4 px-3 flex-shrink-0">
    <div class="flex items-center gap-2 px-3 mb-6">
      <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M12 2L2 7l10 5 10-5-10-5z"/><path d="M2 17l10 5 10-5"/><path d="M2 12l10 5 10-5"/></svg>
      <span class="text-sm font-semibold tracking-tight">syde</span>
    </div>
    <nav class="flex flex-col gap-1" id="nav"></nav>
    <div class="mt-auto pt-4 border-t border-border">
      <div class="px-3 text-xs text-muted-foreground" id="project-name">Loading...</div>
    </div>
  </aside>
  <main class="flex-1 overflow-y-auto">
    <div class="max-w-5xl mx-auto px-6 py-8" id="content">
      <div class="text-muted-foreground text-sm">Loading...</div>
    </div>
  </main>
</div>

<script>
const API='/api';
let proj=null, page='overview';

const pages=['overview','entities','plans','learnings','tasks','designs'];

function nav(p){
  page=p;
  document.querySelectorAll('.sidebar-item').forEach(el=>el.classList.toggle('active',el.dataset.page===p));
  render();
}

function render(){
  const c=document.getElementById('content');
  switch(page){
    case 'overview': renderOverview(c); break;
    case 'entities': renderEntities(c); break;
    case 'plans': renderPlans(c); break;
    case 'learnings': renderLearnings(c); break;
    case 'tasks': renderTasks(c); break;
    case 'designs': renderDesigns(c); break;
  }
}

async function init(){
  const res=await fetch(API+'/projects');
  const data=await res.json();
  if(!data.projects||!data.projects.length){
    document.getElementById('content').innerHTML='<p class="text-muted-foreground mt-20 text-center">No projects. Run <code class="font-mono bg-card px-2 py-1 rounded text-xs border border-border">syde open</code></p>';
    return;
  }
  // Find project from URL path
  const pathSlug=location.pathname.replace(/^\//,'').replace(/\/$/,'');
  proj=data.projects.find(p=>p.slug===pathSlug)||data.projects[0];
  document.getElementById('project-name').textContent=proj.name||proj.slug;

  // Build nav
  const navEl=document.getElementById('nav');
  navEl.innerHTML=pages.map(p=>'<a class="sidebar-item'+(p==='overview'?' active':'')+'" data-page="'+p+'" onclick="nav(\''+p+'\')">'+p[0].toUpperCase()+p.slice(1)+'</a>').join('');
  render();
}

async function api(path){
  const res=await fetch(API+'/'+proj.slug+'/'+path);
  return res.json();
}

async function renderOverview(el){
  const [status,plans,learnings]=await Promise.all([api('status'),api('plans'),api('learnings')]);
  const c=status.counts||{};
  const stats=Object.entries(c).map(([k,v])=>'<div class="card p-4"><div class="text-xs text-muted-foreground uppercase tracking-wider">'+k+'</div><div class="text-2xl font-semibold mt-1">'+v+'</div></div>').join('');

  let planHTML='<p class="text-sm text-muted-foreground">No plans yet.</p>';
  if(plans.plans&&plans.plans.length){
    planHTML=plans.plans.map(p=>{
      const pct=Math.round(p.progress);
      return '<div class="card p-4"><div class="flex items-center justify-between mb-2"><span class="text-sm font-medium">'+p.name+'</span><span class="badge badge-'+p.status+'">'+p.status+'</span></div><div class="progress-bar"><div class="progress-fill" style="width:'+pct+'%"></div></div><div class="text-xs text-muted-foreground mt-1">'+pct+'% complete</div></div>';
    }).join('');
  }

  let learnHTML='';
  if(learnings.learnings&&learnings.learnings.length){
    learnHTML=learnings.learnings.map(l=>{
      const icon=l.category==='gotcha'||l.category==='constraint'?'⚠':'ℹ';
      return '<div class="flex items-center gap-2 py-1.5"><span>'+icon+'</span><span class="badge badge-'+l.category+'">'+l.category+'</span><span class="text-sm">'+esc(l.description)+'</span><span class="badge">'+l.confidence+'</span></div>';
    }).join('');
  }

  el.innerHTML='<div class="mb-8"><h1 class="text-2xl font-semibold tracking-tight">'+(proj.name||proj.slug)+'</h1><p class="text-muted-foreground text-sm mt-1">Design model overview</p></div>'
    +'<div class="grid grid-cols-2 sm:grid-cols-4 gap-4 mb-8">'+stats+'</div>'
    +(planHTML?'<h2 class="text-sm font-medium text-muted-foreground uppercase tracking-wider mb-3">Plans</h2><div class="space-y-3 mb-8">'+planHTML+'</div>':'')
    +(learnHTML?'<h2 class="text-sm font-medium text-muted-foreground uppercase tracking-wider mb-3">Learnings</h2><div class="card p-4">'+learnHTML+'</div>':'');
}

async function renderEntities(el){
  const data=await api('entities');
  const entities=data.entities||[];
  const rows=entities.map(e=>'<tr class="table-row"><td class="p-3"><span class="badge badge-draft">'+e.kind+'</span></td><td class="p-3 font-medium">'+esc(e.name)+'</td><td class="p-3"><span class="badge badge-'+e.status+'">'+e.status+'</span></td><td class="p-3 text-muted-foreground text-sm">'+esc(e.description||'')+'</td><td class="p-3 text-muted-foreground text-sm">'+e.relationship_count+' rels'+(e.learning_count?' · '+e.learning_count+' ⚠':'')+'</td></tr>').join('');
  el.innerHTML='<div class="mb-6"><h1 class="text-xl font-semibold tracking-tight">Entities</h1><p class="text-muted-foreground text-sm mt-1">'+entities.length+' total</p></div><div class="card"><table class="w-full text-sm"><thead><tr class="border-b border-border text-left text-muted-foreground"><th class="p-3 font-medium">Kind</th><th class="p-3 font-medium">Name</th><th class="p-3 font-medium">Status</th><th class="p-3 font-medium">Description</th><th class="p-3 font-medium">Refs</th></tr></thead><tbody>'+rows+'</tbody></table></div>';
}

async function renderPlans(el){
  const data=await api('plans');
  const plans=data.plans||[];
  if(!plans.length){el.innerHTML='<h1 class="text-xl font-semibold mb-4">Plans</h1><p class="text-muted-foreground">No plans yet.</p>';return;}
  const html=plans.map(p=>{
    const pct=Math.round(p.progress);
    const steps=(p.steps||[]).map(s=>{
      const icon=s.status==='completed'?'✓':s.status==='in_progress'?'●':'○';
      return '<div class="flex items-center gap-2 py-1"><span class="'+(s.status==='completed'?'text-green-400':'text-muted-foreground')+'">'+icon+'</span><span class="text-sm">'+esc(s.description||s.entity_name||s.id)+'</span><span class="badge badge-draft">'+s.action+' '+s.entity_kind+'</span></div>';
    }).join('');
    return '<div class="card p-4 mb-4"><div class="flex items-center justify-between mb-3"><h2 class="font-medium">'+esc(p.name)+'</h2><span class="badge badge-'+p.status+'">'+p.status+'</span></div><div class="progress-bar mb-2"><div class="progress-fill" style="width:'+pct+'%"></div></div><div class="text-xs text-muted-foreground mb-3">'+pct+'% complete'+(p.created?' · Created '+p.created.slice(0,10):'')+'</div>'+steps+'</div>';
  }).join('');
  el.innerHTML='<h1 class="text-xl font-semibold mb-4">Plans</h1>'+html;
}

async function renderLearnings(el){
  const data=await api('learnings');
  const items=data.learnings||[];
  if(!items.length){el.innerHTML='<h1 class="text-xl font-semibold mb-4">Learnings</h1><p class="text-muted-foreground">No learnings yet.</p>';return;}
  const grouped={};
  items.forEach(l=>{const c=l.category||'other';(grouped[c]=grouped[c]||[]).push(l)});
  let html='<h1 class="text-xl font-semibold mb-6">Learnings</h1>';
  for(const[cat,ls]of Object.entries(grouped)){
    html+='<h2 class="text-sm font-medium text-muted-foreground uppercase tracking-wider mb-2 mt-6">'+cat+'</h2><div class="space-y-2">';
    ls.forEach(l=>{
      const icon=cat==='gotcha'||cat==='constraint'?'⚠':'ℹ';
      html+='<div class="card p-4"><div class="flex items-center gap-2 mb-1"><span>'+icon+'</span><span class="badge badge-'+cat+'">'+cat+'</span><span class="badge">'+l.confidence+'</span></div><p class="text-sm">'+esc(l.description)+'</p>'+(l.entity_refs&&l.entity_refs.length?'<p class="text-xs text-muted-foreground mt-1">Entities: '+l.entity_refs.join(', ')+'</p>':'')+'<p class="text-xs text-muted-foreground mt-1 font-mono">'+l.file+'</p></div>';
    });
    html+='</div>';
  }
  el.innerHTML=html;
}

async function renderTasks(el){
  const data=await api('tasks');
  const items=data.tasks||[];
  if(!items.length){el.innerHTML='<h1 class="text-xl font-semibold mb-4">Tasks</h1><p class="text-muted-foreground">No tasks yet.</p>';return;}
  const rows=items.map(t=>{
    const icon={completed:'✓',in_progress:'●',blocked:'✗',cancelled:'–',pending:'○'}[t.status]||'○';
    const iconCls=t.status==='completed'?'text-green-400':t.status==='blocked'?'text-red-400':'text-muted-foreground';
    return '<tr class="table-row"><td class="p-3"><span class="'+iconCls+'">'+icon+'</span></td><td class="p-3 font-medium">'+esc(t.name)+'</td><td class="p-3"><span class="badge badge-'+t.status+'">'+t.status+'</span></td><td class="p-3"><span class="badge">'+t.priority+'</span></td><td class="p-3 text-muted-foreground text-sm">'+(t.plan_ref||'—')+'</td></tr>';
  }).join('');
  el.innerHTML='<h1 class="text-xl font-semibold mb-4">Tasks</h1><div class="card"><table class="w-full text-sm"><thead><tr class="border-b border-border text-left text-muted-foreground"><th class="p-3 w-8"></th><th class="p-3 font-medium">Name</th><th class="p-3 font-medium">Status</th><th class="p-3 font-medium">Priority</th><th class="p-3 font-medium">Plan</th></tr></thead><tbody>'+rows+'</tbody></table></div>';
}

async function renderDesigns(el){
  const data=await api('designs');
  const items=data.designs||[];
  if(!items.length){el.innerHTML='<h1 class="text-xl font-semibold mb-4">Designs</h1><p class="text-muted-foreground">No designs yet.</p>';return;}
  const cards=items.map(d=>'<div class="card p-4"><div class="h-24 bg-background rounded border border-border mb-3 flex items-center justify-center text-muted-foreground text-xs">UIML Preview</div><div class="font-medium text-sm">'+esc(d.name)+'</div><div class="text-xs text-muted-foreground">'+d.design_type+' · '+d.status+'</div></div>').join('');
  el.innerHTML='<h1 class="text-xl font-semibold mb-4">Designs</h1><div class="grid grid-cols-3 gap-4">'+cards+'</div>';
}

function esc(s){return s?s.replace(/&/g,'&amp;').replace(/</g,'&lt;').replace(/>/g,'&gt;'):''}

init();
</script>
</body>
</html>`
