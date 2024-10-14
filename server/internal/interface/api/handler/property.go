package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/gotired/real-estate-admin/server/internal/domain/model/property"
	"github.com/gotired/real-estate-admin/server/internal/domain/service"
	"github.com/gotired/real-estate-admin/server/internal/interface/api/response"
)

type Property struct {
	PropertyService *service.Property
	UserService     *service.User
}

func (p *Property) List(c *fiber.Ctx) error {
	var listParams property.ListParams

	if err := c.QueryParser(&listParams); err != nil {
		return response.Error(c, 400, err.Error())
	}

	if listParams.Skip == nil {
		defaultSkip := 0
		listParams.Skip = &defaultSkip
	}
	if listParams.Limit == nil {
		defaultLimit := 10
		listParams.Limit = &defaultLimit
	}
	if listParams.Order == nil {
		defaultOrder := 1
		listParams.Order = &defaultOrder
	}
	if listParams.SortKey == nil {
		defaultSortKey := "title"
		listParams.SortKey = &defaultSortKey
	}
	listProperties, err := p.PropertyService.List(listParams.Skip, listParams.Limit, listParams.Order, listParams.Keyword, listParams.SortKey, listParams.Type)

	if err != nil {
		return response.Error(c, 500, err.Error())
	}

	return response.Success(c, 200, listProperties, "")
}

func (p *Property) Get(c *fiber.Ctx) error {
	propertyID := c.Params("id")

	propertyUUID, err := uuid.Parse(propertyID)
	if err != nil {
		return response.Error(c, 400, err.Error())
	}

	propertyDetail, err := p.PropertyService.GetByID(propertyUUID)
	if err != nil {
		return response.Error(c, 500, err.Error())
	}

	if propertyDetail == nil {
		return response.Error(c, 404, fmt.Sprintf("property id %s not found", propertyID))
	}

	return response.Success(c, 200, propertyDetail, "")
}

func (p *Property) Update(c *fiber.Ctx) error {
	propertyID := c.Params("id")

	propertyUUID, err := uuid.Parse(propertyID)
	if err != nil {
		return response.Error(c, 400, err.Error())
	}

	var updateDetail property.Update
	if err := c.BodyParser(&updateDetail); err != nil {
		return response.Error(c, 400, err.Error())
	}

	if err = p.PropertyService.Update(propertyUUID, updateDetail.Owner, updateDetail.Title, updateDetail.Description, updateDetail.Type, updateDetail.Location, updateDetail.Price, updateDetail.Photo); err != nil {
		return response.Error(c, 500, err.Error())
	}

	return response.Success(c, 200, nil, "update property successfully")
}

func (p *Property) Delete(c *fiber.Ctx) error {
	propertyID := c.Params("id")

	propertyUUID, err := uuid.Parse(propertyID)
	if err != nil {
		return response.Error(c, 500, err.Error())
	}

	if err = p.PropertyService.Delete(propertyUUID); err != nil {
		return response.Error(c, 500, err.Error())
	}

	return response.Success(c, 200, nil, "delete property successfully")
}

func (p *Property) Create(c *fiber.Ctx) error {
	var createDetail property.Create
	if err := c.BodyParser(&createDetail); err != nil {
		return response.Error(c, 400, err.Error())
	}

	userID, err := p.UserService.GetByEmail(createDetail.Email)
	if err != nil {
		return response.Error(c, 500, err.Error())
	}

	if userID == nil {
		return response.Error(c, 404, fmt.Sprintf("email %s not found", createDetail.Email))
	}

	if err = p.PropertyService.Create(*userID, createDetail.Title, createDetail.Description, createDetail.Type, createDetail.Location, createDetail.Photo, createDetail.Price); err != nil {
		return response.Error(c, 500, err.Error())
	}

	return response.Success(c, 200, nil, "create property successfully")
}
