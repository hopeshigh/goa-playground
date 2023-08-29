package models

import "gorm.io/gorm"

type Todo struct {
	ID          int    `gorm:"primaryKey"`
	Title       string `gorm:"type:varchar(255); not null"`
	Description string `gorm:"type:text; not null"`
	Status      string `gorm:"type:varchar(255); not null; default:'pending'"`
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&Todo{})
}
