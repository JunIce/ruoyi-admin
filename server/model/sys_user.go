package model

import (
	"time"
)

// SysUser 用户信息表
type SysUser struct {
	UserID      int64      `gorm:"column:user_id;primaryKey;autoIncrement" json:"userId"`    // 用户ID
	DeptID      *int64     `gorm:"column:dept_id" json:"deptId,omitempty"`                   // 部门ID
	UserName    string     `gorm:"column:user_name;size:30;not null" json:"userName"`        // 用户账号
	NickName    string     `gorm:"column:nick_name;size:30;not null" json:"nickName"`        // 用户昵称
	UserType    string     `gorm:"column:user_type;size:2;default:'00'" json:"userType"`     // 用户类型（00系统用户）
	Email       string     `gorm:"column:email;size:50;default:''" json:"email"`             // 用户邮箱
	Phonenumber string     `gorm:"column:phonenumber;size:11;default:''" json:"phonenumber"` // 手机号码
	Sex         string     `gorm:"column:sex;size:1;default:'0'" json:"sex"`                 // 用户性别（0男 1女 2未知）
	Avatar      string     `gorm:"column:avatar;size:100;default:''" json:"avatar"`          // 头像地址
	Password    string     `gorm:"column:password;size:100;default:''" json:"password"`      // 密码
	Status      string     `gorm:"column:status;size:1;default:'0'" json:"status"`           // 帐号状态（0正常 1停用）
	DelFlag     string     `gorm:"column:del_flag;size:1;default:'0'" json:"delFlag"`        // 删除标志（0代表存在 2代表删除）
	LoginIP     string     `gorm:"column:login_ip;size:128;default:''" json:"loginIp"`       // 最后登录IP
	LoginDate   *time.Time `gorm:"column:login_date" json:"loginDate,omitempty"`             // 最后登录时间
	CreateBy    string     `gorm:"column:create_by;size:64;default:''" json:"createBy"`      // 创建者
	CreateTime  *time.Time `gorm:"column:create_time" json:"createTime,omitempty"`           // 创建时间
	UpdateBy    string     `gorm:"column:update_by;size:64;default:''" json:"updateBy"`      // 更新者
	UpdateTime  *time.Time `gorm:"column:update_time" json:"updateTime,omitempty"`           // 更新时间
	Remark      *string    `gorm:"column:remark;size:500" json:"remark,omitempty"`           // 备注
	CardID      *string    `gorm:"column:card_id;size:100" json:"cardId,omitempty"`          // 员工卡id(16进制的4个byte)
	PwdTime     *time.Time `gorm:"column:pwd_time" json:"pwdTime,omitempty"`                 // 密码更新时间
}

type SysUserVo struct {
	SysUser

	Dept  SysDept   `gorm:"foreignKey:DeptID;references:DeptID;" json:"dept"`
	Roles []SysRole `gorm:"many2many:sys_user_role;joinForeignKey:UserID;joinReferences:RoleID;" json:"roles"` // 用户与角色关联表
}

// 映射数据表
func (e *SysUser) TableName() string {
	return "sys_user"
}
