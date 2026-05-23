package models

import "time"

type Project struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func NewProject(id, name, desc string) Project {
	now := time.Now().Format(time.RFC3339)
	return Project{ID: id, Name: name, Description: desc, CreatedAt: now, UpdatedAt: now}
}

type Section struct {
	ID        string   `json:"id"`
	ProjectID string   `json:"project_id"`
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	Order     int      `json:"order"`
	Tags      []string `json:"tags"`
	CreatedAt string   `json:"created_at"`
}

func NewSection(id, projectID, title string) Section {
	return Section{ID: id, ProjectID: projectID, Title: title, CreatedAt: time.Now().Format(time.RFC3339)}
}

type Source struct {
	ID        string `json:"id"`
	SectionID string `json:"section_id"`
	Title     string `json:"title"`
	URL       string `json:"url"`
	Snippet   string `json:"snippet"`
	Content   string `json:"content"`
	Tags      []string `json:"tags"`
	Notes     string `json:"notes"`
	CreatedAt string `json:"created_at"`
}

func NewSource(id, sectionID, title, url string) Source {
	return Source{ID: id, SectionID: sectionID, Title: title, URL: url, CreatedAt: time.Now().Format(time.RFC3339)}
}

type Asset struct {
	ID       string `json:"id"`
	SourceID string `json:"source_id"`
	Filename string `json:"filename"`
	Data     []byte `json:"-"`
	MimeType string `json:"mime_type"`
}
