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
