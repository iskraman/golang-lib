package main

import (
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

var redisClient = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "changeme",
	DB:       0,
})

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello there 👋")
	})

	app.Listen(":3000")
}
