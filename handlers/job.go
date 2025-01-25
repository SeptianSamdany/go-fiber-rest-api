package handlers

import (
    "github.com/SeptianSamdany/go-fiber-rest-api/entities"
    "github.com/gofiber/fiber/v2"
    "gorm.io/gorm"
)

var database *gorm.DB // Pastikan Anda menginisialisasi database di tempat lain

// Tambahkan Job Baru
func AddJob(c *fiber.Ctx) error {
    var job entities.Job

    if err := c.BodyParser(&job); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Failed to parse job data",
        })
    }

    if err := database.Create(&job).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to create job",
        })
    }

    return c.JSON(job)
}

// Ambil Semua Job
func GetJobs(c *fiber.Ctx) error {
    var jobs []entities.Job

    if err := database.Preload("Users").Find(&jobs).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to retrieve jobs",
        })
    }

    return c.JSON(jobs)
}

