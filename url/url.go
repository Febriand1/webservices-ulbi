package url

import (
	"github.com/Febriand1/webservices-ulbi/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func Web(page *fiber.App) {
	page.Post("/api/whatsauth/request", controller.PostWhatsAuthRequest)  //API from user whatsapp message from iteung gowa
	page.Get("/ws/whatsauth/qr", websocket.New(controller.WsWhatsAuthQR)) //websocket whatsauth
	page.Get("/", controller.Homepage)                                    //ujicoba panggil package musik
	page.Get("/presensi", controller.GetPresensi)
	page.Get("/test", controller.GetAll)
	page.Post("/nilai", controller.InsertData)
	page.Post("/mahasiswa", controller.InsertIdentitas)
	page.Post("/matakuliah", controller.InsertMatkul)
}
