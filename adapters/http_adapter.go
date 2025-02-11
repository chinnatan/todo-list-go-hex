package adapters

import (
	"strconv"
	"todo-list/core"

	"github.com/gofiber/fiber/v2"
)

type httpTodoHandler struct {
	service core.TodoService
}

func NewHttpTodoHandler(service core.TodoService) *httpTodoHandler {
	return &httpTodoHandler{service: service}
}

func (h *httpTodoHandler) CreateTodo(c *fiber.Ctx) error {
	var todo core.Todo
	if err := c.BodyParser(&todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	if err := h.service.CreateTodo(todo); err != nil {
		return c.Status((fiber.StatusInternalServerError)).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "todo created successfully"})
}

func (h *httpTodoHandler) GetAll(c *fiber.Ctx) error {
	items, err := h.service.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": items})
}

func (h *httpTodoHandler) GetById(c *fiber.Ctx) error {
	strId := c.Params("id")
	id, err := strconv.ParseInt(strId, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}

	item, err := h.service.GetById(id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": item})
}
