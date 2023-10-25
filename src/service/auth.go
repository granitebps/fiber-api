package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/ansel1/merry/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/granitebps/fiber-api/pkg/constants"
	"github.com/granitebps/fiber-api/pkg/core"
	"github.com/granitebps/fiber-api/pkg/utils"
	"github.com/granitebps/fiber-api/src/model"
	"github.com/granitebps/fiber-api/src/repository"
	"github.com/granitebps/fiber-api/src/request"
	"github.com/granitebps/fiber-api/src/transformer"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	Core *core.Core

	UserRepo *repository.UserRepository
}

func NewAuthService(c *core.Core, userRepo *repository.UserRepository) *AuthService {
	return &AuthService{
		Core:     c,
		UserRepo: userRepo,
	}
}

func (s *AuthService) Register(ctx context.Context, req request.RegisterRequest) (res transformer.RegisterTransformer, err error) {
	_, err = s.UserRepo.GetByEmail(ctx, req.Email)
	if err == nil {
		err = merry.New("duplicate email", merry.WithHTTPCode(fiber.StatusUnprocessableEntity), merry.WithUserMessage("Email already registered."))
		return
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		err = merry.Wrap(err)
		return
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		err = merry.Wrap(err)
		return
	}

	user := model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(pass),
	}

	err = s.UserRepo.Insert(ctx, &user)
	if err != nil {
		err = merry.Wrap(err)
		return
	}

	t, err := s.generateToken(user)
	if err != nil {
		err = merry.Wrap(err)
		return
	}

	res.Token = t
	return
}

func (s *AuthService) Login(ctx context.Context, req request.LoginRequest) (res transformer.LoginTransformer, err error) {
	user, err := s.UserRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		err = merry.Wrap(err, merry.WithHTTPCode(fiber.StatusUnprocessableEntity), merry.WithUserMessage("Wrong Email/Password."))
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		err = merry.Wrap(err, merry.WithHTTPCode(fiber.StatusUnprocessableEntity), merry.WithUserMessage("Wrong Email/Password."))
		return
	}

	t, err := s.generateToken(user)
	if err != nil {
		err = merry.Wrap(err)
		return
	}

	res.Token = t
	return
}

func (s *AuthService) Me(ctx context.Context, user utils.AuthUser) (res transformer.MeTransformer, err error) {
	res.ID = user.ID
	res.Name = user.Name
	res.Email = user.Email

	return
}

func (s *AuthService) generateToken(u model.User) (t string, err error) {
	claims := jwt.MapClaims{
		"id":    fmt.Sprint(u.ID),
		"name":  u.Name,
		"email": u.Email,
		"exp":   time.Now().Add(time.Minute * constants.JWT_DEFAULT_TTL).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err = token.SignedString([]byte(viper.GetString(constants.JWT_SECRET)))
	if err != nil {
		err = merry.Wrap(err)
		return
	}

	return
}
