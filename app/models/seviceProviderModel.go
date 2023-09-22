package models

type ServiceProvider struct {
	ID   uint    `gorm:"column:id;primaryKey" json:"id" query:"id"`
	Name *string `gorm:"column:name" json:"name" reqHeader:"name"`
}

// protected $table =
func (ServiceProvider) TableName() string {
	return "mst_service_providers"
}
