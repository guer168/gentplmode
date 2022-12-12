package model

import (
	"time"
)

type JyUser struct {
	UID        int64     `gorm:"uid" json:"UID"`
	Account    string    `gorm:"account" json:"account"`
	Password   string    `gorm:"password" json:"password"`
	CreateTime time.Time `gorm:"create_time" json:"createTime"`
	UpdateTime time.Time `gorm:"update_time" json:"updateTime"`
}

// TableName
//  @Description: Getting the table name
//  @return string
func (model *JyUser) TableName() string {
	return "jy_user"
}
