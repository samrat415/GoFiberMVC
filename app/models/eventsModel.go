package models

import (
	"encoding/json"
	"github.com/spf13/viper"
	"time"
)

// protected $cast,
type Events struct {
	ID            *uint           `gorm:"column:id;primaryKey" json:"id" query:"id"`
	Name          *string         `gorm:"column:name" json:"name" reqHeader:"name"`
	IsActive      *bool           `gorm:"column:is_active" json:"is_active" reqHeader:"is_active"`
	StartDate     *string         `gorm:"column:start_date" json:"start_date"`
	EndDate       *string         `gorm:"column:end_date" json:"end_date"`
	SupportEmail  *string         `gorm:"column:support_email" json:"support_email"`
	SupportNumber *string         `gorm:"column:support_number" json:"support_number"`
	MainImage     *string         `gorm:"column:main_image" json:"main_image"`
	Description   *string         `gorm:"column:description" json:"description"`
	EventType     *uint           `gorm:"column:event_type" json:"event_type"`
	Rank          *uint           `gorm:"column:rank" json:"rank"`
	ProviderId    *uint           `gorm:"column:provider" json:"provider_id"`
	Provider      ServiceProvider `gorm:"references:ID;foreignKey:ProviderId" json:"provider_details"`
	CreatedAt     *time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     *time.Time      `gorm:"column:updated_at" json:"updated_at"`
	MainImageURL  *string         `gorm:"-" json:"main_image_url"`
}

func (e *Events) GetMainImageURL() string {
	if e.MainImage != nil {
		return viper.GetString("resourceURL") + *e.MainImage
	}
	return viper.GetString("defaultImage")
}

// Override the MarshalJSON method to include MainImageURL attribute in JSON output
func (e *Events) MarshalJSON() ([]byte, error) {
	type Alias Events
	return json.Marshal(&struct {
		Alias
		MainImageURL string `json:"main_image_url"`
	}{
		Alias:        (Alias)(*e),
		MainImageURL: e.GetMainImageURL(),
	})
}

// Fillable Fields returns the list of fillable field names
func (e *Events) FillableFields() []string {
	return []string{
		"Name",
		"IsActive",
		"StartDate",
		"EndDate",
		"SupportEmail",
		"SupportNumber",
		"MainImage",
		"Description",
		"EventType",
		"Rank",
		"Provider",
	}
}

// protected $table =
func (Events) TableName() string {
	return "tbl_events"
}
