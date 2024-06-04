package controllers

import (
	"raidencommerce/internal/models"

	"github.com/sev-2/raiden"
)

type ProductController struct {
	raiden.ControllerBase
	Http  string `path:"/products" type:"rest"`
	Model models.Product
}
