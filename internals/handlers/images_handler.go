package handlers

import "github.com/gofiber/fiber/v2"

type ImagesHandler struct {
}

func NewImagesHandler() ImagesHandler {
	return ImagesHandler{}
}

func (h ImagesHandler) GetImage(ctx *fiber.Ctx) error {
	return nil
}
