package routes
import(
	"github.com/gofiber/fiber/v2"
	"go-api/handlers"

)
func SetupRoutes(app *fiber.App) {
	app.Post("/register", handlers.Register)
	app.Post("/login", handlers.Login)
}