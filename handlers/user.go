package handlers

import (
	"github.com/SeptianSamdany/go-fiber-rest-api/entities"
	"github.com/SeptianSamdany/go-fiber-rest-api/config"
	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error {
    var users []entities.User

    if err := database.Preload("Job").Find(&users).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to retrieve users",
        })
    }

    return c.JSON(users)
}

func GetUserById(c *fiber.Ctx) error {
	id := c.Params("id")
	var user entities.User

	result := config.Database.Find(&user, id)

	if result.RowsAffected == 0 {
		return c.Status(404).JSON(map[string] string {
			"message": "User not found", 
		})
	}

	return c.Status(200).JSON(user)
}

func AddUser(c *fiber.Ctx) error {
    var user entities.User

    if err := c.BodyParser(&user); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Failed to parse user data",
        })
    }

    // Validasi JobID
    var job entities.Job
    if err := database.First(&job, user.JobID).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Job not found",
        })
    }

    if err := database.Create(&user).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to create user",
        })
    }

    return c.JSON(user)
}

 
func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := new(entities.User)
   
	if err := c.BodyParser(user); err != nil {
	 return c.Status(503).SendString(err.Error())
	}
   
	config.Database.Where("id = ?", id).Updates(&user)
	return c.Status(200).JSON(user)
}
   
func DeleteUserById(c *fiber.Ctx) error {
	id := c.Params("id")

	var user entities.User

	result := config.Database.Delete(&user, id)

	if result.RowsAffected == 0 {
		return c.Status(404).JSON(map[string]string{
		"message": "Data Uer not found, please check again",
		})
	}

	return c.Status(200).JSON(map[string]string{
		"message": "User success deleted",
	})
}