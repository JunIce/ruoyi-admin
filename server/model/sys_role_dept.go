package model

// SysRoleDept 角色和部门关联表
type SysRoleDept struct {
	RoleID int64 `gorm:"column:role_id;primaryKey" json:"roleId"` // 角色ID
	DeptID int64 `gorm:"column:dept_id;primaryKey" json:"deptId"` // 部门ID
}
