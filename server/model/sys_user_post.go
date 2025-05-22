package model

// SysUserPost 用户与岗位关联表
type SysUserPost struct {
	UserID int64 `gorm:"column:user_id;primaryKey" json:"userId"` // 用户ID
	PostID int64 `gorm:"column:post_id;primaryKey" json:"postId"` // 岗位ID
}

func (e *SysUserPost) TableName() string {
	return "sys_user_post"
}
