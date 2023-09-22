package models

type ClientAccount struct {
	ClientId  *uint   `gorm:"column:client_id"`
	AccountId *uint   `gorm:"column:account_id"`
	Account   Account `gorm:"foreignKey:Id;references:AccountId"`
}

// protected $table =
func (ClientAccount) TableName() string {
	return "tbl_client_accounts"
}
