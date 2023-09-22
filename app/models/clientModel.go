package models

import "time"

type Client struct {
	Id           *uint     `gorm:"column:id"`
	MobileNumber string    `gorm:"column:mobile_number"`
	DeletedBy    string    `gorm:"column:deleted_by"`
	DeletedAt    time.Time `gorm:"column:deleted_by"`
	//Documents     []ClientDocument `gorm:"foreignKey:ClientId;references:Id"`
	//ClientAccount ClientAccount    `gorm:"foreignKey:ClientId;references:Id"`
}

func (cl *Client) IsActive() bool {
	if cl.DeletedBy != "" {
		return false
	} else {
		return true
	}
}

//func (cl *Client) Balance() float64 {
//	return cl.ClientAccount.Account.Balance
//}
//func (cl *Client) IsVerified() bool {
//	for _, doc := range cl.Documents {
//		if doc.VerificationFlag {
//			return true
//		}
//	}
//	return false
//}

// protected $table =
func (Client) TableName() string {
	return "tbl_clients"
}
