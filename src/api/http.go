package api

import (
	"catalog/domain"
	"errors"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	productService domain.Service
	loginService   domain.LoginService
}

func (h *handler) Get(ctx *fiber.Ctx) error {
	code := ctx.Params("code")
	p, err := h.productService.Find(code)
	if err != nil {
		return ctx.Status(404).JSON(nil)
	}
	return ctx.JSON(&p)
}

func (h *handler) Post(ctx *fiber.Ctx) error {
	p := &domain.Product{}
	if err := ctx.BodyParser(&p); err != nil {
		return ctx.Status(500).JSON(nil)
	}
	err := h.productService.Store(p)
	if err != nil {
		return ctx.Status(500).JSON(nil)
	}
	return ctx.JSON(&p)
}

func (h *handler) Put(ctx *fiber.Ctx) error {
	p := &domain.Product{}
	if err := ctx.BodyParser(&p); err != nil {
		return ctx.Status(500).JSON(nil)
	}
	err := h.productService.Update(p)
	if err != nil {
		return ctx.Status(500).JSON(nil)
	}
	return ctx.JSON(&p)
}

func (h *handler) Delete(ctx *fiber.Ctx) error {
	code := ctx.Params("code")
	err := h.productService.Delete(code)
	if err != nil {
		return ctx.Status(500).JSON(nil)
	}
	return ctx.Status(201).JSON(fiber.Map{"message": "ok"})
}

func (h *handler) GetAll(ctx *fiber.Ctx) error {
	p, err := h.productService.FindAll()
	if err != nil {
		return ctx.Status(404).JSON(nil)
	}
	return ctx.JSON(&p)
}

func (h *handler) Login(ctx *fiber.Ctx) error {
	var r domain.Login
	if err := ctx.BodyParser(&r); err != nil {
		return ctx.Status(400).JSON(nil)
	}
	token, err := h.loginService.Login(r)
	if err != nil {
		if errors.Is(err, domain.ErrInvalidLogin) {
			return ctx.Status(401).JSON(nil)
		}
		return ctx.Status(500).JSON(nil)
	}
	return ctx.JSON(fiber.Map{"token": token})
}

// NewHandler  New handler instantiates a http handler for our product service
func NewHandler(productService domain.Service, loginService domain.LoginService) *handler {
	return &handler{productService: productService, loginService: loginService}
}
