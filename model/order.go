package model

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	OrderID     uint64     `json:"order_id"`
	CustomerID  uuid.UUID  `json:"customer_id"`
	LineItems   []LineItem `json:"line_items"`
	CreatedAt   *time.Time `json:"created_at"`
	ShippeddAt  *time.Time `json:"shippedd_at"`
	CompletedAt *time.Time `json:"completed_at"`
}

type LineItem struct {
	ItemID  uuid.UUID `json:"item_id"`
	Quantiy uint      `json:"quantiy"`
	Price   uint      `json:"price"`
}
