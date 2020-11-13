package models

import (
	"time"
)

//Order type
type Order struct {
	OrderID    uint      `json:"orderId" gorm:"column:order_id;primaryKey;autoIncrement"`
	OrderDate  time.Time `json:"orderDate" gorm:"column:order_date;autoCreateTime"`
	Status     uint      `json:"orderStatus" gorm:"column:order_status"`
	CustomerID uint      `gorm:"column:customer_id"`
	Customer   Customer  `json:"customer" gorm:"foreignKey:CustomerID;references:customer_id"`
	Items      []Item    `json:"items" gorm:"many2many:Orderdetails;foreignKey:OrderID;joinForeignKey:order_id;joinReferences:order_id"`
}
