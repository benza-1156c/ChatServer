package main

import (
	"chatserver/modules/websocket/hub"
	"chatserver/pkg/database"
	"chatserver/routes"
	"fmt"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	db := database.ConnectPosgres()
	h := hub.NewHub()
	go h.Run()

	cld, err := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to Cloudinary")

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowMethods:     "GET,POST,PUT,DELETE,PATCH",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}))

	routes.SetUpRoutes(app, h, db, cld)

	if err := app.Listen(":80"); err != nil {
		panic(err)
	}
}
