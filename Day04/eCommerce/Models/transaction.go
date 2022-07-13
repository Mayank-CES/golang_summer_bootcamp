package Models


type Transaction struct {

	TransactionId 	string	`json:"transaction_id" gorm:"primaryKey" binding:"required"`
	ProductId    	string 	`json:"product_id" binding:"required"`
	ConsumerId		string	`json:"consumer_id" binding:"required"`
	RetailerId		string	`json:"retailer_id" binding:"required"`
}

func (t *Transaction) TableName() string {
	return "transaction"
}
