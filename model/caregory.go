package model

type Category struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Name      string `json:"name" gorm:"typevarchar(50); not null;uniqe"`
	CreatedAt Time   `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt Time   `json:"updated_at" gorm:"type:timestamp"`
}
