package project

import (
	"context"

	"github.com/marcos-nsantos/portfolio-api/internal/entity"
)

type Writer interface {
	Insert(ctx context.Context, project *entity.Project) error
}

type Reader interface{}

type Repository interface {
	Writer
	Reader
}
