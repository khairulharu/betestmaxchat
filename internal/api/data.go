package api

import (
	"beexamxchat/domain"

	"github.com/gofiber/fiber/v2"
)

type dataApi struct {
	dataService domain.DataService
}

func NewData(app *fiber.App, dataService domain.DataService) {
	handler := dataApi{
		dataService: dataService,
	}

	app.Get("/data", handler.getAllData)
	app.Get("/data/:code", handler.getDataByCode)
	app.Post("/data", handler.createData)
	app.Put("/data/:code", handler.updateData)
	app.Delete("/data/:code", handler.deleteData)
	app.Get("/filter", handler.FilterDataHandler)
}

func (d dataApi) getAllData(c *fiber.Ctx) error {

	response := d.dataService.GetAllData()

	return c.Status(fiber.StatusOK).JSON(response)
}

func (d dataApi) getDataByCode(c *fiber.Ctx) error {

	code := c.Params("code")

	response := d.dataService.GetDataByCode(code)

	return c.Status(response.Code).JSON(response)
}

func (d dataApi) createData(c *fiber.Ctx) error {

	var newData domain.Data

	if err := c.BodyParser(&newData); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	resonse := d.dataService.CreateData(newData)

	return c.Status(resonse.Code).JSON(resonse)
}

func (d dataApi) updateData(c *fiber.Ctx) error {
	code := c.Params("code")

	var updatedData domain.Data

	if err := c.BodyParser(&updatedData); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	response := d.dataService.UpdateData(code, updatedData)

	return c.Status(response.Code).JSON(updatedData)
}

func (d dataApi) deleteData(c *fiber.Ctx) error {
	code := c.Params("code")

	response := d.dataService.DeleteData(code)

	return c.Status(response.Code).JSON(response)
}

func (d dataApi) FilterDataHandler(c *fiber.Ctx) error {
	model := c.Query("model", "")
	tech := c.Query("tech", "")

	response := d.dataService.FilterData(model, tech)

	return c.Status(response.Code).JSON(response)
}
