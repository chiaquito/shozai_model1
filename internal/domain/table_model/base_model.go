package table_model

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID            int       `gorm:"primarykey" json:"id"`
	Version       int       `gorm:"default:1" json:"version"`
	CreatedUserID int       `json:"createdUserId"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedUserID int       `json:"updatedUserId"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

// gorm hooks
func (m *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.Version = 1
	m.CreatedUserID = 1
	m.UpdatedUserID = 1
	return nil
}

// gorm hooks
// func (m *BaseModel) BeforeUpdate(tx *gorm.DB) (err error) {
// 	m.UpdatedUserID = 1
// 	m.Version++
// 	m.CreatedUserID = 1
// 	return nil
// }

// gorm hooks
func (m *BaseModel) BeforeSave(tx *gorm.DB) (err error) {
	m.UpdatedUserID = 1
	m.Version++
	return nil
}

func (m *BaseModel) BeforeUpdate(tx *gorm.DB) (err error) {
	tx.Statement.SetColumn("version", gorm.Expr("version + 6"))
	tx.Statement.SetColumn("updated_user_id", 1)
	tx.Statement.SetColumn("updated_at", time.Now())
	return nil
}
