package web

import (
	db "dbwork/db"
	"github.com/gofiber/fiber/v2"
)



func SelectBase(c *fiber.Ctx) error{

	data := db.Perser()
	return c.JSON(data)

}