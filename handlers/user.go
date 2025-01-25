package handlers

import (
    "github.com/SeptianSamdany/go-fiber-rest-api/config"
    "github.com/SeptianSamdany/go-fiber-rest-api/entities"
    "github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error {
    var users []entities.User

    if err := config.DB.Preload("Job").Find(&users).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to retrieve users",
        })
    }

    return c.JSON(users)
}

func AddUser(c *fiber.Ctx) error {
    var user entities.User

    if err := c.BodyParser(&user); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Failed to parse user data",
        })
    }

    var job entities.Job
    if err := config.DB.First(&job, user.JobID).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Job not found",
        })
    }

    if err := config.DB.Create(&user).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to create user",
        })
    }

    return c.JSON(user)
}
