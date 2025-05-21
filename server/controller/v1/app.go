package v1

import (
	"wms-server/v1/model"

	"github.com/gofiber/fiber/v2"
)

func (bc BaseController) GetInfo(c *fiber.Ctx) error {
	var user model.SysUserVo
	bc.DB.Omit("Password", "PwdTime").First(&user, 1)

	return c.JSON(fiber.Map{
		"code": 200,
		"msg":  "success",
		"user": user,
	})
}
