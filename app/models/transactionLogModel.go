package models

import (
	"gorm.io/gorm"
)

type TransactionLog struct {
	gorm.Model
	UniqueIdentifier              string          `gorm:"column:unique_identifier" json:"id"`
	ProviderId                    string          `gorm:"column:provider" json:"provider_id"`
	ProviderTransactionIdentifier string          `gorm:"column:provider_transaction_identifier" json:"provider_transaction_identifier"`
	TransactionType               string          `gorm:"column:transaction_type" json:"transaction_type"`
	TransactionStatus             string          `gorm:"column:transaction_status" json:"transaction_status"`
	Remarks                       string          `gorm:"column:remarks" json:"remarks"`
	ProviderRemarks               string          `gorm:"column:provider_remarks" json:"provider_remarks"`
	ClientID                      int             `gorm:"column:client_id" json:"client_id"`
	ServiceID                     string          `gorm:"column:service_id" json:"service_id"`
	Service                       Service         `gorm:"references:ID;foreignKey:ServiceID" json:"service"`
	ServiceProvider               ServiceProvider `gorm:"references:ID;foreignKey:ProviderId" json:"provider_details" json:"service_provider"`
	Amount                        float64         `gorm:"column:amount" json:"amount"`
	ProviderFees                  string          `gorm:"column:provider_fees" json:"provider_fees"`
	ServiceFees                   string          `gorm:"column:service_fees" json:"service_fees"`
	Cashback                      float64         `gorm:"column:cashback" json:"cashback"`
	RewardPoints                  string          `gorm:"column:reward_points" json:"reward_points"`
	DiscountCode                  string          `gorm:"column:discount_code" json:"discount_code"`
	DiscountAmount                string          `gorm:"column:discount_amount" json:"discount_amount"`
	TransactionPurpose            string          `gorm:"column:transaction_purpose" json:"transaction_purpose"`
	AmountStatus                  string          `gorm:"column:amount_status" json:"amount_status"`
	Balance                       string          `gorm:"column:balance" json:"balance"`
	TransactionEndResponse        string          `gorm:"column:transaction_end_response" json:"transaction_end_response"`
	TransactionDate               string          `gorm:"column:transaction_date" json:"transaction_date"`
	DeletedAt                     string          `gorm:"column:deleted_at"`
	DeletedBy                     string          `gorm:"column:deleted_by"`
	ProviderAccount               string          `gorm:"column:provider_account"`
	ClientAccount                 string          `gorm:"column:client_account"`
	EnID                          string          `gorm:"column:en_id"`
}

// protected $table =
func (TransactionLog) TableName() string {
	return "tbl_transactions_log"
}
