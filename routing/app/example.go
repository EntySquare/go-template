package app

import (
	"github.com/gofiber/fiber/v2"
	"go-template/pkg"
	"go-template/routing/types"
)

func Example(c *fiber.Ctx) error {
	var msgReq types.Example
	err := c.BodyParser(&msgReq)
	if err != nil {
		return c.JSON(pkg.MessageResponse(pkg.MESSAGE_FAIL, "can not transfer request to struct", "请求参数错误"))
	}
	ex := types.ExampleRes{
		Phone: "13666666666",
	}
	return c.JSON(pkg.SuccessResponse(ex))
}
