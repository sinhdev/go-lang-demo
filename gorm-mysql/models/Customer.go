package models

//Customer type
type Customer struct {
	CustomerID      uint   `json:"customerId" gorm:"column:customer_id;primaryKey;autoIncrement"`
	CustomerName    string `json:"customerName" gorm:"column:customer_name;not null"`
	CustomerAddress string `json:"customerAddress" gorm:"column:customer_address"`
}
