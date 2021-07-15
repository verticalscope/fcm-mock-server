package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type Message struct {
	To           string            `json:"to"`
	Notification Notification      `json:"notification"`
	Data         map[string]string `json:"data"`
	Time         string            `json:"time"`
}

type Notification struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

var messages []Message

func main() {
	app := fiber.New()

	app.Post("/send", send)
	app.Delete("/messages", clear)
	app.Get("/messages", log)

	app.Listen(":4004")
}

func send(c *fiber.Ctx) error {
	auth := c.Get("Authorization")
	if len(auth) == 0 {
		return c.SendStatus(401)
	}
	message := new(Message)
	if err := c.BodyParser(message); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	currentTime := time.Now()
	message.Time = currentTime.Format("2006-01-02 15:04:05")
	messages = append(messages, *message)
	return c.SendStatus(200)
}

func log(c *fiber.Ctx) error {
	return c.JSON(messages)
}

func clear(c *fiber.Ctx) error {
	messages = nil
	return c.SendStatus(200)
}
