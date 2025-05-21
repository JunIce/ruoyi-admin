package model

import (
	"time"
)

// SysDept 部门信息表
type SysDept struct {
	DeptID     int64      `gorm:"column:dept_id;primaryKey;autoIncrement" json:"deptId"` // 部门id
	ParentID   int64      `gorm:"column:parent_id;default:0" json:"parentId"`            // 父部门id
	Ancestors  string     `gorm:"column:ancestors;size:50;default:''" json:"ancestors"`  // 祖级列表
	DeptName   string     `gorm:"column:dept_name;size:30;default:''" json:"deptName"`   // 部门名称
	OrderNum   int        `gorm:"column:order_num;default:0" json:"orderNum"`            // 显示顺序
	Leader     *string    `gorm:"column:leader;size:20" json:"leader,omitempty"`         // 负责人
	Phone      *string    `gorm:"column:phone;size:11" json:"phone,omitempty"`           // 联系电话
	Email      *string    `gorm:"column:email;size:50" json:"email,omitempty"`           // 邮箱
	Status     string     `gorm:"column:status;size:1;default:'0'" json:"status"`        // 部门状态（0正常 1停用）
	DelFlag    string     `gorm:"column:del_flag;size:1;default:'0'" json:"delFlag"`     // 删除标志（0代表存在 2代表删除）
	CreateBy   string     `gorm:"column:create_by;size:64;default:''" json:"createBy"`   // 创建者
	CreateTime *time.Time `gorm:"column:create_time" json:"createTime,omitempty"`        // 创建时间
	UpdateBy   string     `gorm:"column:update_by;size:64;default:''" json:"updateBy"`   // 更新者
	UpdateTime *time.Time `gorm:"column:update_time" json:"updateTime,omitempty"`        // 更新时间
}

// 映射数据表
func (e *SysDept) TableName() string {
	return "sys_dept"
}
