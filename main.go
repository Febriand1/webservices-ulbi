package main

import (
	"log"

	"github.com/Febriand1/webservices-ulbi/config"

	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/whatsauth/whatsauth"

	"github.com/Febriand1/webservices-ulbi/url"

	"github.com/gofiber/fiber/v2"
)

func main() {
	go whatsauth.RunHub()
	site := fiber.New(config.Iteung)
	site.Use(cors.New(config.Cors))
	url.Web(site)
	log.Fatal(site.Listen(musik.Dangdut()))
}
