package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"os"
)

func main() {
	log.Println("Welcome to golang server template")

	app := fiber.New()

	app.Group("/api")
	app.Use(cors.New())
	app.Use(logger.New(logger.Config{
		TimeZone: "Asia/Seoul",
		Output:   os.Stderr,
	}))

	if err := app.Listen(fmt.Sprintf(":%v", os.Getenv("PORT"))); err != nil {
		log.Panicf("failed to listen on port %v..., check your config files", os.Getenv("PORT"))
	}
}
