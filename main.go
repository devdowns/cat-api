package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type CatFact struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
	Time   string `json:"time"`
}

func main() {
	app := fiber.New()

	app.Get("/catfact", func(c *fiber.Ctx) error {
		timer := time.Now()
		// Make an HTTP GET request
		response, err := http.Get("https://catfact.ninja/fact")
		if err != nil {
			log.Fatalf("Failed to make the request: %v", err)
		}
		defer response.Body.Close()

		var catFact CatFact
		json.NewDecoder(response.Body).Decode(&catFact)

		catFact.Time = time.Since(timer).String()
		// Return the cat fact as JSON
		return c.JSON(catFact)
	})

	app.Listen(":3000")
}
