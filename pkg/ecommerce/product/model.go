package product

import (
	"e-commerce-golang/pkg/ecommerce/category"
	"time"
)

type Product struct {
	Id        int                 `json:"id"`
	Name      string              `json:"name"`
	Price     float32             `json:"price"`
	Category  []category.Category `json:"category"`
	Sale      bool                `json:"sale"`
	PrevPrice float32             `json:"prev_price"`
	CreatedAt time.Time           `json:"createdAt"`
	UpdatedAt time.Time           `json:"updated_at"`
	DeletedAt time.Time           `json:"deleted_at"`
}
