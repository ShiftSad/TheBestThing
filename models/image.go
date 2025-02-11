package models

type Image struct {
	ID   uint   `gorm:"primaryKey,autoIncrement"`
	Name string `gorm:"uniqueIndex"`
	Data []byte `gorm:"type:bytea"`
}
