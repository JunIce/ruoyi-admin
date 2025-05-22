package service

import (
	"wms-server/v1/model"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

type RoleService struct {
	DB *gorm.DB
}

type RoleListReq struct {
	model.SysRole
	PageNum  int
	PageSize int
	Params   struct {
		BeginTime string
		EndTime   string
	}
}

func (s RoleService) GetRoleList(req *RoleListReq) ([]model.SysRoleVo, int64, error) {
	var roleList []model.SysRoleVo
	var total int64

	log.Info("req:", req)

	queryBuilder := s.DB.Model(&model.SysRole{})
	if req.RoleName != "" {
		queryBuilder = queryBuilder.Where("role_name LIKE ?", "%"+req.RoleName+"%")
	}
	if req.RoleKey != "" {
		queryBuilder = queryBuilder.Where("roke_key LIKE ?", "%"+req.RoleKey+"%")
	}
	if req.Status != "" {
		queryBuilder = queryBuilder.Where("status LIKE ?", "%"+req.Status+"%")
	}

	if req.Params.BeginTime != "" {
		queryBuilder = queryBuilder.Where("create_time >= ?", req.Params.BeginTime)
	}
	if req.Params.EndTime != "" {
		queryBuilder = queryBuilder.Where("create_time <= ?", req.Params.EndTime)
	}

	queryBuilder.Count(&total)
	result := queryBuilder.Limit(req.PageSize).Offset((req.PageNum - 1) * req.PageSize).Find(&roleList)

	if result.Error != nil {
		// 处理错误
		log.Error("Error occurred while creating user:", result.Error)
	}

	return roleList, total, nil
}

func (s RoleService) GetRoleByID(id int) (model.SysRoleVo, error) {
	var role model.SysRoleVo
	s.DB.Preload("Dept").Preload("Roles").Where("role_id = ?", id).First(&role)

	return role, nil
}

func (s RoleService) AddRole(role model.SysRoleAdd) (model.SysRoleAdd, error) {
	result := s.DB.Create(&role)

	if result.Error != nil {
		// 处理错误
		log.Error("Error occurred while creating user:", result.Error)
	} else {
		var roleMenus []model.SysRoleMenu
		for _, menuId := range role.MenuIds {
			userRole := model.SysRoleMenu{
				MenuID: menuId,
				RoleID: role.RoleID,
			}
			roleMenus = append(roleMenus, userRole)
		}
		s.DB.Model(&model.SysRoleMenu{}).Create(&roleMenus)

	}
	return role, result.Error
}

func (s RoleService) UpdateRole(role model.SysRoleAdd) (model.SysRoleAdd, error) {
	result := s.DB.Save(&role)

	if result.Error != nil {
		// 处理错误
		log.Error("Error occurred while creating user:", result.Error)
	}

	// 角色
	var roleMenus []model.SysRoleMenu

	// 删除旧的用户角色
	s.DB.Where("role_id = ?", role.RoleID).Delete(&model.SysRoleMenu{})

	for _, menuId := range role.MenuIds {
		roleMenu := model.SysRoleMenu{
			MenuID: menuId,
			RoleID: role.RoleID,
		}
		roleMenus = append(roleMenus, roleMenu)
	}
	s.DB.Model(&model.SysRoleMenu{}).Create(&roleMenus)

	return role, result.Error
}

func (s RoleService) DeleteRole(id int) error {
	// 删除用户
	result := s.DB.Where("role_id = ?", id).Delete(&model.SysRole{})
	if result.Error != nil {
		// 处理错误
		log.Error("Error occurred while creating user:", result.Error)
	}
	return result.Error
}

func (s RoleService) GetDeptTree() ([]model.SysDept, error) {
	var deptList []model.SysDept
	s.DB.Find(&deptList)
	return deptList, nil
}
