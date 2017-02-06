package db

import "time"

//ImageDescription テーブル定義
type ImageDescriptions struct {
	ID          int
	OriginID    int
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
