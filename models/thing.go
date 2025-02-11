package models

type Thing struct {
	ID          uint   `gorm:"primaryKey,autoIncrement"`
	Name        string `gorm:"uniqueIndex"`
	Description string
	Image       string
}
