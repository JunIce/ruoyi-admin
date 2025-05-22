package model

import (
	"time"
)

// SysJob 定时任务调度表
type SysJob struct {
	JobID          int64      `gorm:"column:job_id;primaryKey;autoIncrement" json:"jobId"`                            // 任务ID
	JobName        string     `gorm:"column:job_name;size:64;not null;default:'';primaryKey" json:"jobName"`          // 任务名称
	JobGroup       string     `gorm:"column:job_group;size:64;not null;default:'DEFAULT';primaryKey" json:"jobGroup"` // 任务组名
	InvokeTarget   string     `gorm:"column:invoke_target;size:500;not null" json:"invokeTarget"`                     // 调用目标字符串
	CronExpression string     `gorm:"column:cron_expression;size:255;default:''" json:"cronExpression"`               // cron执行表达式
	MisfirePolicy  string     `gorm:"column:misfire_policy;size:20;default:'3'" json:"misfirePolicy"`                 // 计划执行错误策略（1立即执行 2执行一次 3放弃执行）
	Concurrent     string     `gorm:"column:concurrent;size:1;default:'1'" json:"concurrent"`                         // 是否并发执行（0允许 1禁止）
	Status         string     `gorm:"column:status;size:1;default:'0'" json:"status"`                                 // 状态（0正常 1暂停）
	CreateBy       string     `gorm:"column:create_by;size:64;default:''" json:"createBy"`                            // 创建者
	CreateTime     *time.Time `gorm:"column:create_time" json:"createTime,omitempty"`                                 // 创建时间
	UpdateBy       string     `gorm:"column:update_by;size:64;default:''" json:"updateBy"`                            // 更新者
	UpdateTime     *time.Time `gorm:"column:update_time" json:"updateTime,omitempty"`                                 // 更新时间
	Remark         *string    `gorm:"column:remark;size:500" json:"remark,omitempty"`                                 // 备注信息
}
