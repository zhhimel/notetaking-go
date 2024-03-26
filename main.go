package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
    "os"
    routes "notetaking/routes"
)

func main() {
    app := fiber.New()
    app.Use(cors.New())
		app.Route("/api",routes.SetRoutes )
		port := os.Getenv("PORT")
    if port == "" {
        port = "4000"
    }
		app.Listen(":" + port)
}
