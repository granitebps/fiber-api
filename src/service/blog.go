package service

import (
	"context"
	"errors"

	"github.com/ansel1/merry/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/granitebps/fiber-api/pkg/core"
	"github.com/granitebps/fiber-api/src/model"
	"github.com/granitebps/fiber-api/src/repository"
	"github.com/granitebps/fiber-api/src/request"
	"github.com/granitebps/fiber-api/src/transformer"
	"gorm.io/gorm"
)

type BlogService struct {
	Core *core.Core

	BlogRepo *repository.BlogRepository
}

func NewBlogService(c *core.Core, blogRepo *repository.BlogRepository) *BlogService {
	return &BlogService{
		Core:     c,
		BlogRepo: blogRepo,
	}
}

func (s *BlogService) GetAllBlog(ctx context.Context) (result []transformer.BlogTransformer, err error) {
	res, err := s.BlogRepo.GetAll(ctx)
	if err != nil {
		err = merry.Wrap(err)
		return
	}

	for _, b := range res {
		result = append(result, transformer.BlogTransformer{
			ID:          uint64(b.ID),
			Title:       b.Title,
			Description: b.Description,
			Image:       b.Image,
			CreatedAt:   b.CreatedAt,
		})
	}

	return
}

func (s *BlogService) GetByID(ctx context.Context, id uint64) (result transformer.BlogTransformer, err error) {
	res, err := s.BlogRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = merry.Wrap(err, merry.WithHTTPCode(fiber.StatusNotFound), merry.WithUserMessage("blog not found"))
			return
		}
		return
	}

	result.ID = uint64(res.ID)
	result.Title = res.Title
	result.Description = res.Description
	result.Image = res.Image
	result.CreatedAt = res.CreatedAt

	return
}

func (s *BlogService) Create(ctx context.Context, req request.CreateBlogRequest) (result transformer.BlogTransformer, err error) {
	var res model.Blog
	res.Title = req.Title
	res.Description = req.Description
	res.Image = req.Image

	err = s.BlogRepo.Insert(ctx, &res)
	if err != nil {
		err = merry.Wrap(err)
		return
	}

	result.ID = uint64(res.ID)
	result.Title = res.Title
	result.Description = res.Description
	result.Image = res.Image
	result.CreatedAt = res.CreatedAt

	return
}

func (s *BlogService) Update(ctx context.Context, id uint64, req request.UpdateBlogRequest) (result transformer.BlogTransformer, err error) {
	res, err := s.BlogRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = merry.Wrap(err, merry.WithHTTPCode(fiber.StatusNotFound), merry.WithUserMessage("blog not found"))
			return
		}
		return
	}

	res.Title = req.Title
	res.Description = req.Description
	res.Image = req.Image

	err = s.BlogRepo.Update(ctx, id, &res)
	if err != nil {
		err = merry.Wrap(err)
		return
	}

	result.ID = uint64(res.ID)
	result.Title = res.Title
	result.Description = res.Description
	result.Image = res.Image
	result.CreatedAt = res.CreatedAt

	return
}

func (s *BlogService) Delete(ctx context.Context, id uint64) (err error) {
	_, err = s.BlogRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = merry.Wrap(err, merry.WithHTTPCode(fiber.StatusNotFound), merry.WithUserMessage("blog not found"))
			return
		}
		return
	}

	err = s.BlogRepo.Delete(ctx, id)
	if err != nil {
		err = merry.Wrap(err)
		return
	}

	return
}
