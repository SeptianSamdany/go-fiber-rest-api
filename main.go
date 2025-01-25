package main

import (
    "github.com/SeptianSamdany/go-fiber-rest-api/config"
    "github.com/SeptianSamdany/go-fiber-rest-api/handlers"
    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    // Koneksi ke database
    config.Connect()

    // Route Job
    app.Post("/job", handlers.AddJob)
    app.Get("/job", handlers.GetJobs)

    // Route User
    app.Get("/user", handlers.GetUser)
    app.Post("/user", handlers.AddUser)
    app.Get("/user/:id", handlers.GetUserById)
    app.Put("/user/:id", handlers.UpdateUser)
    app.Delete("/user/:id", handlers.DeleteUserById)

    app.Listen(":3000")
}
