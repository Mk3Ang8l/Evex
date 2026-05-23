package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"

	"evex/database"
	"evex/models"
)

type App struct {
	ctx context.Context
	db  *database.DB
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	db, err := database.New()
	if err != nil {
		fmt.Println("DB error:", err)
		return
	}
	a.db = db
	fmt.Println("Evex ready")
}


// fonction qui cree des projets
func (a *App) CreateProject(name string, description string) (models.Project, error) {
	p := models.NewProject(newID(), name, description)
	return p, a.db.AddProject(p)
}

// fonction qui récupère les projets
func (a *App) GetProjects() ([]models.Project, error) {
	return a.db.GetProjects(), nil
}

// fonction qui crée des sections
func (a *App) CreateSection(projectID string, title string) (models.Section, error) {
	s := models.NewSection(newID(), projectID, title)
	return s, a.db.AddSection(s)
}

// fonction qui récupère les sections
func (a *App) GetSections(projectID string) ([]models.Section, error) {
	return a.db.GetSections(projectID), nil
}

// fonction qui crée des sources
func (a *App) AddSource(sectionID string, title string, url string) (models.Source, error) {
	s := models.NewSource(newID(), sectionID, title, url)
	return s, a.db.AddSource(s)
}

// fonction qui récupère les sources
func (a *App) GetSources(sectionID string) ([]models.Source, error) {
	return a.db.GetSources(sectionID), nil
}

// fonction qui crée des assets
func (a *App) AddAsset(sourceID string, filename string) (models.Asset, error) {
	asset := models.Asset{ID: newID(), SourceID: sourceID, Filename: filename}
	return asset, a.db.AddAsset(asset)
}

func newID() string {
	n, _ := rand.Int(rand.Reader, big.NewInt(9007199254740991))
	return fmt.Sprintf("%d", n.Int64())
}
