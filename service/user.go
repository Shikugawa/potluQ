package service

import (
	"context"

	"github.com/Shikugawa/potraq/controller"
	"github.com/Shikugawa/potraq/ent"
	"github.com/Shikugawa/potraq/ent/user"
)

type UserService struct {
	client *ent.Client
}

func (service *UserService) CreateUser(ctx context.Context, user *ent.User) error {
	_, err := service.client.User.
		Create().
		SetEmail(user.Email).
		SetName(user.Name).
		SetPassword(user.Password).
		Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (service *UserService) FindByUser(ctx context.Context, email string) (*ent.User, error) {
	user, err := service.client.User.
		Query().
		Where(user.EmailEQ(email)).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (service *UserService) VerityPassword(ctx context.Context, credential *controller.Credential) bool {
	user, err := service.FindByUser(ctx, credential.Email)
	if err != nil {
		return false
	}
	if credential.Password != user.Password {
		return false
	}
	return true
}
