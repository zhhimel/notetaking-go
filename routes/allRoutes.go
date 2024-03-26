package routes
import (
	controllers "notetaking/controllers"
	"github.com/gofiber/fiber/v2"
)
func SetRoutes(api fiber.Router){
	api.Get("/notes",controllers.GetNotes)
}