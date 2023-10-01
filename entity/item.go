package entity

import "time"

type Items struct {
	Item_id     uint      `gorm:"primaryKey" json:"item_id"`
	Item_Code   string    `gorm:"uniqueKey" json:"item_code"`
	Description string    `gorm:"not null;type:varchar(191)" json:"description"`
	Quantity    int       `gorm:"not null" json:"quantity"`
	Order_id    uint      `gorm:"foreignKey:Order_id" json:"order_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
