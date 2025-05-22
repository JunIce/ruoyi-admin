package v1

import (
	"strconv"
	"wms-server/v1/model"
	"wms-server/v1/service"

	"github.com/gofiber/fiber/v2"
)

// 用户列表分页数据
func (bc BaseController) GetUserList(c *fiber.Ctx) error {
	var userService = service.UserService{
		DB: bc.DB,
	}

	pageNum := c.Query("pageNum", "1")
	pageNumInt, _ := strconv.Atoi(pageNum)

	PageSize := c.Query("pageSize", "10")
	PageSizeInt, _ := strconv.Atoi(PageSize)

	var req = &service.UserListReq{
		PageNum:  pageNumInt,
		PageSize: PageSizeInt,
		SysUser: model.SysUser{
			Email:       c.Query("email", ""),
			UserName:    c.Query("userName", ""),
			NickName:    c.Query("nickName", ""),
			Phonenumber: c.Query("phonenumber", ""),
		},
	}

	result, total, err := userService.GetUserList(req)

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
func (bc BaseController) PostUser(c *fiber.Ctx) error {
	var userService = service.UserService{
		DB: bc.DB,
	}

	var user model.SysUserAdd

	if err := c.BodyParser(&user); err != nil {
		return c.JSON(fiber.Map{
			"code": 400,
			"msg":  "参数错误",
			"data": err.Error(),
		})
	}

	u, err := userService.AddUser(user)

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

func (bc BaseController) UpdateUser(c *fiber.Ctx) error {
	var userService = service.UserService{
		DB: bc.DB,
	}

	var user model.SysUserAdd

	if err := c.BodyParser(&user); err != nil {
		return c.JSON(fiber.Map{
			"code": 400,
			"msg":  "参数错误",
			"data": err.Error(),
		})
	}

	u, err := userService.UpdateUser(user)

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
func (bc BaseController) DeleteUser(c *fiber.Ctx) error {
	var userService = service.UserService{
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

	err := userService.DeleteUser(idInt)

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
func (bc BaseController) GetUserByID(c *fiber.Ctx) error {
	var userService = service.UserService{
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

	result, err := userService.GetUserByID(idInt)

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

func (bc BaseController) GetDeptTree(c *fiber.Ctx) error {
	var userService = service.UserService{
		DB: bc.DB,
	}

	result, err := userService.GetDeptTree()

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
