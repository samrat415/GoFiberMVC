package models

type ClientDocument struct {
	Id               *uint `gorm:"column:id"`
	ClientId         int   `gorm:"column:client_id"`
	VerificationFlag bool  `gorm:"column:verification_flag"`
}

// protected $table =
func (ClientDocument) TableName() string {
	return "tbl_client_documents"
}
