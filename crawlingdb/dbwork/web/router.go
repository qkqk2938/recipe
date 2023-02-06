package web

import (
	"github.com/gofiber/fiber/v2"
)


func Router() {
	app := fiber.New()

	app.Static("/", "./view/index.html") 

    app.Post("/api/selectbase", SelectBase)

    app.Listen(":3000")
}
