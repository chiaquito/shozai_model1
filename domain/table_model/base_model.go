package table_model

import "time"

type BaseModel struct {
	ID            int       `json:"id"`
	CreatedUserID int       `json:"created_user_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedUserID int       `json:"updated_user_id"`
	UpdatedAt     time.Time `json:"updated_at"`
}
