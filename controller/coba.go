package controller

import(
	"github.com/aiteung/musik"
	cek "github.com/aiteung/presensi"
	"github.com/gofiber/fiber/v2"
	"github.com/Febriand1/webservices-ulbi/config"
)

func Homepage(c *fiber.Ctx) error {
	ipaddr := musik.GetIPaddress()
	return c.JSON(ipaddr)
}

func GetPresensi(c *fiber.Ctx) error {
     ps := cek.GetPresensiCurrentMonth(config.Ulbimongoconn)
     return c.JSON(ps)
}