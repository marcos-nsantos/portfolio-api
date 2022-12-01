package project

import (
	"context"
	"time"

	"github.com/marcos-nsantos/portfolio-api/internal/entity"
)

type inMemory struct {
	projects []*entity.Project
}

func newInMemory() *inMemory {
	return &inMemory{
		projects: make([]*entity.Project, 0),
	}
}

func (i *inMemory) Insert(ctx context.Context, project *entity.Project) error {
	project.ID = i.generateID()
	project.CreatedAt = time.Now()
	project.UpdatedAt = time.Now()
	i.projects = append(i.projects, project)
	return nil
}

func (i *inMemory) Update(ctx context.Context, project *entity.Project) error {
	//TODO implement me
	panic("implement me")
}

func (i *inMemory) Delete(ctx context.Context, id uint) error {
	//TODO implement me
	panic("implement me")
}

func (i *inMemory) FindAll(ctx context.Context) ([]*entity.Project, error) {
	//TODO implement me
	panic("implement me")
}

func (i *inMemory) FindByID(ctx context.Context, id uint) (*entity.Project, error) {
	//TODO implement me
	panic("implement me")
}

func (i *inMemory) generateID() uint {
	return uint(len(i.projects) + 1)
}
