package models

import (
	"time"
)

type Migration struct {
	Version string    `xorm:"not null pk VARCHAR(255)"`
	DoneAt  time.Time `xorm:"not null TIMESTAMP"`
}
