package controllers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Todo : todo model
type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos = []*Todo{
	{
		ID:        1,
		Title:     "Walk the dog 🦮",
		Completed: false,
	},
	{
		ID:        2,
		Title:     "Walk the cat 🐈",
		Completed: false,
	},
}

// GetTodos : get all todos
func GetTodos(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"todos": todos,
		},
	})
}

// CreateTodo : Create a todo
func CreateTodo(c *fiber.Ctx) error {
	type Request struct {
		Title string `json:"title"`
	}

	var body Request

	err := c.BodyParser(&body)

	// if error
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
		})
	}

	// create a todo variable
	todo := &Todo{
		ID:        len(todos) + 1,
		Title:     body.Title,
		Completed: false,
	}

	// append in todos
	todos = append(todos, todo)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"todo": todo,
		},
	})
}

// GetTodo : get a single todo
// PARAM: id
func GetTodo(c *fiber.Ctx) error {
	// get parameter value
	paramID := c.Params("id")

	// convert parameter value string to int
	id, err := strconv.Atoi(paramID)

	// if error in parsing string to int
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse Id",
		})
	}

	// find todo and return
	for _, todo := range todos {
		if todo.ID == id {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"success": true,
				"data": fiber.Map{
					"todo": todo,
				},
			})
		}
	}

	// if todo not available
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"success": false,
		"message": "Todo not found",
	})
}

// UpdateTodo : Update a todo
// PARAM: id
func UpdateTodo(c *fiber.Ctx) error {
	// find parameter
	paramID := c.Params("id")

	// convert parameter string to int
	id, err := strconv.Atoi(paramID)

	// if parameter cannot parse
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse id",
		})
	}

	// request structure
	type Request struct {
		Title     *string `json:"title"`
		Completed *bool   `json:"completed"`
	}

	var body Request
	err = c.BodyParser(&body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
		})
	}

	var todo *Todo

	for _, t := range todos {
		if t.ID == id {
			todo = t
			break
		}
	}

	if todo.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Not found",
		})
	}

	if body.Title != nil {
		todo.Title = *body.Title
	}

	if body.Completed != nil {
		todo.Completed = *body.Completed
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"todo": todo,
		},
	})
}

// DeleteTodo : Delete a todo
// PARAM: id
func DeleteTodo(c *fiber.Ctx) error {
	// get param
	paramID := c.Params("id")

	// convert param string to int
	id, err := strconv.Atoi(paramID)

	// if parameter cannot parse
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse id",
		})
	}

	// find and delete todo
	for i, todo := range todos {
		if todo.ID == id {

			todos = append(todos[:i], todos[i+1:]...)

			return c.SendStatus(fiber.StatusNoContent)
		}
	}

	// if todo not found
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"success": false,
		"message": "Todo not found",
	})
}
