package models

//Item type
type Item struct {
	ItemID          uint    `json:"itemId" gorm:"column:item_id;primaryKey;autoIncrement"`
	ItemName        string  `json:"itemName" gorm:"column:item_name"`
	UnitPrice       float64 `json:"unitPrice" gorm:"column:unit_price"`
	Amount          int32   `json:"amount" gorm:"column:amount"`
	ItemStatus      int16   `json:"status" gorm:"column:item_status"`
	ItemDescription string  `json:"description" gorm:"column:item_description"`
	// CreateAt        time.Time `json:"createAt" gorm:"autoCreateTime"`
	// UpdateAt        time.Time `json:"updateAt" gorm:"autoUpdateTime"`
}
