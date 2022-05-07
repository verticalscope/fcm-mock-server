package main

import (
	"sort"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

type MessageV1 struct {
	Message Message `json:"Message"`
}

type Message struct {
	Authorization string            `json:"authorization,omitempty"`
	To            string            `json:"to,omitempty"`
	Token         string            `json:"token,omitempty"`
	Topic         string            `json:"topic,omitempty"`
	Condition     string            `json:"condition,omitempty"`
	Project       string            `json:"project,omitempty"`
	Notification  Notification      `json:"notification"`
	Data          map[string]string `json:"data,omitempty"`
	Time          string            `json:"time"`
}

type Notification struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

var messages []Message

func main() {
	app := fiber.New()

	app.Post("/send", send)
	app.Post("/v1/projects/:project_id/messages\\:send", sendV1)
	app.Delete("/api/messages", clear)
	app.Get("/api/messages", log)

	app.Listen(":4004")
}

func send(c *fiber.Ctx) error {
	auth := c.Get("Authorization")
	if len(auth) == 0 || !strings.HasPrefix(auth, "key=") {
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

func sendV1(c *fiber.Ctx) error {
	auth := c.Get("Authorization")
	if len(auth) == 0 || !strings.HasPrefix(auth, "Bearer") {
		return c.SendStatus(401)
	}
	messageV1 := new(MessageV1)
	if err := c.BodyParser(messageV1); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	message := messageV1.Message
	currentTime := time.Now()
	message.Time = currentTime.Format("2006-01-02 15:04:05")
	message.Project = c.Params("project_id")
	messages = append(messages, message)
	return c.SendStatus(200)
}

func log(c *fiber.Ctx) error {
	sort.SliceStable(messages, func(i, j int) bool {
		return messages[i].Time > messages[j].Time
	})

	return c.JSON(messages)
}

func clear(c *fiber.Ctx) error {
	messages = nil
	return c.SendStatus(200)
}
