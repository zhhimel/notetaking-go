package controllers
import (
	
	models "notetaking/models"
	"github.com/gofiber/fiber/v2"
	"strconv"
)
func GetNotes(c *fiber.Ctx) error {
	var notes []models.Note
    if err := models.DB.Find(&notes).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch notes"})
    }
    return c.Status(fiber.StatusOK).JSON(notes)
}
func  CreateNote(c *fiber.Ctx) error{
	note:= new(models.Note)
	if err := c.BodyParser(note);err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Invalid request body"})
	}
	if err := models.DB.Create(note).Error;err!=nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"Failed to create note"})
	}
	return c.Status(fiber.StatusCreated).JSON(note)	
}
func DeleteNote(c *fiber.Ctx) error {
	id := c.Params("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid note ID"})
	}
	if err := models.DB.Delete(&models.Note{}, idUint); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete note"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": id + " Note deleted"})
}

func UpdateNote(c *fiber.Ctx) error {
	id:=c.Params("id")
	idUint,err:=strconv.ParseUint(id,10,64)
	if err!=nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error":"Invalid note id"})
	}
	var existingNote models.Note
	if err=models.DB.First(&existingNote,idUint).Error; err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":id+"No note with the given id"})
	}
	updatedNote:= new(models.Note)
	if err=c.BodyParser(updatedNote); err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Invalid update data"})
	}
	existingNote.Title=updatedNote.Title
	existingNote.Details=  updatedNote.Details
	if err := models.DB.Save(&existingNote).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update note"})
}
	return c.Status(fiber.StatusOK).JSON(existingNote)
}