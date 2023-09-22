package middlewares

import (
	"eventsapp/app/initializers"
	"time"
)

type OauthAccessTokens struct {
	Id        string     `gorm:"column:id;primaryKey"`
	ClientId  *uint      `gorm:"column:client_id"`
	UserId    *uint      `gorm:"column:user_id"`
	Name      *string    `gorm:"column:name"`
	Revoked   *bool      `gorm:"column:revoked"`
	ExpiresAt *time.Time `gorm:"column:expires_at"`
}

func (oauth *OauthAccessTokens) Validate() bool {
	db := initializers.OauthDb
	db.Find(oauth)
	// Verify if Name is "Client"
	if oauth.Name != nil && *oauth.Name != "Client" {
		// Name is not "Client", invalid
		return false
	}

	// Verify if expires_at is later than now
	if oauth.ExpiresAt != nil {
		now := time.Now()
		if oauth.ExpiresAt.Before(now) {
			// expires_at is earlier than now, invalid
			return false
		}
	}
	// Verify if revoked is false
	if oauth.Revoked != nil && *oauth.Revoked {
		// revoked is true, invalid
		return false
	}
	// All conditions passed, oauth is valid
	return true
}

func (OauthAccessTokens) TableName() string {
	return "oauth_access_tokens"
}
