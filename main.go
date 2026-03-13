package main 
import(
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go-api/routes"
	"go-api/config"
)
func main(){
	godotenv.Load()

	config.ConnectDB()

	app := fiber.New()

	routes.SetupRoutes(app)
	
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("BonSOIRRRRR !")
	})

	app.Listen(":3000")
}
