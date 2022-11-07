package routes

import (
	"eztravel-wh/internals/handlers"
	"github.com/gofiber/fiber/v2"
)

type LocationRouteController struct {
	locationHander handlers.LocationHandler
}

func NewLocationRouteController(locationHandler handlers.LocationHandler) LocationRouteController {
	return LocationRouteController{
		locationHander: locationHandler,
	}
}

func (lrc LocationRouteController) LocationRoute(router fiber.Router) {
	locationRouter := router.Group("/location")

	locationRouter.Get("/autocomplete", lrc.locationHander.GetLocationsAutoComplete)
	locationRouter.Get("/:hex", lrc.locationHander.GetLocation)
	locationRouter.Get("/", lrc.locationHander.GetLocations)
}
