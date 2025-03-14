package main

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

type Language struct {
	Language       string   `json:"language"`
	Appeared       int      `json:"appeared"`
	Created        []string `json:"created"`
	Functional     bool     `json:"functional"`
	ObjectOriented bool     `json:"object-oriented"`
	Relation       struct {
		InfluencedBy []string `json:"influenced-by"`
		Influences   []string `json:"influences"`
	} `json:"relation"`
}

var languages []Language

func isPalindrome(text string) bool {
	text = strings.ToLower(strings.ReplaceAll(text, " ", ""))
	reversed := ""
	for i := len(text) - 1; i >= 0; i-- {
		reversed += string(text[i])
	}
	return text == reversed
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello Go developers")
	})

	app.Get("/language", func(c *fiber.Ctx) error {
		data := Language{
			Language:       "C",
			Appeared:       1972,
			Created:        []string{"Dennis Ritchie"},
			Functional:     true,
			ObjectOriented: false,
		}
		data.Relation.InfluencedBy = []string{"B", "ALGOL 68", "Assembly", "FORTRAN"}
		data.Relation.Influences = []string{"C++", "Objective-C", "C#", "Java", "JavaScript", "PHP", "Go"}
		return c.JSON(data)
	})

	app.Get("/palindrome", func(c *fiber.Ctx) error {
		text := c.Query("text")
		if text == "" {
			return c.Status(fiber.StatusBadRequest).SendString("Text parameter is required")
		}
		if isPalindrome(text) {
			return c.SendString("Palindrome")
		}
		return c.Status(fiber.StatusBadRequest).SendString("Not palindrome")
	})

	app.Get("/languages", func(c *fiber.Ctx) error {
		return c.JSON(languages)
	})

	app.Get("/language/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil || id < 0 || id >= len(languages) {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
		}
		return c.JSON(languages[id])
	})

	app.Post("/language", func(c *fiber.Ctx) error {
		var lang Language
		if err := c.BodyParser(&lang); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
		}
		languages = append(languages, lang)
		return c.Status(fiber.StatusCreated).JSON(lang)
	})

	app.Patch("/language/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil || id < 0 || id >= len(languages) {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
		}
		var update Language
		if err := c.BodyParser(&update); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
		}
		languages[id] = update
		return c.JSON(update)
	})

	app.Delete("/language/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil || id < 0 || id >= len(languages) {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
		}
		languages = append(languages[:id], languages[id+1:]...)
		return c.SendString("Deleted successfully")
	})

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Method not allowed")
	})

	app.Listen(":3000")
}
