package controllers

import (
	"raidencommerce/internal/models"

	"github.com/sev-2/raiden"
)

// OrderitemController is the controller for the Orderitem model
type OrderitemController struct {
	raiden.ControllerBase
	Http  string `path:"/orderitems" type:"rest"`
	Model models.Orderitem
}
