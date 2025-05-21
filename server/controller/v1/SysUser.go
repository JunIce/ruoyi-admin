package v1

import (
	"strconv"
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
func (bc BaseController) PostUser(c *fiber.Ctx) {

}
