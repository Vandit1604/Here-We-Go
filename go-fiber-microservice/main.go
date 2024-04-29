package main

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	setupRoutes(app)

	app.Listen(":3000")
	log.Println("Check localhost:3000")
}

// ---------- types.go

type Blog struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type blogPosts []Blog

var blogs = blogPosts{
	{ID: 0, Title: "Hello", Content: "Hello World"},
	{ID: 1, Title: "Fiber", Content: "Go Fiber"},
	{ID: 2, Title: "Vandit", Content: "Vandit Singh"},
}

// ---------- rotues.go

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/post", GetPosts)
	app.Get("/api/v1/post/:id", GetPost)
}

func GetPosts(c *fiber.Ctx) error {
	return c.JSON(blogs)
}

func GetPost(c *fiber.Ctx) error {
	id := c.Params("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(fiber.Map{"error": "Invalid ID"})
	}
	for _, s := range blogs {
		if s.ID == i {
			return c.JSON(s)
		}
	}
	return c.JSON(fiber.Map{"error": "Post not found"})
}
