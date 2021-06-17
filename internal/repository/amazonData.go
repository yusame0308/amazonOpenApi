package repository

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

type AmazonData struct {
	Asin        string `gorm:"unique"`
	MakerName   string
	Price       int64
	ProductName string
	Reason      string
	Url         string
	IsDelete    soft_delete.DeletedAt `gorm:"softDelete:flag"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
