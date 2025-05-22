package v1

import "gorm.io/gorm"

type BaseController struct {
	DB *gorm.DB // 假设这是你的定义
}
