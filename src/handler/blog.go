package handler

import (
	"github.com/ansel1/merry/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/granitebps/fiber-api/pkg/core"
	"github.com/granitebps/fiber-api/pkg/utils"
	"github.com/granitebps/fiber-api/src/request"
	"github.com/granitebps/fiber-api/src/service"
)

type BlogHandler struct {
	Core *core.Core

	BlogService *service.BlogService
}

func NewBlogHandler(c *core.Core, blogService *service.BlogService) *BlogHandler {
	return &BlogHandler{
		Core:        c,
		BlogService: blogService,
	}
}

// GetBlogs
// @summary Get all blogs
// @tags blog
// @produce json
// @router /api/v1/blogs [get]
// @success 200 {object} utils.JSONResponse{data=[]transformer.BlogTransformer}
// @Security Bearer
func (h *BlogHandler) Index(c *fiber.Ctx) error {
	res, err := h.BlogService.GetAllBlog(c.UserContext())
	if err != nil {
		err = merry.Wrap(err)
		return utils.ReturnErrorResponse(c, err, nil)
	}

	return utils.ReturnSuccessResponse(c, fiber.StatusOK, "Success", res)
}

// GetBlog
// @summary Get single blog by ID
// @tags blog
// @produce json
// @param id path int true "blog id"
// @router /api/v1/blogs/{id} [get]
// @success 200 {object} utils.JSONResponse{data=transformer.BlogTransformer}
// @failure 404
// @Security Bearer
func (h *BlogHandler) Show(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		err = merry.Wrap(err, merry.WithHTTPCode(fiber.StatusBadRequest))
		return utils.ReturnErrorResponse(c, err, nil)
	}

	res, err := h.BlogService.GetByID(c.UserContext(), uint64(id))
	if err != nil {
		err = merry.Wrap(err)
		return utils.ReturnErrorResponse(c, err, nil)
	}

	return utils.ReturnSuccessResponse(c, fiber.StatusOK, "Success", res)
}

// StoreBlog
// @summary Save blog to database
// @tags blog
// @produce json
// @param payload body request.CreateBlogRequest true "JSON payload"
// @router /api/v1/blogs [post]
// @success 201 {object} utils.JSONResponse{data=transformer.BlogTransformer}
// @success 422
// @Security Bearer
func (h *BlogHandler) Store(c *fiber.Ctx) error {
	var req request.CreateBlogRequest
	errorField, err := h.Core.Validator.Validate(c, &req)
	if err != nil {
		err = merry.Wrap(err, merry.WithHTTPCode(fiber.StatusUnprocessableEntity))
		return utils.ReturnErrorResponse(c, err, errorField)
	}

	res, err := h.BlogService.Create(c.UserContext(), req)
	if err != nil {
		err = merry.Wrap(err)
		return utils.ReturnErrorResponse(c, err, nil)
	}

	return utils.ReturnSuccessResponse(c, fiber.StatusCreated, "Success", res)
}

// UpdateBlog
// @summary Update existing blog by ID
// @tags blog
// @produce json
// @param id path int true "blog id"
// @param payload body request.UpdateBlogRequest true "JSON payload"
// @router /api/v1/blogs/{id} [put]
// @success 200 {object} utils.JSONResponse{data=transformer.BlogTransformer}
// @success 404
// @success 422
// @Security Bearer
func (h *BlogHandler) Update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		err = merry.Wrap(err, merry.WithHTTPCode(fiber.StatusBadRequest))
		return utils.ReturnErrorResponse(c, err, nil)
	}

	var req request.UpdateBlogRequest
	errorField, err := h.Core.Validator.Validate(c, &req)
	if err != nil {
		err = merry.Wrap(err, merry.WithHTTPCode(fiber.StatusUnprocessableEntity))
		return utils.ReturnErrorResponse(c, err, errorField)
	}

	res, err := h.BlogService.Update(c.UserContext(), uint64(id), req)
	if err != nil {
		err = merry.Wrap(err)
		return utils.ReturnErrorResponse(c, err, nil)
	}

	return utils.ReturnSuccessResponse(c, fiber.StatusOK, "Success", res)
}

// DeleteBlog
// @summary Delete existing blog by ID
// @tags blog
// @produce json
// @param id path int true "blog id"
// @router /api/v1/blogs/{id} [delete]
// @success 200 {object} utils.JSONResponse
// @success 404
// @Security Bearer
func (h *BlogHandler) Destroy(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		err = merry.Wrap(err, merry.WithHTTPCode(fiber.StatusBadRequest))
		return utils.ReturnErrorResponse(c, err, nil)
	}

	err = h.BlogService.Delete(c.UserContext(), uint64(id))
	if err != nil {
		err = merry.Wrap(err)
		return utils.ReturnErrorResponse(c, err, nil)
	}

	return utils.ReturnSuccessResponse(c, fiber.StatusOK, "Success", nil)
}
