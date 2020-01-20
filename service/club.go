package service

import (
	"context"

	"github.com/Shikugawa/potluq/ent"
)

type ClubService struct {
	Client *ent.Client
}

func (service *ClubService) GetAll(ctx context.Context) ([]*ent.Club, error) {
	clubs, err := service.Client.Club.
		Query().
		All(ctx)
	if err != nil {
		return nil, err
	}
	return clubs, nil
}

func (service *ClubService) Add(ctx context.Context, club *ent.Club, jukeBox *ent.User) error {
	_, err := service.Client.Club.
		Create().
		AddUser(jukeBox).
		SetName(club.Name).
		Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (service *ClubService) Delete(ctx context.Context, club *ent.Club) error {
	err := service.Client.Club.DeleteOne(club).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
