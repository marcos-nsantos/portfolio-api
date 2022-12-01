package project

import (
	"context"
	"testing"

	"github.com/marcos-nsantos/portfolio-api/internal/entity"
	"github.com/stretchr/testify/assert"
)

func newFixture() *entity.Project {
	return &entity.Project{
		Name:        "Test Name",
		Description: "Test Description",
		URL:         "https://github.com/marcos-nsantos/test",
	}
}

func TestCreate(t *testing.T) {
	repo := newInMemory()
	services := NewServices(repo)
	project := newFixture()
	err := services.Create(context.Background(), project)
	assert.NoError(t, err)
	assert.False(t, project.CreatedAt.IsZero())
}
