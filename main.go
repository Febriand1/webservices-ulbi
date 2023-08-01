package main

import (
	"log"

	"github.com/Febriand1/webservices-ulbi/config"

	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/whatsauth/whatsauth"

	"github.com/Febriand1/webservices-ulbi/url"

	"github.com/gofiber/fiber/v2"

	// "github.com/gofiber/swagger"

	_ "github.com/Febriand1/webservices-ulbi/docs"
)

// @title TES  SWAG
// @version 1.0
// @description This is a sample server.

// @contact.name API Support
// @contact.url https://github.com/Febriand1
// @contact.email 1214039@std.ulbi.ac.id

// @host percobaan.herokuapp.com
// @BasePath /
// @schemes https http

func main() {
	go whatsauth.RunHub()
	site := fiber.New(config.Iteung)
	site.Use(cors.New(config.Cors))
	url.Webs(site)
	log.Fatal(site.Listen(musik.Dangdut()))
}
