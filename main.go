package main

import (
	"goFiberMongoCrud/handlers"

	"github.com/gofiber/fiber/v2")


func main() {
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World 👋!")
    })
	app.Get("/getColleges",handlers.GetColleges)

    app.Listen(":3000")
}