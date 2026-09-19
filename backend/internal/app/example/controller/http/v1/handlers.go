package v1

import (
	fiber "github.com/gofiber/fiber/v2"

	utilsjwt "skadi/backend/internal/pkg/utils/jwt"
)

// ExampleController represents a controller with example handlers to check auth.
type ExampleController struct{}

// NewController returns a new instance of ExampleController.
func NewController() *ExampleController {
	return &ExampleController{}
}

// @summary		Проверка. [Все]
// @description	Проверочный эндпоинт с доступом абсолютно всем.
// @router			/example/free [get]
// @id				example-free
// @tags			example
// @accept			json
// @produce		json
// @success		200	{object}	exampleData
func (c *ExampleController) Free(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(exampleData{
		Handler: "free",
		Access:  "any",
	})
}

// @summary		Проверка. [Только авторизованные]
// @description	Проверочный эндпоинт с доступом только для авторизованных юзеров.
// @router			/example/private [get]
// @id				example-private
// @tags			example
// @produce		json
// @security		JWTAccess
// @success		200	{object}	exampleData
// @failure		401	"Неверный токен (пустой, истекший или неверный формат)"
// @failure		403	"Юзер не авторизован, доступ запрещён"
func (c *ExampleController) Private(ctx *fiber.Ctx) error {
	userClaims := utilsjwt.ParseUserClaimsFromRequest(ctx)
	return ctx.Status(fiber.StatusOK).JSON(exampleData{
		UserClaims: userClaims,
		Handler:    "private",
		Access:     "auth user",
	})
}

// @summary		Проверка. [Только админ]
// @description	Проверочный эндпоинт с доступом только для админов.
// @router			/example/admin [get]
// @id				example-admin
// @tags			example
// @produce		json
// @security		JWTAccess
// @success		200	{object}	exampleData
// @failure		401	"Неверный токен (пустой, истекший или неверный формат)"
// @failure		403	"Юзер не админ, доступ запрещён"
func (c *ExampleController) Admin(ctx *fiber.Ctx) error {
	userClaims := utilsjwt.ParseUserClaimsFromRequest(ctx)
	return ctx.Status(fiber.StatusOK).JSON(exampleData{
		UserClaims: userClaims,
		Handler:    "admin",
		Access:     "admin only",
	})
}

// @summary		Проверка. [Только преподаватель]
// @description	Проверочный эндпоинт с доступом только для преподавателей.
// @router			/example/teacher [get]
// @id				example-teacher
// @tags			example
// @produce		json
// @security		JWTAccess
// @success		200	{object}	exampleData
// @failure		401	"Неверный токен (пустой, истекший или неверный формат)"
// @failure		403	"Юзер не преподаватель, доступ запрещён"
func (c *ExampleController) Teacher(ctx *fiber.Ctx) error {
	userClaims := utilsjwt.ParseUserClaimsFromRequest(ctx)
	return ctx.Status(fiber.StatusOK).JSON(exampleData{
		UserClaims: userClaims,
		Handler:    "teacher",
		Access:     "teacher only",
	})
}

// @summary		Проверка. [Только ученик]
// @description	Проверочный эндпоинт с доступом только для студентов.
// @router			/example/student [get]
// @id				example-student
// @tags			example
// @produce		json
// @security		JWTAccess
// @success		200	{object}	exampleData
// @failure		401	"Неверный токен (пустой, истекший или неверный формат)"
// @failure		403	"Юзер не студент, доступ запрещён"
func (c *ExampleController) Student(ctx *fiber.Ctx) error {
	userClaims := utilsjwt.ParseUserClaimsFromRequest(ctx)
	return ctx.Status(fiber.StatusOK).JSON(exampleData{
		UserClaims: userClaims,
		Handler:    "student",
		Access:     "student only",
	})
}
