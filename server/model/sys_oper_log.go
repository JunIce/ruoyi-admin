package model

import (
	"time"
)

// SysOperLog 操作日志记录
type SysOperLog struct {
	OperID        int64      `gorm:"column:oper_id;primaryKey;autoIncrement" json:"operId"`         // 日志主键
	Title         string     `gorm:"column:title;size:50;default:''" json:"title"`                  // 模块标题
	BusinessType  int        `gorm:"column:business_type;default:0" json:"businessType"`            // 业务类型（0其它 1新增 2修改 3删除）
	Method        string     `gorm:"column:method;size:100;default:''" json:"method"`               // 方法名称
	RequestMethod string     `gorm:"column:request_method;size:10;default:''" json:"requestMethod"` // 请求方式
	OperatorType  int        `gorm:"column:operator_type;default:0" json:"operatorType"`            // 操作类别（0其它 1后台用户 2手机端用户）
	OperName      string     `gorm:"column:oper_name;size:50;default:''" json:"operName"`           // 操作人员
	DeptName      string     `gorm:"column:dept_name;size:50;default:''" json:"deptName"`           // 部门名称
	OperURL       string     `gorm:"column:oper_url;size:255;default:''" json:"operUrl"`            // 请求URL
	OperIP        string     `gorm:"column:oper_ip;size:128;default:''" json:"operIp"`              // 主机地址
	OperLocation  string     `gorm:"column:oper_location;size:255;default:''" json:"operLocation"`  // 操作地点
	OperParam     string     `gorm:"column:oper_param;size:2000;default:''" json:"operParam"`       // 请求参数
	JsonResult    string     `gorm:"column:json_result;size:2000;default:''" json:"jsonResult"`     // 返回参数
	Status        int        `gorm:"column:status;default:0" json:"status"`                         // 操作状态（0正常 1异常）
	ErrorMsg      string     `gorm:"column:error_msg;size:2000;default:''" json:"errorMsg"`         // 错误消息
	OperTime      *time.Time `gorm:"column:oper_time" json:"operTime,omitempty"`                    // 操作时间
	CostTime      int64      `gorm:"column:cost_time;default:0" json:"costTime"`                    // 消耗时间，单位毫秒
}
