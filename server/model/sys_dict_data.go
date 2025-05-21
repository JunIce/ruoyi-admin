package model

import (
	"time"
)

// SysDictData 字典数据表
type SysDictData struct {
	DictCode   int64      `gorm:"column:dict_code;primaryKey;autoIncrement" json:"dictCode"` // 字典编码
	DictSort   int        `gorm:"column:dict_sort;default:0" json:"dictSort"`                // 字典排序
	DictLabel  string     `gorm:"column:dict_label;size:100;default:''" json:"dictLabel"`    // 字典标签
	DictValue  string     `gorm:"column:dict_value;size:100;default:''" json:"dictValue"`    // 字典键值
	DictType   string     `gorm:"column:dict_type;size:100;default:''" json:"dictType"`      // 字典类型（关联 sys_dict_type）
	CssClass   *string    `gorm:"column:css_class;size:100" json:"cssClass,omitempty"`       // 样式属性（其他样式扩展）
	ListClass  *string    `gorm:"column:list_class;size:100" json:"listClass,omitempty"`     // 表格回显样式
	IsDefault  string     `gorm:"column:is_default;size:1;default:'N'" json:"isDefault"`     // 是否默认（Y是 N否）
	Status     string     `gorm:"column:status;size:1;default:'0'" json:"status"`            // 状态（0正常 1停用）
	CreateBy   string     `gorm:"column:create_by;size:64;default:''" json:"createBy"`       // 创建者
	CreateTime *time.Time `gorm:"column:create_time" json:"createTime,omitempty"`            // 创建时间
	UpdateBy   string     `gorm:"column:update_by;size:64;default:''" json:"updateBy"`       // 更新者
	UpdateTime *time.Time `gorm:"column:update_time" json:"updateTime,omitempty"`            // 更新时间
	Remark     *string    `gorm:"column:remark;size:500" json:"remark,omitempty"`            // 备注
}
