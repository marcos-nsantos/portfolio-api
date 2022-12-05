//go:build unit

package user

import (
	"context"
	"github.com/marcos-nsantos/portfolio-api/internal/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func newFixture() *entity.User {
	return &entity.User{
		FirstName: "Marcos",
		LastName:  "Santos",
		Email:     "email@email.com",
		Password:  "password",
	}
}

func TestCreate(t *testing.T) {
	repo := newInMemory()
	services := NewServices(repo)
	user := newFixture()
	err := services.Create(context.Background(), user)
	assert.NoError(t, err)
}

func TestGetByID(t *testing.T) {
	repo := newInMemory()
	services := NewServices(repo)
	user := newFixture()
	err := services.Create(context.Background(), user)
	assert.NoError(t, err)
	user, err = services.GetByID(context.Background(), user.ID)
	assert.NoError(t, err)
	assert.NotEmpty(t, user.ID)
	assert.Equal(t, "Marcos", user.FirstName)
	assert.Equal(t, "Santos", user.LastName)
	assert.Equal(t, "email@email.com", user.Email)
}

func TestGetAll(t *testing.T) {
	repo := newInMemory()
	services := NewServices(repo)
	user := newFixture()
	err := services.Create(context.Background(), user)
	assert.NoError(t, err)
	users, err := services.GetAll(context.Background())
	assert.NoError(t, err)
	assert.Len(t, users, 1)
}
