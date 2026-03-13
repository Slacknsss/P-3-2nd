package handlers
import(
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"go-api/config"
	"go-api/models"
	"os"
	"github.com/golang-jwt/jwt/v4"
)
func Register(c *fiber.Ctx) error {
	input := new(struct {
		Username string `json:"username"`
		Email    string `json:"email"`
        Password string `json:"password"`
	})

    if err := c.BodyParser(input); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Données invalides"})	
    }

    hash, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 10)

    user := models.User{
        Username: input.Username,
        Email:    input.Email,
        Password: string(hash),
    }

    config.DB.Create(&user)

    return c.Status(201).JSON(fiber.Map{"message": "Compte créé !"})
}

func Login(c *fiber.Ctx) error {
    input := new(struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    })

    c.BodyParser(input)

    var user models.User
    config.DB.Where("email = ?", input.Email).First(&user)

    token := jwt.New(jwt.SigningMethodHS256)
    claims := token.Claims.(jwt.MapClaims)
    claims["user_id"] = user.ID
    t, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

    return c.JSON(fiber.Map{"token": t})
}