package lead

import (
	"github.com/gofiber/fiber"
	"github.com/mehfooz-ul-haq/go-fiber-crm/database"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}

func GetLeads(c *fiber.Ctx) {
	db := database.DbConn

	var leads []Lead
	db.Find(&leads)
	c.Status(200).JSON(leads)
}

func GetLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DbConn

	var lead Lead
	db.Find(&lead, id)
	c.Status(200).JSON(lead)
}

func NewLead(c *fiber.Ctx) {
	db := database.DbConn
	lead := new(Lead)

	if err := c.BodyParser(lead); err != nil {
		c.Status(503).Send(err)
		return
	}

	db.Create(&lead)
	c.Status(200).JSON(lead)
}

func DeleteLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DbConn

	var lead Lead
	db.First(&lead, id)

	if lead.Name == "" {
		c.Status(500).Send("no lead found with ID")
	}

	db.Delete(&lead)
	c.Status(200).Send("Lead delete successfully.")
}
