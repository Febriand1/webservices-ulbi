package url

import (
	"github.com/Febriand1/webservices-ulbi/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func Webs(page *fiber.App) {
	page.Post("/api/whatsauth/request", controller.PostWhatsAuthRequest)  //API from user whatsapp message from iteung gowa
	page.Get("/ws/whatsauth/qr", websocket.New(controller.WsWhatsAuthQR)) //websocket whatsauth
	page.Get("/", controller.Home)                                        //ujicoba panggil package musik

	//percobaan
	page.Get("/members", controller.GetMembers)
	page.Get("/customers", controller.GetCustomers)
	page.Get("/incomes", controller.GetIncomes)

	page.Get("/members/:id", controller.GetMembersByID)
	page.Get("/customers/:id", controller.GetCustomersByID)
	page.Get("/incomes/:id", controller.GetIncomesByID)

	page.Post("/insmembers", controller.InsertMembers)
	page.Post("/inscustomers", controller.InsertCustomers)
	page.Post("/insincomes", controller.InsertIncomes)

	page.Put("/updmembers/:id", controller.UpdateMembers)
	page.Put("/updcustomers/:id", controller.UpdateCustomers)
	page.Put("/updincomes/:id", controller.UpdateIncomes)

	page.Delete("/delmembers/:id", controller.DeleteMembersByID)
	page.Delete("/delcustomers/:id", controller.DeleteCustomersByID)
	page.Delete("/delincomes/:id", controller.DeleteIncomesByID)

	page.Post("/insadmin", controller.InsertLoginAdmin)
	page.Post("/loginadmin", controller.LoginByAdmin)
	page.Post("/auth", controller.Authenticate)
}
