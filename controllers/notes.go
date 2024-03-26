package controllers
import (
	"github.com/gofiber/fiber/v2"
	models "notetaking/models"
)
func GetNotes(c *fiber.Ctx) error {
	var notes []models.Note
    if err := models.DB.Find(&notes).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch notes"})
    }
    return c.Status(fiber.StatusOK).JSON(notes)
}
