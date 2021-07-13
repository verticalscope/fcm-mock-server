package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type Message struct {
	To           string       `json:"to"`
	Notification Notification `json:"notification"`
	Data         Data         `json:"data"`
}

type Data struct {
	Click_action string `json:"click_action"`
	Action       string `json:"action"`
	Itemtype     string `json:"itemtype"`
	Value        string `json:"value"`
}

type Notification struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type Activity struct {
	Device_token string `json:"device_token"`
	Title        string `json:"title"`
	Click_action string `json:"click_action"`
	Action       string `json:"action"`
	Itemtype     string `json:"itemtype"`
	Value        string `json:"value"`
	Time         string `json:"time"`
}

var activities []Activity

func main() {
	app := fiber.New()

	app.Post("/send", send)
	app.Put("/clearlog", clear)
	app.Get("/log", log)

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
	activity := Activity{
		Action:       message.Data.Action,
		Device_token: message.To,
		Title:        message.Notification.Title,
		Click_action: message.Data.Click_action,
		Itemtype:     message.Data.Itemtype,
		Value:        message.Data.Value,
		Time:         currentTime.Format("2006-01-02 15:04:05"),
	}
	activities = append(activities, activity)
	return c.SendStatus(200)
}

func log(c *fiber.Ctx) error {
	return c.JSON(activities)
}

func clear(c *fiber.Ctx) error {
	activities = nil
	return c.SendStatus(200)
}
