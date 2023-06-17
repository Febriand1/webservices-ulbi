package url

import (
	"github.com/Febriand1/webservices-ulbi/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/gofiber/websocket/v2"
)

func Web(page *fiber.App) {
	page.Post("/api/whatsauth/request", controller.PostWhatsAuthRequest)  //API from user whatsapp message from iteung gowa
	page.Get("/ws/whatsauth/qr", websocket.New(controller.WsWhatsAuthQR)) //websocket whatsauth
	page.Get("/", controller.Home)                                        //ujicoba panggil package musik
	page.Get("/presensiold", controller.GetPresensiold)

	//nilai
	page.Get("/presensi", controller.GetPresensi)
	page.Get("/matakuliah", controller.GetMatakuliah)
	page.Get("/grade", controller.GetGrade)
	page.Get("/test", controller.GetAll)

	page.Post("/insmahasiswa", controller.InsertIdentitas)
	page.Post("/insmatakuliah", controller.InsertMatkul)
	page.Post("/insdosen", controller.InsertDosen)

	//TB
	page.Get("/nilai", controller.GetAllNilai)
	page.Get("/nilai/:id", controller.GetNilaiID)
	page.Post("/insnilai", controller.InsertDataNilai)
	page.Put("/updnilai/:id", controller.UpdateDataNilai)
	page.Delete("/delnilai/:id", controller.DeleteNilaiID)
	//TB

	//login
	page.Post("/loginadmin", controller.LoginAdmin)
	//login


	// //login
	// page.Get("/admin", controller.GetAdmin)
	// page.Get("/admin/:id", controller.GetAdminID)
	// page.Post("/insadmin", controller.InsertLoginAdmin)
	// page.Get("/user", controller.GetUser)
	// page.Get("/user/:id", controller.GetUserID)
	// page.Post("/insuser", controller.InsertLoginUser)
	// //login

	//sekarang presensi
	page.Get("/presensinew", controller.GetAllPresensi) //menampilkan seluruh data presensi
	page.Get("/presensinew/:id", controller.GetPresensiID) //menampilkan data presensi berdasarkan id
	//insert data presensi
	page.Post("/ins", controller.InsertData1)
	//edit data presensi
	page.Put("/upd/:id", controller.UpdateData)
	//delete data presensi
	page.Delete("/delete/:id", controller.DeletePresensiByID)

	//swagger
	page.Get("/docs/*", swagger.HandlerDefault)
}


