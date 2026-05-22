package database

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"

	"evex/models"
)

type DB struct {
	mu       sync.RWMutex
	filePath string
	Data     Store
}

type Store struct {
	Projects []models.Project `json:"projects"`
	Sections []models.Section `json:"sections"`
	Sources  []models.Source  `json:"sources"`
	Assets   []models.Asset   `json:"assets"`
}

func New() (*DB, error) {
	dir, err := ensureDir()
	if err != nil {
		return nil, err
	}
	db := &DB{filePath: filepath.Join(dir, "evex.json")}
	if err := db.load(); err != nil {
		return nil, err
	}
	return db, nil
}

func ensureDir() (string, error) {
	dir := filepath.Join(os.TempDir(), "evex")
	return dir, os.MkdirAll(dir, 0755)
}

func (db *DB) load() error {
	data, err := os.ReadFile(db.filePath)
	if err != nil {
		db.Data = Store{}
		return nil
	}
	return json.Unmarshal(data, &db.Data)
}

func (db *DB) save() error {
	data, err := json.MarshalIndent(db.Data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(db.filePath, data, 0644)
}

func (db *DB) AddProject(p models.Project) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.Data.Projects = append(db.Data.Projects, p)
	return db.save()
}

func (db *DB) GetProjects() []models.Project {
	db.mu.RLock()
	defer db.mu.RUnlock()
	return db.Data.Projects
}

func (db *DB) AddSection(s models.Section) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.Data.Sections = append(db.Data.Sections, s)
	return db.save()
}

func (db *DB) GetSections(projectID string) []models.Section {
	db.mu.RLock()
	defer db.mu.RUnlock()
	var result []models.Section
	for _, s := range db.Data.Sections {
		if s.ProjectID == projectID {
			result = append(result, s)
		}
	}
	return result
}

func (db *DB) AddSource(s models.Source) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.Data.Sources = append(db.Data.Sources, s)
	return db.save()
}

func (db *DB) GetSources(sectionID string) []models.Source {
	db.mu.RLock()
	defer db.mu.RUnlock()
	var result []models.Source
	for _, s := range db.Data.Sources {
		if s.SectionID == sectionID {
			result = append(result, s)
		}
	}
	return result
}
