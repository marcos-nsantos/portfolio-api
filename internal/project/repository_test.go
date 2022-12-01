//go:build integration

package project

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/marcos-nsantos/portfolio-api/internal/database"
	"github.com/marcos-nsantos/portfolio-api/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var db *gorm.DB

func TestMain(m *testing.M) {
	conn, err := database.New()
	if err != nil {
		fmt.Println(err)
	}
	db = conn.DB
	if err := conn.CreateTables(); err != nil {
		fmt.Println(err)
	}
	code := m.Run()
	os.Exit(code)
}

func TestInsert(t *testing.T) {
	project := &entity.Project{
		Name:        "Test",
		Description: "Test",
		URL:         "https://github.com/marcos-nsantos/test",
	}
	repo := NewRepo(db)
	err := repo.Insert(context.Background(), project)
	assert.NoError(t, err)
	assert.NotEmpty(t, project.ID)
	assert.NotEmpty(t, project.CreatedAt)
	assert.NotEmpty(t, project.UpdatedAt)
}

func TestFindAll(t *testing.T) {
	repo := NewRepo(db)
	projects, err := repo.FindAll(context.Background())
	assert.NoError(t, err)
	assert.NotEmpty(t, projects)
}

func TestFindByID(t *testing.T) {
	repo := NewRepo(db)
	project, err := repo.FindByID(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, uint(1), project.ID)
	assert.Equal(t, "Test", project.Name)
	assert.Equal(t, "Test", project.Description)
	assert.Equal(t, "https://github.com/marcos-nsantos/test", project.URL)
	assert.NotEmpty(t, project.CreatedAt)
	assert.NotEmpty(t, project.UpdatedAt)
}
