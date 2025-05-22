package service

import (
	"wms-server/v1/model"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

type DeptService struct {
	DB *gorm.DB
}

type DeptListReq struct {
	model.SysDept
	PageNum  int
	PageSize int
}

func (s DeptService) GetDeptList(req *DeptListReq) ([]model.SysDeptVo, int64, error) {
	var roleList []model.SysDeptVo
	var total int64

	log.Info("req:", req)

	queryBuilder := s.DB.Model(&model.SysDept{})
	if req.DeptName != "" {
		queryBuilder = queryBuilder.Where("dept_name LIKE ?", "%"+req.DeptName+"%")
	}
	if req.Status != "" {
		queryBuilder = queryBuilder.Where("status LIKE ?", "%"+req.Status+"%")
	}

	queryBuilder.Count(&total)
	result := queryBuilder.Limit(req.PageSize).Offset((req.PageNum - 1) * req.PageSize).Find(&roleList)

	if result.Error != nil {
		// 处理错误
		log.Error("Error occurred while creating user:", result.Error)
	}

	return roleList, total, nil
}

func (s DeptService) GetDeptByID(id int) (model.SysDeptVo, error) {
	var role model.SysDeptVo
	s.DB.Where("dept_id = ?", id).First(&role)

	return role, nil
}

func (s DeptService) AddDept(role model.SysDeptAdd) (model.SysDeptAdd, error) {
	result := s.DB.Create(&role)

	if result.Error != nil {
		// 处理错误
		log.Error("Error occurred while creating user:", result.Error)
	}
	return role, result.Error
}

func (s DeptService) UpdateDept(role model.SysDeptAdd) (model.SysDeptAdd, error) {
	result := s.DB.Save(&role)

	if result.Error != nil {
		// 处理错误
		log.Error("Error occurred while creating user:", result.Error)
	}

	return role, result.Error
}

func (s DeptService) DeleteDept(id int) error {
	// 删除用户
	result := s.DB.Where("role_id = ?", id).Delete(&model.SysDept{})
	if result.Error != nil {
		// 处理错误
		log.Error("Error occurred while creating user:", result.Error)
	}
	return result.Error
}
