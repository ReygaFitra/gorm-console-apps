package entity

import "time"

type Products struct {
	ProductCode string `gorm:"column:product_code;primaryKey"`
	ProductName string `gorm:"column:product_name"`
	Stock       int    `gorm:"column:stock"`
	Price       int    `gorm:"column:price"`
}

type Orders struct {
	IdOrders  int `gorm:"column:id_order;primaryKey;autoIncrement"`
	OrderDate time.Time `gorm:"column:order_date"`
	PaymentMethod string `gorm:"column:payment_method"`
	OrderTotal int `gorm:"column:order_total"`
}

type OrderDetails struct {
	IdOrderDetail int `gorm:"column:id_order_detail;primaryKey;autoIncrement"`
	OrdersIdOrder int `gorm:"column:id_order"`
	Orders Orders `gorm:"foreignKey:OrdersIdOrder"`
	ProductsProductCode string `gorm:"column:product_code"`
	Products Products `gorm:"foreignKey:ProductsProductCode"`
	Qty int `gorm:"column:qty"`
}