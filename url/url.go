package url

import (
	"github.com/Febriand1/webservices-ulbi/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func Web(page *fiber.App) {
	page.Post("/api/whatsauth/request", controller.PostWhatsAuthRequest)  //API from user whatsapp message from iteung gowa
	page.Get("/ws/whatsauth/qr", websocket.New(controller.WsWhatsAuthQR)) //websocket whatsauth
	page.Get("/", controller.Home)                                        //ujicoba panggil package musik
	page.Get("/presensiold", controller.GetPresensiold)
	page.Get("/presensi", controller.GetPresensi)
	page.Get("/matakuliah", controller.GetMatakuliah)
	page.Get("/grade", controller.GetGrade)
	page.Get("/test", controller.GetAll)
	page.Post("/insnilai", controller.InsertData)
	page.Post("/insmahasiswa", controller.InsertIdentitas)
	page.Post("/insmatakuliah", controller.InsertMatkul)
	page.Post("/insdosen", controller.InsertDosen)


	//sekarang
	page.Get("/presensinew", controller.GetAllPresensi) //menampilkan seluruh data presensi
	page.Get("/presensinew/:id", controller.GetPresensiID) //menampilkan data presensi berdasarkan id
}
