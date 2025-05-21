package model

import (
	"time"
)

// SysDictType 字典类型表
type SysDictType struct {
	DictID     int64      `gorm:"column:dict_id;primaryKey;autoIncrement" json:"dictId"`       // 字典主键
	DictName   string     `gorm:"column:dict_name;size:100;default:''" json:"dictName"`        // 字典名称
	DictType   string     `gorm:"column:dict_type;size:100;unique;default:''" json:"dictType"` // 字典类型（唯一）
	Status     string     `gorm:"column:status;size:1;default:'0'" json:"status"`              // 状态（0正常 1停用）
	CreateBy   string     `gorm:"column:create_by;size:64;default:''" json:"createBy"`         // 创建者
	CreateTime *time.Time `gorm:"column:create_time" json:"createTime,omitempty"`              // 创建时间
	UpdateBy   string     `gorm:"column:update_by;size:64;default:''" json:"updateBy"`         // 更新者
	UpdateTime *time.Time `gorm:"column:update_time" json:"updateTime,omitempty"`              // 更新时间
	Remark     *string    `gorm:"column:remark;size:500" json:"remark,omitempty"`              // 备注
}
