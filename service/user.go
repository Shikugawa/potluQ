package service

import (
	"context"

	"github.com/Shikugawa/potraq/ent"
	"github.com/Shikugawa/potraq/ent/user"
	"github.com/Shikugawa/potraq/message"
)

type UserService struct {
	Client *ent.Client
}

func (service *UserService) CreateUser(ctx context.Context, user *ent.User) error {
	_, err := service.Client.User.
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

func (service *UserService) FindByEmail(ctx context.Context, email string) (*ent.User, error) {
	user, err := service.Client.User.
		Query().
		Where(user.EmailEQ(email)).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (service *UserService) FindByUserId(ctx context.Context, name string) (*ent.User, error) {
	user, err := service.Client.User.
		Query().
		Where(user.NameEQ(name)).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (service *UserService) VerityPassword(ctx context.Context, credential *message.Credential) bool {
	user, err := service.FindByEmail(ctx, credential.Email)
	if err != nil {
		return false
	}
	if credential.Password != user.Password {
		return false
	}
	return true
}
