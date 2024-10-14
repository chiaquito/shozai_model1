package table_model

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID            int       `json:"id"`
	CreatedUserID int       `json:"created_user_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedUserID int       `json:"updated_user_id"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// gorm hooks
func (m *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now()
	m.CreatedAt = now
	m.CreatedUserID = 1
	m.UpdatedAt = now
	m.UpdatedUserID = 1
	return nil
}
