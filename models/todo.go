package models

type Todo struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"size:255"`
	Owner string `gorm:"size:255"`
}
