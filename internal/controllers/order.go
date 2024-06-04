package controllers

import (
	"raidencommerce/internal/models"

	"github.com/sev-2/raiden"
)

// OrderController is the controller for the Order model
type OrderController struct {
	raiden.ControllerBase
	Http  string `path:"/orders" type:"rest"`
	Model models.Order
}
