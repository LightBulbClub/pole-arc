package models

import (
	"gorm.io/gorm"
)

// AssociationLog 社团活动日志
type AssociationLog struct {
	gorm.Model
	AssociationName string `gorm:"not null" json:"association_name"` // 社团名称
	ActivityTitle   string `gorm:"not null" json:"activity_title"`   // 活动标题
	ActivityType    string `gorm:"not null" json:"activity_type"`    // 活动类型
	Description     string `gorm:"type:text" json:"description"`     // 活动描述
	Location        string `json:"location"`                         // 活动地点
	Participants    int    `json:"participants"`                     // 参与人数
	Duration        int    `json:"duration"`                         // 活动时长（分钟）
	Organizer       string `gorm:"not null" json:"organizer"`        // 组织者
	Status          string `gorm:"default:'pending'" json:"status"`  // 状态：pending, approved, rejected
	Remarks         string `gorm:"type:text" json:"remarks"`         // 备注
}
