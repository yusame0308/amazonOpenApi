package repository

import (
	"time"
)

type AmazonData struct {
	Asin        string
	MakerName   string
	Price       int64
	ProductName string
	Reason      string
	Url         string
	Status      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
