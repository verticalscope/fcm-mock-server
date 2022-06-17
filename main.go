package main

import (
	"sort"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type MessageV1 struct {
	Message      Message `json:"Message"`
	ValidateOnly bool    `json:"validate_only,omitempty"`
	Project      string  `json:"project,omitempty"`
	Time         string  `json:"time"`
}

type Message struct {
	Authorization string            `json:"authorization,omitempty"`
	To            string            `json:"to,omitempty"`
	Token         string            `json:"token,omitempty"`
	Topic         string            `json:"topic,omitempty"`
	Condition     string            `json:"condition,omitempty"`
	Notification  Notification      `json:"notification"`
	Data          map[string]string `json:"data,omitempty"`
}

type Notification struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

var messages []MessageV1

func main() {
	messages = []MessageV1{}
	app := fiber.New()

	// Default middleware config
	app.Use(logger.New())

	app.Post("/v1/projects/:project_id/messages\\:send", sendV1)
	app.Delete("/api/messages", deleteMessages)
	app.Get("/api/messages", getMessages)

	app.Listen(":4004")
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
	currentTime := time.Now()
	messageV1.Time = currentTime.Format("2006-01-02 15:04:05")
	messageV1.Project = c.Params("project_id")
	messages = append(messages, *messageV1)
	return c.Status(200).SendString("{}")
}

func getMessages(c *fiber.Ctx) error {
	sort.SliceStable(messages, func(i, j int) bool {
		return messages[i].Time > messages[j].Time
	})

	return c.JSON(messages)
}

func deleteMessages(c *fiber.Ctx) error {
	messages = []MessageV1{}
	return c.SendStatus(200)
}
