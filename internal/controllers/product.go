package controllers

import (
	"raidencommerce/internal/models"

	"github.com/sev-2/raiden"
)

// create ProductRequest struct
type ProductRequest struct{}

// create ProductResponse struct
type ProductResponse struct {
	Message string `json:"message"`
}

// create ProductController struct
type ProductController struct {
	raiden.ControllerBase
	Http    string `path:"/products" type:"rest"`
	Payload *ProductRequest
	Result  ProductResponse
	Model   models.Product
}
