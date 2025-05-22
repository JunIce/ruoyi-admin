package main

import (
	"net/http"
	"os"
	"runtime/debug"
	"time"
	v1 "wms-server/v1/controller/v1"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// var db *gorm.DB

func init() {
	env := os.Getenv("ENV")

	var err error
	switch env {
	case "production":
		err = godotenv.Load(".env.production")
	case "development":
		err = godotenv.Load(".env.development")
	case "local":
		err = godotenv.Load(".env.local")
	default:
		err = godotenv.Load(".env.development")
	}

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	log.SetLevel(log.LevelDebug)
	log.SetOutput(os.Stdout)
}

func main() {
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")

	dsn := DB_USER + ":" + DB_PASSWORD + "@tcp(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?charset=utf8mb4&parseTime=True&loc=Local"

	log.Info(dsn)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // data source name
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Log level
		NowFunc: func() time.Time {
			return time.Now().In(time.FixedZone("CST", 8*3600)) // 或者使用 time.LoadLocation("Asia/Shanghai")
		},
	})

	if err != nil {
		return
	}

	app := fiber.New()

	// 错误处理中间件
	app.Use(func(c *fiber.Ctx) error {
		// 尝试执行下一个中间件/处理器
		err := c.Next()
		if err != nil {
			// 输出错误堆栈
			log.Info("Error: %v\nStack:\n%s", err, debug.Stack())

			// 返回友好的错误消息给客户端
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": "Internal Server Error",
			})
		}
		return nil
	})

	system := v1.System{}

	v1Controller := v1.BaseController{DB: db}

	app.Get("/captcha", system.GenerateCaptchaHandler)
	app.Get("/captcha/:id.png", system.CaptchaImageHandler)
	app.Get("/getInfo", v1Controller.GetInfo)
	app.Get("/system/user/list", v1Controller.GetUserList)
	app.Get("/system/user/deptTree", v1Controller.GetDeptTree)
	app.Get("/system/user/:id", v1Controller.GetUserByID)
	app.Post("/system/user", v1Controller.PostUser)
	app.Put("/system/user", v1Controller.UpdateUser)
	app.Delete("/system/user/:id", v1Controller.DeleteUser)

	log.Info("Server is running on port :3000")
	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}
