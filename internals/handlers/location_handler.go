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
	locations, err := h.locationService.GetLocation()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(locations)
	return ctx.SendString("get locations")
}

func (h LocationHandler) GetLocationsAutoComplete(ctx *fiber.Ctx) error {
	term := ctx.Query("term")
	matchedOptions := h.locationService.GetLocationsAutoComplete(term)
	return ctx.JSON(matchedOptions)
}
