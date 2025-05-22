package model

import (
	"time"
)

// SysMenu 菜单权限表
type SysMenu struct {
	MenuID     int64      `gorm:"column:menu_id;primaryKey;autoIncrement" json:"menuId"` // 菜单ID
	MenuName   string     `gorm:"column:menu_name;size:50;not null" json:"menuName"`     // 菜单名称
	ParentID   int64      `gorm:"column:parent_id;default:0" json:"parentId"`            // 父菜单ID
	OrderNum   int        `gorm:"column:order_num;default:0" json:"orderNum"`            // 显示顺序
	Path       string     `gorm:"column:path;size:200;default:''" json:"path"`           // 路由地址
	Component  *string    `gorm:"column:component;size:255" json:"component,omitempty"`  // 组件路径
	Query      *string    `gorm:"column:query;size:255" json:"query,omitempty"`          // 路由参数
	IsFrame    int        `gorm:"column:is_frame;default:1" json:"isFrame"`              // 是否为外链（0是 1否）
	IsCache    int        `gorm:"column:is_cache;default:0" json:"isCache"`              // 是否缓存（0缓存 1不缓存）
	MenuType   string     `gorm:"column:menu_type;size:1;default:''" json:"menuType"`    // 菜单类型（M目录 C菜单 F按钮）
	Visible    string     `gorm:"column:visible;size:1;default:'0'" json:"visible"`      // 菜单状态（0显示 1隐藏）
	Status     string     `gorm:"column:status;size:1;default:'0'" json:"status"`        // 菜单状态（0正常 1停用）
	Perms      *string    `gorm:"column:perms;size:100" json:"perms,omitempty"`          // 权限标识
	Icon       string     `gorm:"column:icon;size:100;default:'#'" json:"icon"`          // 菜单图标
	CreateBy   string     `gorm:"column:create_by;size:64;default:''" json:"createBy"`   // 创建者
	CreateTime *time.Time `gorm:"column:create_time" json:"createTime,omitempty"`        // 创建时间
	UpdateBy   string     `gorm:"column:update_by;size:64;default:''" json:"updateBy"`   // 更新者
	UpdateTime *time.Time `gorm:"column:update_time" json:"updateTime,omitempty"`        // 更新时间
	Remark     string     `gorm:"column:remark;size:500;default:''" json:"remark"`       // 备注
	Category   string     `gorm:"column:category;size:100;default:'0'" json:"category"`  // 菜单分类，PC端-0, PAD端-1, PDA-2
}

// 映射数据表
func (e *SysMenu) TableName() string {
	return "sys_menu"
}
