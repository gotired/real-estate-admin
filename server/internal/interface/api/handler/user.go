package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/gotired/real-estate-admin/server/internal/domain/model/user"
	"github.com/gotired/real-estate-admin/server/internal/domain/service"
	"github.com/gotired/real-estate-admin/server/internal/interface/api/response"
)

type User struct {
	Service *service.User
}

func (u *User) List(c *fiber.Ctx) error {
	users, err := u.Service.List()
	if err != nil {
		return response.Error(c, 500, err.Error())
	}
	return response.Success(c, 200, users)
}

func (u *User) Get(c *fiber.Ctx) error {
	userID := c.Params("id")

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return response.Error(c, 400, err.Error())
	}

	userDetails, err := u.Service.Get(&userUUID)
	if err != nil {
		return response.Error(c, 500, err.Error())
	}
	return response.Success(c, 200, userDetails)
}

func (u *User) Create(c *fiber.Ctx) error {
	var body user.Create

	if err := c.BodyParser(&body); err != nil {
		return response.Error(c, 400, err.Error())
	}

	if err := u.Service.Create(&body.Name, &body.Email, &body.Avatar); err != nil {
		return response.Error(c, 500, err.Error())
	}
	return response.Success(c, 201, "User created successfully")
}
