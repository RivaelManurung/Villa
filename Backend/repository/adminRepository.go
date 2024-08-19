package repository

import (
	"Backend/models/dto"
	"Backend/models/entity"
	"Backend/database"
	"Backend/utils"
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type AdminRepository interface {
	AdminLogin(ctx *fiber.Ctx, input dto.RequestAdminLogin) (*entity.Admin, error)
	GetProfile(ctx *fiber.Ctx) (*entity.Admin, error)
}

type adminRepository struct{}

func NewAdminRepository() AdminRepository {
	return &adminRepository{}
}

func (r *adminRepository) AdminLogin(ctx *fiber.Ctx, input dto.RequestAdminLogin) (*entity.Admin, error) {
	validation := validator.New()

	if err := validation.Struct(input); err != nil {
		return nil, err
	}

	var admin entity.Admin
	result := database.DB.First(&admin, "username = ?", input.Username)
	if result.Error != nil {
		return nil, result.Error
	}

	if !utils.CheckPassword(input.Password, admin.Password) {
		return nil, errors.New("password not match")
	}

	return &admin, nil
}

func (r *adminRepository) GetProfile(ctx *fiber.Ctx) (*entity.Admin, error) {
	admin := ctx.Locals("admin").(entity.Admin)
	return &admin, nil
}
