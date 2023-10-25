package repository

import (
	"context"

	"github.com/granitebps/fiber-api/pkg/core"
	"github.com/granitebps/fiber-api/src/model"
)

type BlogRepository struct {
	Core *core.Core
}

func NewBlogRepository(c *core.Core) *BlogRepository {
	return &BlogRepository{
		Core: c,
	}
}

func (r *BlogRepository) GetAll(ctx context.Context) (res []model.Blog, err error) {
	err = r.Core.Database.Db.
		WithContext(ctx).
		Find(&res).
		Error
	return
}

func (r *BlogRepository) GetByID(ctx context.Context, id uint64) (res model.Blog, err error) {
	err = r.Core.Database.Db.
		WithContext(ctx).
		Where("id = ?", id).
		First(&res).
		Error
	return
}

func (r *BlogRepository) Insert(ctx context.Context, data *model.Blog) (err error) {
	err = r.Core.Database.Db.
		WithContext(ctx).
		Create(&data).
		Error
	return
}

func (r *BlogRepository) Update(ctx context.Context, id uint64, data *model.Blog) (err error) {
	err = r.Core.Database.Db.
		WithContext(ctx).
		Model(&model.Blog{}).
		Where("id = ?", id).
		Updates(&data).
		Error
	return
}

func (r *BlogRepository) Delete(ctx context.Context, id uint64) (err error) {
	err = r.Core.Database.Db.
		WithContext(ctx).
		Where("id = ?", id).
		Delete(&model.Blog{}).
		Error
	return
}
