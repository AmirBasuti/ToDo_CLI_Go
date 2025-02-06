package Model

import "time"

type Todo struct {
	ID        uint
	Task      string `gorm:"size:64"`
	Complete  bool   `gorm:"default:False"`
	CreatedAt time.Time
}
