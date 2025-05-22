package model

// SysRoleMenu 角色和菜单关联表
type SysRoleMenu struct {
	RoleID int64 `gorm:"column:role_id;primaryKey" json:"roleId"` // 角色ID
	MenuID int64 `gorm:"column:menu_id;primaryKey" json:"menuId"` // 菜单ID
}
