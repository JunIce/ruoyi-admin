package v1

import (
	"strconv"
	"wms-server/v1/model"
	"wms-server/v1/service"

	"github.com/gofiber/fiber/v2"
)

// 部门列表分页数据
func (bc BaseController) GetDeptList(c *fiber.Ctx) error {
	var deptService = service.DeptService{
		DB: bc.DB,
	}

	var req service.DeptListReq

	// 使用 QueryParser 自动绑定查询参数
	if err := c.QueryParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	result, total, err := deptService.GetDeptList(&req)

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

// 部门新增页面
func (bc BaseController) PostDept(c *fiber.Ctx) error {
	var deptService = service.DeptService{
		DB: bc.DB,
	}

	var role model.SysDeptAdd

	if err := c.BodyParser(&role); err != nil {
		return c.JSON(fiber.Map{
			"code": 400,
			"msg":  "参数错误",
			"data": err.Error(),
		})
	}

	u, err := deptService.AddDept(role)

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

func (bc BaseController) UpdateDept(c *fiber.Ctx) error {
	var deptService = service.DeptService{
		DB: bc.DB,
	}

	var role model.SysDeptAdd

	if err := c.BodyParser(&role); err != nil {
		return c.JSON(fiber.Map{
			"code": 400,
			"msg":  "参数错误",
			"data": err.Error(),
		})
	}

	u, err := deptService.UpdateDept(role)

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

// 部门删除
func (bc BaseController) DeleteDept(c *fiber.Ctx) error {
	var deptService = service.DeptService{
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

	err := deptService.DeleteDept(idInt)

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

// 部门查询
func (bc BaseController) GetDeptByID(c *fiber.Ctx) error {
	var deptService = service.DeptService{
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

	result, err := deptService.GetDeptByID(idInt)

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
