package db

import (
	"time"

	"github.com/jinzhu/gorm"
)

const SelectLimit = 10

//Images テーブル定義
type Images struct {
	ID          int
	StoreID     int
	OriginID    int
	Kind        int
	Description string
	Order       int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (i *Images) Create(db *gorm.DB) {
	db.Create(i)
}
