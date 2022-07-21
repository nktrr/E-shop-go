package entity

import "time"

type ProductInfo struct {
	ProductID          string    `json:"product_id" gorm:"column:product_id"`
	CategoryType       int       `json:"category_type" gorm:"column:category_type"`
	CreateTime         time.Time `json:"create_time" gorm:"column:create_time"`
	ProductDescription string    `json:"product_description" gorm:"column:product_description"`
	ProductIcon        string    `json:"product_icon" gorm:"column:product_icon"`
	ProductName        string    `json:"product_name" gorm:"column:product_name"`
	ProductPrice       float64   `json:"product_price" gorm:"column:product_price"`
	ProductStatus      int       `json:"product_status" gorm:"column:product_status"`
	ProductStock       int       `json:"product_stock" gorm:"column:product_stock"`
	UpdateTime         time.Time `json:"update_time" gorm:"column:update_time"`
}

func (m *ProductInfo) TableName() string {
	return "product_info"
}
