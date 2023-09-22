package models

type Account struct {
	Id            *uint   `gorm:"column:id"`
	AccountNumber string  `gorm:"column:account_number"`
	AccountName   string  `gorm:"column:account_name"`
	Balance       float64 `gorm:"column:balance"`
	Status        bool    `gorm:"column:status"`
	LockFlag      bool    `gorm:"column:lock_flag"`
}

// protected $table =
func (Account) TableName() string {
	return "tbl_account"
}
