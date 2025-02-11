package main

import (
	"todo-list/adapters"
	"todo-list/core"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()

	// init db using sqlite
	db, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&core.Todo{})

	// setup
	todoRepo := adapters.NewGormTodoRepository(db)
	todoService := core.NewTodoService(todoRepo)
	todoHandler := adapters.NewHttpTodoHandler(todoService)

	// routes
	app.Post("/todo", todoHandler.CreateTodo)
	app.Get("/todos", todoHandler.GetAll)
	app.Get("/todo/:id", todoHandler.GetById)

	// start server
	app.Listen(":3000")
}
