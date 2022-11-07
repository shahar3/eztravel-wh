package handlers

import (
	"eztravel-wh/internals/services"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strings"
)

type LocationHandler struct {
	locationService services.LocationService
}

func NewLocationHandler(locationService services.LocationService) LocationHandler {
	return LocationHandler{locationService: locationService}
}

func (h LocationHandler) GetLocation(ctx *fiber.Ctx) error {
	mongoIdHex := ctx.Params("hex")
	location, err := h.locationService.GetLocation(mongoIdHex)
	if err != nil {
		return ctx.JSON(map[string]string{"error": err.Error()})
	}

	return ctx.JSON(location)
}

func (h LocationHandler) GetLocations(ctx *fiber.Ctx) error {
	locations := h.locationService.GetLocations()

	return ctx.JSON(locations)
}

func (h LocationHandler) GetLocationsAutoComplete(ctx *fiber.Ctx) error {
	term := ctx.Query("term")
	if strings.Trim(term, " ") == "" {
		return ctx.JSON(map[string]string{"error": fmt.Sprintf("term %s is not valid", term)})
	}
	matchedOptions := h.locationService.GetLocationsAutoComplete(term)
	return ctx.JSON(matchedOptions)
}
