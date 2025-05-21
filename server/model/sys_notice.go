package model

import (
	"time"
)

// SysNotice 通知公告表
type SysNotice struct {
	NoticeID      int        `gorm:"column:notice_id;primaryKey;autoIncrement" json:"noticeId"` // 公告ID
	NoticeTitle   string     `gorm:"column:notice_title;size:50;not null" json:"noticeTitle"`   // 公告标题
	NoticeType    string     `gorm:"column:notice_type;size:3;not null" json:"noticeType"`      // 公告类型（1通知 2公告）
	NoticeContent []byte     `gorm:"column:notice_content;type:longblob" json:"noticeContent"`  // 公告内容（存储为二进制）
	Status        string     `gorm:"column:status;size:1;default:'0'" json:"status"`            // 公告状态（0正常 1关闭）
	CreateBy      string     `gorm:"column:create_by;size:64;default:''" json:"createBy"`       // 创建者
	CreateTime    *time.Time `gorm:"column:create_time" json:"createTime,omitempty"`            // 创建时间
	UpdateBy      string     `gorm:"column:update_by;size:64;default:''" json:"updateBy"`       // 更新者
	UpdateTime    *time.Time `gorm:"column:update_time" json:"updateTime,omitempty"`            // 更新时间
	Remark        *string    `gorm:"column:remark;size:255" json:"remark,omitempty"`            // 备注
}
