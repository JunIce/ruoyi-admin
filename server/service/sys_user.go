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
	model.SysUser
	PageNum  int
	PageSize int
	// DeptId   string
}

func (s UserService) GetUserList(req *UserListReq) ([]model.SysUserVo, int64, error) {
	var userList []model.SysUserVo
	var total int64

	log.Info("req:", req)

	queryBuilder := s.DB.Model(&model.SysUser{}).Preload("Dept").Preload("Roles")
	if req.UserName != "" {
		queryBuilder = queryBuilder.Where("user_name LIKE ?", "%"+req.UserName+"%")
	}
	if req.NickName != "" {
		queryBuilder = queryBuilder.Where("nick_name LIKE ?", "%"+req.NickName+"%")
	}
	if req.Email != "" {
		queryBuilder = queryBuilder.Where("email LIKE ?", "%"+req.Email+"%")
	}
	if req.Phonenumber != "" {
		queryBuilder = queryBuilder.Where("phonenumber LIKE ?", "%"+req.Phonenumber+"%")
	}
	if req.DeptID != nil {
		queryBuilder = queryBuilder.Where("dept_id = ?", req.DeptID)
	}

	queryBuilder.Count(&total)
	result := queryBuilder.Limit(req.PageSize).Offset((req.PageNum - 1) * req.PageSize).Omit("Password").Find(&userList)

	if result.Error != nil {
		// 处理错误
		log.Error("Error occurred while creating user:", result.Error)
	}

	return userList, total, nil
}

func (s UserService) GetUserByID(id int) (model.SysUserVo, error) {
	var user model.SysUserVo
	// var user model.SysUser
	s.DB.Preload("Dept").Preload("Roles").Where("user_id = ?", id).Omit("Password").First(&user)

	return user, nil
}

func (s UserService) AddUser(user model.SysUserAdd) (model.SysUserAdd, error) {
	result := s.DB.Create(&user)

	if result.Error != nil {
		// 处理错误
		log.Error("Error occurred while creating user:", result.Error)
	} else {
		var userRoles []model.SysUserRole
		for _, role := range user.Roles {
			userRole := model.SysUserRole{
				UserID: user.UserID,
				RoleID: role,
			}
			userRoles = append(userRoles, userRole)
		}
		s.DB.Model(&model.SysUserRole{}).Create(&userRoles)

		var userPosts []model.SysUserPost
		for _, post := range user.PostIds {
			userPost := model.SysUserPost{
				UserID: user.UserID,
				PostID: post,
			}
			userPosts = append(userPosts, userPost)
		}
		s.DB.Model(&model.SysUserPost{}).Create(&userPosts)
	}
	return user, result.Error
}

func (s UserService) UpdateUser(user model.SysUserAdd) (model.SysUserAdd, error) {
	result := s.DB.Save(&user)

	if result.Error != nil {
		// 处理错误
		log.Error("Error occurred while creating user:", result.Error)
	}

	// 角色
	var userRoles []model.SysUserRole

	// 删除旧的用户角色
	s.DB.Where("user_id = ?", user.UserID).Delete(&model.SysUserRole{})

	for _, role := range user.Roles {
		userRole := model.SysUserRole{
			UserID: user.UserID,
			RoleID: role,
		}
		userRoles = append(userRoles, userRole)
	}
	s.DB.Model(&model.SysUserRole{}).Create(&userRoles)

	// 岗位
	var userPosts []model.SysUserPost
	// 删除旧的用户岗位
	s.DB.Where("user_id = ?", user.UserID).Delete(&model.SysUserPost{})

	for _, post := range user.PostIds {
		userPost := model.SysUserPost{
			UserID: user.UserID,
			PostID: post,
		}
		userPosts = append(userPosts, userPost)
	}
	s.DB.Model(&model.SysUserPost{}).Create(&userPosts)
	return user, result.Error
}

func (s UserService) DeleteUser(id int) error {
	// 删除用户
	result := s.DB.Where("user_id = ?", id).Delete(&model.SysUser{})
	if result.Error != nil {
		// 处理错误
		log.Error("Error occurred while creating user:", result.Error)
	}
	return result.Error
}

func (s UserService) GetDeptTree() ([]model.SysDept, error) {
	var deptList []model.SysDept
	s.DB.Find(&deptList)
	return deptList, nil
}
