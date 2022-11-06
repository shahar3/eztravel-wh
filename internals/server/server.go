package server

import (
	"context"
	"eztravel-wh/internals/handlers"
	"eztravel-wh/internals/routes"
	"eztravel-wh/internals/services"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

var (
	locationService services.LocationService
	locationHandler handlers.LocationHandler
	locationRouter  routes.LocationRouteController
)

type Server struct {
	mongoClient *mongo.Client
}

func NewServer(client *mongo.Client) Server {
	return Server{mongoClient: client}
}

func (s Server) Init() {
	ctx := context.TODO()

	mongoDb := s.mongoClient.Database("eztravel")
	locationCollection := mongoDb.Collection("locations")

	locationService = services.NewLocationService(locationCollection, ctx)
	locationHandler = handlers.NewLocationHandler(locationService)
	locationRouter = routes.NewLocationRouteController(locationHandler)

	locationService.Init()
}

func (s Server) StartServer() {
	port := os.Getenv("PORT")

	s.Init()
	app := fiber.New()
	api := app.Group("/api")

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("ok")
	})

	locationRouter.LocationRoute(api)
	app.Listen(fmt.Sprintf(":%s", port))
}
