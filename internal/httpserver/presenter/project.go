package presenter

import "github.com/marcos-nsantos/portfolio-api/internal/entity"

type ProjectPresenter struct {
	ID          uint64 `json:"id"`
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

func NewProjectsPresenter(projects []*entity.Project) []*ProjectPresenter {
	var projectPresenters []*ProjectPresenter
	for _, project := range projects {
		projectPresenters = append(projectPresenters, NewProjectPresenter(project))
	}
	return projectPresenters
}
