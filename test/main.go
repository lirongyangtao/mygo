package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(":3000"))
}

//package main
//
//import "github.com/flamego/flamego"
//
//func main() {
//	f := flamego.Classic()
//	f.Get("/", func() string {
//		return "Hello, Flamego!"
//	})
//	f.Run()
//}
