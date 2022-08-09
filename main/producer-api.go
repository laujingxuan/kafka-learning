package main

import (
	"encoding/json"
	"kafka-sarama/kafkaHandler"
	"github.com/gofiber/fiber/v2"
)

type Comment struct {
	Text string `form:"text" json:"text"`
}

func createComment(c *fiber.Ctx) error {
	// Instantiate new Message struct
	cmt := new(Comment)
	if err := c.BodyParser(cmt); err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return err
	}
	// convert body into bytes and send it to kafka
	cmtInBytes, err := json.Marshal(cmt)
	kafkaHandler.PushCommentToQueue("comments", cmtInBytes)
	// Return Comment in JSON format
	err = c.JSON(&fiber.Map{
		"success": true,
		"message": "Comment pushed successfully",
		"comment": cmt,
	})
	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "Error creating product",
		})
		return err
	}
	return err
}

func main() {
	app := fiber.New()
	api := app.Group("/api/v1")
	api.Post("/comment", createComment)
	app.Listen(":3000")
}
