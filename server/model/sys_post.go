package model

import (
	"time"
)

// SysPost 岗位信息表
type SysPost struct {
	PostID     int64      `gorm:"column:post_id;primaryKey;autoIncrement" json:"postId"` // 岗位ID
	PostCode   string     `gorm:"column:post_code;size:64;not null" json:"postCode"`     // 岗位编码
	PostName   string     `gorm:"column:post_name;size:50;not null" json:"postName"`     // 岗位名称
	PostSort   int        `gorm:"column:post_sort;not null" json:"postSort"`             // 显示顺序
	Status     string     `gorm:"column:status;size:1;not null" json:"status"`           // 状态（0正常 1停用）
	CreateBy   string     `gorm:"column:create_by;size:64;default:''" json:"createBy"`   // 创建者
	CreateTime *time.Time `gorm:"column:create_time" json:"createTime,omitempty"`        // 创建时间
	UpdateBy   string     `gorm:"column:update_by;size:64;default:''" json:"updateBy"`   // 更新者
	UpdateTime *time.Time `gorm:"column:update_time" json:"updateTime,omitempty"`        // 更新时间
	Remark     *string    `gorm:"column:remark;size:500" json:"remark,omitempty"`        // 备注
}
