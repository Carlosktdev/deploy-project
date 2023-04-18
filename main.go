package main

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func main(){
	app := fiber.New()
	names := [5]string{"Juan", "Maria", "Pedro", "Ana", "Luis"} 

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!!")
    })

	app.Get("/api/getNames", func(c *fiber.Ctx) error {
         data := struct {
            Names []string `json:"names"`
        }{
            Names: names[:],
        }
        jsonData, err := json.Marshal(data)
        if err != nil {
            return err
        }
        return c.Status(200).Send(jsonData)
    })

	app.Listen("0.0.0.0:3000")
}