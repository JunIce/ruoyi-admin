package model

import (
	"time"
)

// SysRole corresponds to the 'sys_role' table
type SysRole struct {
	// RoleID 角色ID
	RoleID int64 `gorm:"column:role_id;primaryKey;autoIncrement" json:"roleId"`
	// RoleName 角色名称
	RoleName string `gorm:"column:role_name;not null;size:30" json:"roleName"`
	// RoleKey 角色权限字符串
	RoleKey string `gorm:"column:role_key;not null;size:100" json:"roleKey"`
	// RoleSort 显示顺序
	RoleSort int `gorm:"column:role_sort;not null" json:"roleSort"`
	// DataScope 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
	DataScope string `gorm:"column:data_scope;size:1;default:1" json:"dataScope"`
	// MenuCheckStrictly 菜单树选择项是否关联显示
	MenuCheckStrictly bool `gorm:"column:menu_check_strictly;type:tinyint(1);default:1" json:"menuCheckStrictly"`
	// DeptCheckStrictly 部门树选择项是否关联显示
	DeptCheckStrictly bool `gorm:"column:dept_check_strictly;type:tinyint(1);default:1" json:"deptCheckStrictly"`
	// Status 角色状态（0正常 1停用）
	Status string `gorm:"column:status;not null;size:1" json:"status"`
	// DelFlag 删除标志（0代表存在 2代表删除）
	DelFlag string `gorm:"column:del_flag;size:1;default:0" json:"delFlag"`
	// CreateBy 创建者
	CreateBy string `gorm:"column:create_by;size:64;default:''" json:"createBy"`
	// CreateTime 创建时间
	CreateTime *time.Time `gorm:"column:create_time" json:"createTime"`
	// UpdateBy 更新者
	UpdateBy string `gorm:"column:update_by;size:64;default:''" json:"updateBy"`
	// UpdateTime 更新时间
	UpdateTime *time.Time `gorm:"column:update_time" json:"updateTime"`
	// Remark 备注
	Remark string `gorm:"column:remark;size:500" json:"remark"`
}

type SysRoleVo struct {
	SysRole
}

type SysRoleAdd struct {
	SysRole
	MenuIds []int64 `gorm:"-" json:"menuIds"` // 菜单ID
}

// TableName specifies the table name for GORM
func (*SysRole) TableName() string {
	return "sys_role"
}
