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
