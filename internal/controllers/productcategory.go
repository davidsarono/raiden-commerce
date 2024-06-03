package controllers

import (
	"raidencommerce/internal/models"

	"github.com/sev-2/raiden"
)

type ProductCategoryRequest struct{}

type ProductCategoryResponse struct {
	Message string `json:"message"`
}

type ProductCategoryController struct {
	raiden.ControllerBase
	Http    string `path:"/product-categories" type:"rest"`
	Payload *ProductCategoryRequest
	Result  ProductCategoryResponse
	Model   models.Productcategory
}
