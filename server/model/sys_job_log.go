package model

import (
	"time"
)

// SysJobLog 定时任务调度日志表
type SysJobLog struct {
	JobLogID      int64      `gorm:"column:job_log_id;primaryKey;autoIncrement" json:"jobLogId"`      // 任务日志ID
	JobName       string     `gorm:"column:job_name;size:64;not null" json:"jobName"`                 // 任务名称
	JobGroup      string     `gorm:"column:job_group;size:64;not null" json:"jobGroup"`               // 任务组名
	InvokeTarget  string     `gorm:"column:invoke_target;size:500;not null" json:"invokeTarget"`      // 调用目标字符串
	JobMessage    *string    `gorm:"column:job_message;size:500" json:"jobMessage,omitempty"`         // 日志信息
	Status        string     `gorm:"column:status;size:1;default:'0'" json:"status"`                  // 执行状态（0正常 1失败）
	ExceptionInfo string     `gorm:"column:exception_info;size:2000;default:''" json:"exceptionInfo"` // 异常信息
	CreateTime    *time.Time `gorm:"column:create_time" json:"createTime,omitempty"`                  // 创建时间
}
