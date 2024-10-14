package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/gotired/real-estate-admin/server/internal/domain/model/user"
	"github.com/gotired/real-estate-admin/server/internal/domain/service"
	"github.com/gotired/real-estate-admin/server/internal/interface/api/response"
)

type User struct {
	UserService     *service.User
	PropertyService *service.Property
}

func (u *User) List(c *fiber.Ctx) error {
	skip, err := c.ParamsInt("skip", 0)
	if err != nil {
		return response.Error(c, 400, err.Error())
	}
	limit, err := c.ParamsInt("limit", 10)
	if err != nil {
		return response.Error(c, 400, err.Error())
	}

	if limit < 0 {
		return response.Error(c, 400, "parameter 'limit' must be greater than 0")
	}
	if skip < 0 {
		return response.Error(c, 400, "parameter 'skip' must be greater than 0")
	}

	users, err := u.UserService.List(&skip, &limit)
	if err != nil {
		return response.Error(c, 500, err.Error())
	}
	return response.Success(c, 200, users, "")
}

func (u *User) Get(c *fiber.Ctx) error {
	userID := c.Params("id")

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return response.Error(c, 400, err.Error())
	}

	userDetails, err := u.UserService.Get(userUUID)
	if err != nil {
		return response.Error(c, 500, err.Error())
	}
	userProperty, err := u.PropertyService.GetByOwnerID(userUUID)
	if err != nil {
		return response.Error(c, 500, err.Error())
	}
	userDetails.Property = userProperty

	return response.Success(c, 200, userDetails, "")
}

func (u *User) Create(c *fiber.Ctx) error {
	var body user.Create

	if err := c.BodyParser(&body); err != nil {
		return response.Error(c, 400, err.Error())
	}

	if err := u.UserService.Create(body.Name, body.Email, body.Avatar); err != nil {
		return response.Error(c, 500, err.Error())
	}
	return response.Success(c, 201, nil, "User created successfully")
}
