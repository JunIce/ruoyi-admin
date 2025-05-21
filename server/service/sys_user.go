package service

import (
	"wms-server/v1/model"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

type UserListReq struct {
	PageNum  int
	PageSize int
	// DeptId   string
}

func (s UserService) GetUserList(req *UserListReq) ([]model.SysUserVo, int64, error) {
	var userList []model.SysUserVo
	var total int64

	log.Info("PageNum:", req.PageNum)
	log.Info("PageSize:", req.PageSize)

	s.DB.Preload("Dept").Preload("Roles").Where("1=1").Limit(req.PageSize).Offset((req.PageNum - 1) * req.PageSize).Omit("Password").Find(&userList)
	s.DB.Model(&model.SysUser{}).Where("1=1").Count(&total)

	// if err != nil {
	// 	log.Error(err)
	// 	return nil, 0, errors.New("查询失败")
	// }

	return userList, total, nil
}
