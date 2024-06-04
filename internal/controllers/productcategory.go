package controllers

import (
	"raidencommerce/internal/models"

	"github.com/sev-2/raiden"
)

type ProductCategoryController struct {
	raiden.ControllerBase
	Http  string `path:"/product-categories" type:"rest"`
	Model models.Productcategory
}
