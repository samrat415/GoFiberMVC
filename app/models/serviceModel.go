package models

type Service struct {
	ID   *uint  `gorm:"column:id;primaryKey" json:"id" query:"id"`
	Name string `gorm:"column:name" json:"name" query:"name"`
}

// Fillable Fields returns the list of fillable field names
func (e *Service) FillableFields() []string {
	return []string{
		"Name",
	}
}

// protected $table =
func (Service) TableName() string {
	return "tbl_services"
}
