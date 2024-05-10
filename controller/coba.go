package controller

import (
	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2"
	adawong "github.com/rayfanaqbil/BE_UTS_Webservice"
)

func Homepage(c *fiber.Ctx) error {
	ipaddr := musik.GetIPaddress()
	return c.JSON(ipaddr)
}

func GetGetAllGames(c *fiber.Ctx) error {
	ps := adawong.GetAllDataGames()
	return c.JSON(ps)
}