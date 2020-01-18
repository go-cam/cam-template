package models

import (
	"time"
)

type User struct {
	Id       int       `xorm:"not null pk autoincr INT(11)"`
	Name     string    `xorm:"not null comment('username') VARCHAR(64)"`
	DeleteAt time.Time `xorm:"comment('delete timestamp') TIMESTAMP"`
	UpdateAt time.Time `xorm:"default 'CURRENT_TIMESTAMP' comment('update timestamp') TIMESTAMP"`
	CreateAt time.Time `xorm:"default 'CURRENT_TIMESTAMP' comment('create timestamp') TIMESTAMP"`
}
