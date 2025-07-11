package v1

import (
	"strconv"
	"wms-server/v1/model"
	"wms-server/v1/service"

	"github.com/gofiber/fiber/v2"
)

// 菜单列表分页数据
func (bc BaseController) GetMenuList(c *fiber.Ctx) error {
	var menuService = service.MenuService{
		DB: bc.DB,
	}

	pageNum := c.Query("pageNum", "1")
	pageNumInt, _ := strconv.Atoi(pageNum)

	PageSize := c.Query("pageSize", "10")
	PageSizeInt, _ := strconv.Atoi(PageSize)

	var req = &service.MenuListReq{
		PageNum:  pageNumInt,
		PageSize: PageSizeInt,
		SysMenu: model.SysMenu{
			Status:   c.Query("status", ""),
			Category: c.Query("category", ""),
			MenuName: c.Query("menuName", ""),
		},
	}

	result, err := menuService.GetMenuList(req)

	if err != nil {
		return c.JSON(fiber.Map{
			"code":  500,
			"msg":   "查询失败",
			"data":  nil,
			"total": 0,
		})
	}

	return c.JSON(fiber.Map{
		"code": 0,
		"msg":  "查询成功",
		"data": result,
	})
}

// 菜单新增页面
func (bc BaseController) PostMenu(c *fiber.Ctx) error {
	var menuService = service.MenuService{
		DB: bc.DB,
	}

	var menu model.SysMenu

	if err := c.BodyParser(&menu); err != nil {
		return c.JSON(fiber.Map{
			"code": 400,
			"msg":  "参数错误",
			"data": err.Error(),
		})
	}

	u, err := menuService.AddMenu(menu)

	if err != nil {
		return c.JSON(fiber.Map{
			"code": 500,
			"msg":  "新增失败",
			"data": nil,
		})
	}

	return c.JSON(fiber.Map{
		"code": 200,
		"msg":  "新增成功",
		"data": u,
	})

}

func (bc BaseController) UpdateMenu(c *fiber.Ctx) error {
	var menuService = service.MenuService{
		DB: bc.DB,
	}

	var menu model.SysMenu

	if err := c.BodyParser(&menu); err != nil {
		return c.JSON(fiber.Map{
			"code": 400,
			"msg":  "参数错误",
			"data": err.Error(),
		})
	}

	u, err := menuService.UpdateMenu(menu)

	if err != nil {
		return c.JSON(fiber.Map{
			"code": 500,
			"msg":  "查询失败",
			"data": err,
		})
	}

	return c.JSON(fiber.Map{
		"code": 200,
		"msg":  "更新成功",
		"data": u,
	})
}

// 菜单删除
func (bc BaseController) DeleteMenu(c *fiber.Ctx) error {
	var menuService = service.MenuService{
		DB: bc.DB,
	}

	id := c.Params("id")
	if id == "" {
		return c.JSON(fiber.Map{
			"code": 400,
			"msg":  "参数错误",
			"data": nil,
		})
	}

	idInt, _ := strconv.Atoi(id)

	err := menuService.DeleteMenu(idInt)

	if err != nil {
		return c.JSON(fiber.Map{
			"code": 500,
			"msg":  "删除失败",
			"data": nil,
		})
	}

	return c.JSON(fiber.Map{
		"code": 200,
		"msg":  "删除成功",
	})
}

// 菜单查询
func (bc BaseController) GetMenuByID(c *fiber.Ctx) error {
	var menuService = service.MenuService{
		DB: bc.DB,
	}

	id := c.Params("id")
	if id == "" {
		return c.JSON(fiber.Map{
			"code": 400,
			"msg":  "参数错误",
			"data": nil,
		})
	}

	idInt, _ := strconv.Atoi(id)

	result, err := menuService.GetMenuByID(idInt)

	if err != nil {
		return c.JSON(fiber.Map{
			"code": 500,
			"msg":  "查询失败",
			"data": nil,
		})
	}

	return c.JSON(fiber.Map{
		"code": 200,
		"msg":  "查询成功",
		"data": result,
	})
}

// 部门列表分页数据
func (bc BaseController) GetRouters(c *fiber.Ctx) error {
	var menuService = service.MenuService{
		DB: bc.DB,
	}

	result, err := menuService.GetMenuTree()
	if err != nil {
		return c.JSON(fiber.Map{
			"code": -1,
			"msg":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"code": 0,
		"msg":  "查询成功",
		"data": result,
	})
}
