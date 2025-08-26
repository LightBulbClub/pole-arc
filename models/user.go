package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"username"`
	Realname string `gorm:"not null" json:"realname"`
	Password string `gorm:"not null" json:"-"` // 密码字段通常不应直接返回给客户端
	Email    string `gorm:"unique;not null" json:"email"`
	Grade    string `gorm:"not null" json:"grade"`
	Class    string `gorm:"not null" json:"class"` // 班级
}

type Teacher struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"username"`
	Realname string `gorm:"not null" json:"realname"`
	Password string `gorm:"not null" json:"-"` // 密码字段通常不应直接返回给客户端
	Email    string `gorm:"unique;not null" json:"email"`
	Subject  string `gorm:"not null" json:"subject"` // 教授科目
	Grade    string `gorm:"not null" json:"grade"`   // 年级
	Class    string `gorm:"not null" json:"class"`   // 班级
}
