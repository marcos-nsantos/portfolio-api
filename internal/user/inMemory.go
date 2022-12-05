package user

import (
	"context"
	"github.com/marcos-nsantos/portfolio-api/internal/entity"
)

type inMemory struct {
	users []*entity.User
}

func newInMemory() *inMemory {
	return &inMemory{}
}

func (i *inMemory) Insert(ctx context.Context, user *entity.User) error {
	user.ID = i.generateID()
	i.users = append(i.users, user)
	return nil
}

func (i *inMemory) Update(ctx context.Context, user *entity.User) error {
	//TODO implement me
	panic("implement me")
}

func (i *inMemory) UpdatePassword(ctx context.Context, user *entity.User) error {
	//TODO implement me
	panic("implement me")
}

func (i *inMemory) Delete(ctx context.Context, id uint64) error {
	//TODO implement me
	panic("implement me")
}

func (i *inMemory) FindAll(ctx context.Context) ([]*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (i *inMemory) FindByID(ctx context.Context, id uint64) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (i *inMemory) generateID() uint64 {
	return uint64(len(i.users) + 1)
}
