package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()

	app.Post("/notifications", notificationsHandler)
	app.Get("/healthy", healthyHandler)
	app.Get("/ready", readyHandler)

	log.Fatal(app.Listen(":8888"))
}

func notificationsHandler(ctx *fiber.Ctx) error {
	body := ctx.Body()
	log.Printf("Body: %s", string(body))
	log.Printf("-------")
	return ctx.SendStatus(fiber.StatusAccepted)
}

func healthyHandler(ctx *fiber.Ctx) error {
	return ctx.SendStatus(200)
}

func readyHandler(ctx *fiber.Ctx) error {
	return ctx.SendStatus(200)
}
