package handlers

import (
	"eztravel-wh/internals/services"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type LocationHandler struct {
	locationService services.LocationService
}

func NewLocationHandler(locationService services.LocationService) LocationHandler {
	return LocationHandler{locationService: locationService}
}

func (h LocationHandler) GetLocation(ctx *fiber.Ctx) error {
	mongoIdHex := ctx.Query("hex")
	location, err := h.locationService.GetLocation(mongoIdHex)
	if err != nil {
		fmt.Println(err)
	}

	return ctx.JSON(location)
}

func (h LocationHandler) GetLocationsAutoComplete(ctx *fiber.Ctx) error {
	term := ctx.Query("term")
	matchedOptions := h.locationService.GetLocationsAutoComplete(term)
	return ctx.JSON(matchedOptions)
}
