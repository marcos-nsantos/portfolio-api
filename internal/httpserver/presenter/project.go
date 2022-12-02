package presenter

import "github.com/marcos-nsantos/portfolio-api/internal/entity"

type ProjectPresenter struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

func NewProjectPresenter(project *entity.Project) *ProjectPresenter {
	return &ProjectPresenter{
		ID:          project.ID,
		Name:        project.Name,
		Description: project.Description,
		URL:         project.URL,
	}
}