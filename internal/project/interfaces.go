package project

import (
	"context"

	"github.com/marcos-nsantos/portfolio-api/internal/entity"
)

type Writer interface {
	Insert(ctx context.Context, project *entity.Project) error
	Update(ctx context.Context, project *entity.Project) error
	Delete(ctx context.Context, id uint) error
}

type Reader interface {
	FindAll(ctx context.Context) ([]*entity.Project, error)
	FindByID(ctx context.Context, id uint) (*entity.Project, error)
}

type Repository interface {
	Writer
	Reader
}

type Service interface {
	Create(ctx context.Context, project *entity.Project) error
}
