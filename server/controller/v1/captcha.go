package v1

import (
	"github.com/gofiber/fiber/v2/log"

	"github.com/dchest/captcha"
	"github.com/gofiber/fiber/v2"
)

type System struct {
}

const (
	captchaIdLen  = 8
	captchaWidth  = 240
	captchaHeight = 80
)

// GenerateCaptchaHandler 获取验证码
func (e System) GenerateCaptchaHandler(c *fiber.Ctx) error {
	newCaptcha := captcha.New()

	log.Info("DriverDigitFunc answer: %s", newCaptcha)

	return c.JSON(fiber.Map{
		"code": 200,
		"data": newCaptcha,
		"msg":  "success",
	})
}

func (e System) CaptchaImageHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	log.Info("id: %s", id)
	// return c.WriteString()

	// captcha.Server(captcha.StdWidth, captcha.StdHeight).ServeHTTP(c.Context(), id)
	return c.JSON(fiber.Map{
		"code": 200,
		"data": id,
		"msg":  "success",
	})
}
