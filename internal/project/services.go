package project

import (
	"context"

	"github.com/marcos-nsantos/portfolio-api/internal/entity"
)

type Services struct {
	repo Repository
}

func NewServices(repo Repository) *Services {
	return &Services{repo: repo}
}

func (s *Services) Create(ctx context.Context, project *entity.Project) error {
	return s.repo.Insert(ctx, project)
}

func (s *Services) GetByID(ctx context.Context, id uint) (*entity.Project, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *Services) GetAll(ctx context.Context) ([]*entity.Project, error) {
	return s.repo.FindAll(ctx)
}

func (s *Services) Update(ctx context.Context, project *entity.Project) error {
	if _, err := s.GetByID(ctx, project.ID); err != nil {
		return err
	}
	return s.repo.Update(ctx, project)
}

func (s *Services) Delete(ctx context.Context, id uint) error {
	if _, err := s.GetByID(ctx, id); err != nil {
		return err
	}
	return s.repo.Delete(ctx, id)
}
