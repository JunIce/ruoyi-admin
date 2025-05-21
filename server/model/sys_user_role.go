package model

// SysUserRole 用户和角色关联表
type SysUserRole struct {
	UserID int64 `gorm:"column:user_id;primaryKey;foreignKey:UserID" json:"userId"` // 用户ID
	RoleID int64 `gorm:"column:role_id;primaryKey;foreignKey:RoleID" json:"roleId"` // 角色ID
}

// 映射数据表
func (e *SysUserRole) TableName() string {
	return "sys_user_role"
}
