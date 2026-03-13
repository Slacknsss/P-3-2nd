package handlers

import (
    "github.com/gofiber/fiber/v2"
    "go-api/config"
    "go-api/models"
)

func AddTodo(c *fiber.Ctx) error {
    input := new(struct {
        Title     string `json:"title"`
        Completed bool   `json:"completed"`
    })

    c.BodyParser(input)

    todo := models.Todo{
        Title:     input.Title,
        Completed: input.Completed,
        UserID:    1,
    }

    config.DB.Create(&todo)

    return c.Status(201).JSON(fiber.Map{"message": "Todo créé !"})
}

func DeleteTodo(c *fiber.Ctx) error {
    id := c.Params("id")
    config.DB.Delete(&models.Todo{}, id)
    return c.JSON(fiber.Map{"message": "Todo supprimé !"})
}