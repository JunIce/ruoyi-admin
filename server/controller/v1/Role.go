package v1

import (
	"strconv"
	"wms-server/v1/model"
	"wms-server/v1/service"

	"github.com/gofiber/fiber/v2"
)

// 用户列表分页数据
func (bc BaseController) GetRoleList(c *fiber.Ctx) error {
	var roleService = service.RoleService{
		DB: bc.DB,
	}

	var req service.RoleListReq

	// 使用 QueryParser 自动绑定查询参数
	if err := c.QueryParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	result, total, err := roleService.GetRoleList(&req)

	if err != nil {
		return c.JSON(fiber.Map{
			"code":  500,
			"msg":   "查询失败",
			"data":  nil,
			"total": 0,
		})
	}

	return c.JSON(fiber.Map{
		"code":  0,
		"msg":   "查询成功",
		"rows":  result,
		"total": total,
	})
}

// 用户新增页面
func (bc BaseController) PostRole(c *fiber.Ctx) error {
	var roleService = service.RoleService{
		DB: bc.DB,
	}

	var role model.SysRoleAdd

	if err := c.BodyParser(&role); err != nil {
		return c.JSON(fiber.Map{
			"code": 400,
			"msg":  "参数错误",
			"data": err.Error(),
		})
	}

	u, err := roleService.AddRole(role)

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

func (bc BaseController) UpdateRole(c *fiber.Ctx) error {
	var roleService = service.RoleService{
		DB: bc.DB,
	}

	var role model.SysRoleAdd

	if err := c.BodyParser(&role); err != nil {
		return c.JSON(fiber.Map{
			"code": 400,
			"msg":  "参数错误",
			"data": err.Error(),
		})
	}

	u, err := roleService.UpdateRole(role)

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

// 用户删除
func (bc BaseController) DeleteRole(c *fiber.Ctx) error {
	var roleService = service.RoleService{
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

	err := roleService.DeleteRole(idInt)

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

// 用户查询
func (bc BaseController) GetRoleByID(c *fiber.Ctx) error {
	var roleService = service.RoleService{
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

	result, err := roleService.GetRoleByID(idInt)

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
