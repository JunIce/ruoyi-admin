package service

import (
	"strings"
	"wms-server/v1/model"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

type MenuService struct {
	DB *gorm.DB
}

type MenuListReq struct {
	model.SysMenu
	PageNum  int
	PageSize int
}

func (s MenuService) GetMenuList(req *MenuListReq) ([]model.SysMenu, error) {
	var menuList []model.SysMenu

	log.Info("req:", req)

	queryBuilder := s.DB.Model(&model.SysMenu{})
	if req.MenuName != "" {
		queryBuilder = queryBuilder.Where("menu_name LIKE ?", "%"+req.MenuName+"%")
	}
	if req.Category != "" {
		queryBuilder = queryBuilder.Where("category LIKE ?", "%"+req.Category+"%")
	}
	if req.Status != "" {
		queryBuilder = queryBuilder.Where("email = ?", req.Status)
	}

	// queryBuilder.Count(&total)
	result := queryBuilder.Find(&menuList)

	if result.Error != nil {
		// 处理错误
		log.Error("Error occurred while creating menu:", result.Error)
	}

	return menuList, nil
}

func (s MenuService) GetMenuByID(id int) (model.SysMenu, error) {
	var menu model.SysMenu
	// var menu model.SysMenu
	s.DB.Preload("Dept").Preload("Roles").Where("menu_id = ?", id).Omit("Password").First(&menu)

	return menu, nil
}

func (s MenuService) AddMenu(menu model.SysMenu) (model.SysMenu, error) {
	result := s.DB.Create(&menu)
	return menu, result.Error
}

func (s MenuService) UpdateMenu(menu model.SysMenu) (model.SysMenu, error) {
	result := s.DB.Save(&menu)

	return menu, result.Error
}

func (s MenuService) DeleteMenu(id int) error {
	// 删除菜单
	result := s.DB.Where("menu_id = ?", id).Delete(&model.SysMenu{})
	if result.Error != nil {
		// 处理错误
		log.Error("Error occurred while creating menu:", result.Error)
	}
	return result.Error
}

func (s MenuService) GetMenuTree() ([]model.SysMenuRouteVo, error) {
	var menuList []model.SysMenu

	if err := s.DB.Where("menu_type in ('M', 'C')").Where("status = ?", "0").Where("category = ?", "0").Order("order_num asc").Find(&menuList).Error; err != nil {
		return nil, err
	}

	return s.BuildTree(menuList, 0), nil
}

// 获取路由地址
func getRouterPath(menu model.SysMenu) string {
	routerPath := menu.Path
	if menu.MenuType == "M" {
		if menu.Path != "/" {
			routerPath = "/" + menu.Path
		}
	} else if menu.MenuType == "C" && menu.ParentID == 0 {
		routerPath = "/"
	}
	return routerPath
}

// 构建树结构
func (e MenuService) BuildTree(list []model.SysMenu, parentId int64) []model.SysMenuRouteVo {
	var tree []model.SysMenuRouteVo

	for _, menu := range list {
		if menu.ParentID == parentId {
			route := model.SysMenuRouteVo{
				Hidden: menu.Visible == "1",
				Name:   strings.Title(menu.Path),
				Path:   getRouterPath(menu),
			}

			meta := model.Meta{Title: menu.MenuName, Icon: menu.Icon, NoCache: menu.IsCache == 1}
			route.Meta = meta

			if menu.MenuType == "M" {
				route.Component = "Layout"
				route.Redirect = "noRedirect"
				route.AlwaysShow = true
			} else if menu.MenuType == "C" {
				route.Component = menu.Component
			}

			children := e.BuildTree(list, menu.MenuID)
			if len(children) > 0 {
				route.Children = children
			}
			tree = append(tree, route)
		}
	}
	return tree
}
