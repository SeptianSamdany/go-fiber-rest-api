package handlers

import (
    "github.com/SeptianSamdany/go-fiber-rest-api/config"
    "github.com/SeptianSamdany/go-fiber-rest-api/entities"
    "github.com/gofiber/fiber/v2"
)

// Tambahkan Job Baru
func AddJob(c *fiber.Ctx) error {
    var job entities.Job
    if err := c.BodyParser(&job); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
    }

    if err := config.DB.Create(&job).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create job"})
    }

    return c.JSON(job)
}

// Ambil Semua Job
func GetJobs(c *fiber.Ctx) error {
    var jobs []entities.Job
    if err := config.DB.Preload("Users").Find(&jobs).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch jobs"})
    }

    return c.JSON(jobs)
}
