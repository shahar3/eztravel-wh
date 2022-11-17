package routes

import (
	"eztravel-wh/internals/handlers"
	"github.com/gofiber/fiber/v2"
)

type ImagesRouteController struct {
	imagesHandler handlers.ImagesHandler
}

func NewImagesRouteController(imagesHandler handlers.ImagesHandler) ImagesRouteController {
	return ImagesRouteController{imagesHandler: imagesHandler}
}

func (irc ImagesRouteController) ImagesRoute(router fiber.Router) {
	imagesRouter := router.Group("/images")

	imagesRouter.Get("/", irc.imagesHandler.GetImage)
}
