package models

import "time"

type BaseInfo struct {
	Modified  int
	CreatedBy string
	CreatedAt time.Time
	UpdatedBy string
	UpdatedAt time.Time
}
