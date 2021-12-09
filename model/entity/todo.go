package entity

import "time"

type Todo struct {
	ID        uint `gorm:"primarykey"`
	Name      string
	Complete  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
