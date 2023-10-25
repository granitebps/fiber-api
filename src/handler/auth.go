package handler

import (
	"github.com/ansel1/merry/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/granitebps/fiber-api/pkg/core"
	"github.com/granitebps/fiber-api/pkg/utils"
	"github.com/granitebps/fiber-api/src/request"
	"github.com/granitebps/fiber-api/src/service"
)

type AuthHandler struct {
	Core *core.Core

	AuthService *service.AuthService
}

func NewAuthHandler(c *core.Core, authService *service.AuthService) *AuthHandler {
	return &AuthHandler{
		Core:        c,
		AuthService: authService,
	}
}

// Register
// @summary Register a new user
// @tags auth
// @produce json
// @param payload body request.RegisterRequest true "JSON payload"
// @router /api/v1/auth/register [post]
// @success 201 {object} utils.JSONResponse{data=transformer.RegisterTransformer}
// @success 422
func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req request.RegisterRequest
	errorField, err := h.Core.Validator.Validate(c, &req)
	if err != nil {
		err = merry.Wrap(err, merry.WithHTTPCode(fiber.StatusUnprocessableEntity))
		return utils.ReturnErrorResponse(c, err, errorField)
	}

	res, err := h.AuthService.Register(c.UserContext(), req)
	if err != nil {
		err = merry.Wrap(err)
		return utils.ReturnErrorResponse(c, err, nil)
	}

	return utils.ReturnSuccessResponse(c, fiber.StatusOK, "Success", res)
}

// Login
// @summary Login existing user
// @tags auth
// @produce json
// @param payload body request.LoginRequest true "JSON payload"
// @router /api/v1/auth/login [post]
// @success 201 {object} utils.JSONResponse{data=transformer.LoginTransformer}
// @success 422
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req request.LoginRequest
	errorField, err := h.Core.Validator.Validate(c, &req)
	if err != nil {
		err = merry.Wrap(err, merry.WithHTTPCode(fiber.StatusUnprocessableEntity))
		return utils.ReturnErrorResponse(c, err, errorField)
	}

	res, err := h.AuthService.Login(c.UserContext(), req)
	if err != nil {
		err = merry.Wrap(err)
		return utils.ReturnErrorResponse(c, err, nil)
	}

	return utils.ReturnSuccessResponse(c, fiber.StatusOK, "Success", res)
}

// Me
// @summary Get authenticated user
// @tags auth
// @produce json
// @router /api/v1/auth/me [get]
// @success 201 {object} utils.JSONResponse{data=transformer.MeTransformer}
// @success 401
// @Security Bearer
func (h *AuthHandler) Me(c *fiber.Ctx) error {
	res, err := h.AuthService.Me(c.UserContext(), utils.GetAuthUser(c))
	if err != nil {
		err = merry.Wrap(err)
		return utils.ReturnErrorResponse(c, err, nil)
	}

	return utils.ReturnSuccessResponse(c, fiber.StatusOK, "Success", res)
}
