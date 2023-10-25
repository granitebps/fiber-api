package repository

import (
	"context"

	"github.com/granitebps/fiber-api/pkg/core"
	"github.com/granitebps/fiber-api/src/model"
)

type UserRepository struct {
	Core *core.Core
}

func NewUserRepository(c *core.Core) *UserRepository {
	return &UserRepository{
		Core: c,
	}
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (res model.User, err error) {
	err = r.Core.Database.Db.
		WithContext(ctx).
		Where("email = ?", email).
		First(&res).
		Error
	return
}

func (r *UserRepository) Insert(ctx context.Context, data *model.User) (err error) {
	err = r.Core.Database.Db.
		WithContext(ctx).
		Create(&data).
		Error
	return
}
