package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/google/uuid"
)

type Payload struct {
	Project     string         `json:"project"`
	Time        time.Time      `json:"time"`
	Name        string         `json:"name"`
	RequestBody FCMRequestBody `json:"requestBody"`
}

type FCMRequestBody struct {
	Message      map[string]interface{} `json:"message"`
	ValidateOnly bool                   `json:"validate_only,omitempty"`
}

var payloads []Payload

func main() {
	payloads = []Payload{}
	app := fiber.New()

	// Default middleware config
	app.Use(logger.New())

	app.Post("/v1/projects/:project_id/messages\\:send", sendV1)
	app.Delete("/api/messages", deleteMessages)
	app.Get("/api/messages", getPayloads)

	app.Listen(":4004")
}

func sendV1(c *fiber.Ctx) error {
	auth := c.Get("Authorization")
	if len(auth) == 0 || !strings.HasPrefix(auth, "Bearer") {
		return c.SendStatus(401)
	}
	requestBody := new(FCMRequestBody)
	if err := c.BodyParser(requestBody); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	projectID := c.Params("project_id")
	name := fmt.Sprintf("projects/%s/messages/%s", projectID, uuid.New())

	payload := Payload{
		Time:        time.Now(),
		Project:     projectID,
		Name:        name,
		RequestBody: *requestBody,
	}

	payloads = append(payloads, payload)

	responseBody := map[string]interface{}{
		"name": name,
	}

	return c.Status(200).JSON(responseBody)
}

func getPayloads(c *fiber.Ctx) error {
	return c.JSON(payloads)
}

func deleteMessages(c *fiber.Ctx) error {
	payloads = []Payload{}
	return c.SendStatus(200)
}
