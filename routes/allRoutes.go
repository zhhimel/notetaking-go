package routes
import (
	controllers "notetaking/controllers"
	"github.com/gofiber/fiber/v2"
)
func SetRoutes(api fiber.Router){
	api.Get("/notes",controllers.GetNotes)
	api.Post("/createnote",controllers.CreateNote)
	api.Delete("/notes/:id", controllers.DeleteNote)
	api.Put("/notes/:id", controllers.UpdateNote)
}