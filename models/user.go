package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"username"`
	Password string `gorm:"not null" json:"-"` // 密码字段通常不应直接返回给客户端
	Email    string `gorm:"unique;not null" json:"email"`
}
