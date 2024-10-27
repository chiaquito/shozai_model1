package table_model

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID            int       `gorm:"primarykey" json:"id"`
	Version       int       `gorm:"default:1" json:"version"`
	CreatedUserID int       `json:"created_user_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedUserID int       `json:"updated_user_id"`
	UpdatedAt     time.Time `json:"updated_at"`
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
