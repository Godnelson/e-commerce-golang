package operation

import (
	"e-commerce-golang/pkg/ecommerce/costumer"
	"e-commerce-golang/pkg/ecommerce/product"
	"time"
)

type Operation struct {
	Id            int               `json:"id"`
	Product       []product.Product `json:"product"`
	Costumer      costumer.Costumer `json:"costumer"`
	TotalPrice    float32           `json:"total_price"`
	PaymentMethod string            `json:"payment_method"`
	CreatedAt     time.Time         `json:"createdAt"`
	UpdatedAt     time.Time         `json:"updated_at"`
	DeletedAt     time.Time         `json:"deleted_at"`
}
