package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"go-template/database"
	"go-template/routing"
	"time"
)

func main() {
	database.ConnectDB()
	fiberApp := fiber.New()
	// 创建一个速率限制器，每秒最多只允许10个请求
	fiberApp.Use(limiter.New(limiter.Config{
		Max:        10,
		Expiration: 2 * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP() // 使用客户端IP作为限流key
		},
	}))
	// 添加 CORS 中间件
	fiberApp.Use(func(c *fiber.Ctx) error {
		// 允许所有域名进行跨域请求
		c.Set("Access-Control-Allow-Origin", "*")
		// 允许 GET、POST、PUT、DELETE 和 OPTIONS 方法进行跨域请求
		c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		// 允许客户端发送的请求头
		c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization, token")
		// 在响应中添加 CORS 头
		if c.Method() == "OPTIONS" {
			return c.SendStatus(fiber.StatusOK)
		} else {
			return c.Next()
		}
	})
	// 将速率限制器添加到路由中间件中
	fiberApp.Use(cors.New())
	routing.Setup(fiberApp)
	err := fiberApp.Listen(":3000")
	if err != nil {
		fmt.Println(err.Error())
	}
}
