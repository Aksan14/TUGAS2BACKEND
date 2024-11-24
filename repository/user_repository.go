package repository

import (
	"context"
	"golang-database-user/model"
)

type UserRepository interface {
	InsertUser(ctx context.Context, user model.MstUser) (model.MstUser, error)
	TesEmail(ctx context.Context, email string) (*model.MstUser, error)
	UpdateUser(ctx context.Context, user model.MstUser, userId string) (model.MstUser, error)
	DeleteUser(ctx context.Context, userId string) error
	ReadUsers(ctx context.Context) ([]model.MstUser, error)
}