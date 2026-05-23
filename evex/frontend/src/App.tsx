import { useState, useEffect } from 'react';
import './App.css';

interface Project { id: string; name: string; description: string; created_at: string; updated_at: string; }
interface Section { id: string; project_id: string; title: string; content: string; order: number; tags: string[]; created_at: string; }

function App() {
  const [projects, setProjects] = useState<Project[]>([]);
  const [sections, setSections] = useState<Section[]>([]);
  const [view, setView] = useState<"welcome" | "project">("welcome");
  const [selectedProject, setSelectedProject] = useState<string | null>(null);

  const go = typeof window !== "undefined" && (window as any).go?.main?.App;

  const loadProjects = async () => {
    if (!go) return;
    try {
      const list = await go.GetProjects();
      setProjects(list);
      if (list.length > 0) setView("project");
    } catch {}
  };

  useEffect(() => { loadProjects(); }, []);

  const handleCreateProject = async () => {
    if (!go) return;
    const name = prompt("Nom du projet :");
    if (!name) return;
    await go.CreateProject(name, "");
    loadProjects();
  };

  const handleSelectProject = async (id: string) => {
    setSelectedProject(id);
    if (!go) return;
    const list = await go.GetSections(id);
    setSections(list);
  };

  const handleCreateSection = async () => {
    if (!go || !selectedProject) return;
    const title = prompt("Titre de la section :");
    if (!title) return;
    try {
      await go.CreateSection(selectedProject, title);
      const list = await go.GetSections(selectedProject);
      setSections(list);
    } catch (e) {
      console.error("CreateSection failed", e);
    }
  };

  if (view === "welcome") {
    return (
      <div className="welcome">
        <div className="welcome-content">
          <h1>Evex Research</h1>
          <p className="subtitle">
            Organise tes recherches, indexe le web,<br />
            garde tout au même endroit.
          </p>
          {go && (
            <div className="actions">
              <button className="btn-primary" onClick={handleCreateProject}>+ Créer un projet</button>
              <button className="btn-secondary" onClick={() => setView("project")}>Voir mes projets</button>
            </div>
          )}
        </div>
      </div>
    );
  }

  return (
    <div className="app">
      <header>
        <h1>Evex Research</h1>
        <div className="header-actions">
          <button className="btn-small" onClick={() => setView("welcome")}>Accueil</button>
          <button className="btn-small" onClick={handleCreateProject}>+ Projet</button>
          {selectedProject && <button className="btn-small" onClick={handleCreateSection}>+ Section</button>}
        </div>
      </header>
      <div className="layout">
        <aside className="sidebar">
          <h3>Projets</h3>
          <ul>
            {projects.map(p => (
              <li key={p.id} className={selectedProject === p.id ? "active" : ""} onClick={() => handleSelectProject(p.id)}>
                {p.name}
              </li>
            ))}
          </ul>
        </aside>
        <main>
          {selectedProject ? (
            <div className="sections">
              <h2>Sections</h2>
              {sections.length === 0 && <p className="empty">Aucune section. Créez-en une.</p>}
              {sections.map(s => (
                <div key={s.id} className="card">
                  <h3>{s.title}</h3>
                  {s.content && <p>{s.content}</p>}
                </div>
              ))}
            </div>
          ) : (
            <p className="empty">Sélectionne un projet</p>
          )}
        </main>
      </div>
    </div>
  );
}

export default App;
