package model

import (
	"time"
)

// SysConfig 参数配置表
type SysConfig struct {
	ConfigID    int        `gorm:"column:config_id;primaryKey;autoIncrement" json:"configId"`  // 参数主键
	ConfigName  string     `gorm:"column:config_name;size:100;default:''" json:"configName"`   // 参数名称
	ConfigKey   string     `gorm:"column:config_key;size:100;default:''" json:"configKey"`     // 参数键名
	ConfigValue string     `gorm:"column:config_value;size:500;default:''" json:"configValue"` // 参数键值
	ConfigType  string     `gorm:"column:config_type;size:1;default:'N'" json:"configType"`    // 系统内置（Y是 N否）
	CreateBy    string     `gorm:"column:create_by;size:64;default:''" json:"createBy"`        // 创建者
	CreateTime  *time.Time `gorm:"column:create_time" json:"createTime,omitempty"`             // 创建时间
	UpdateBy    string     `gorm:"column:update_by;size:64;default:''" json:"updateBy"`        // 更新者
	UpdateTime  *time.Time `gorm:"column:update_time" json:"updateTime,omitempty"`             // 更新时间
	Remark      *string    `gorm:"column:remark;size:500" json:"remark,omitempty"`             // 备注
}
