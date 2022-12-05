package project

import (
	"context"
	"time"

	"github.com/marcos-nsantos/portfolio-api/internal/entity"
	"gorm.io/gorm"
)

type inMemory struct {
	projects []*entity.Project
}

func newInMemory() *inMemory {
	return &inMemory{
		projects: make([]*entity.Project, 0),
	}
}

func (i *inMemory) Insert(_ context.Context, project *entity.Project) error {
	project.ID = i.generateID()
	project.CreatedAt = time.Now()
	project.UpdatedAt = time.Now()
	i.projects = append(i.projects, project)
	return nil
}

func (i *inMemory) FindByID(_ context.Context, id uint64) (*entity.Project, error) {
	for _, project := range i.projects {
		if project.ID == id {
			return project, nil
		}
	}
	return nil, gorm.ErrRecordNotFound
}

func (i *inMemory) FindAll(_ context.Context) ([]*entity.Project, error) {
	return i.projects, nil
}

func (i *inMemory) Update(_ context.Context, project *entity.Project) error {
	for index, p := range i.projects {
		if p.ID == project.ID {
			i.projects[index] = project
			return nil
		}
	}
	return nil
}

func (i *inMemory) Delete(_ context.Context, id uint64) error {
	for index, project := range i.projects {
		if project.ID == id {
			i.projects = append(i.projects[:index], i.projects[index+1:]...)
			return nil
		}
	}
	return nil
}

func (i *inMemory) generateID() uint64 {
	return uint64(len(i.projects) + 1)
}
