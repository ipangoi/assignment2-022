package entity

import (
	"time"
)

type Orders struct {
	Order_id     uint      `gorm:"primaryKey" json:"order_id"`
	CustomerName string    `gorm:"not null;type:varchar(191)" json:"customer_name"`
	OrderedAt    time.Time `json:"ordered_at"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Items        []Items   `gorm:"foreignKey:Order_id" json:"items"`
}

type UpdateOrder struct {
	CustomerName string    `json:"customer_name,omitempty"`
	UpdatedAt    time.Time `json:"updated_at"`
	Items        []Items   `json:"items"`
}
