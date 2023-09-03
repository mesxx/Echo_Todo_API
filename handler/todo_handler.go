package handler

import (
	"net/http"
	"strconv"

	"todo_api/model"
	"todo_api/usecase"

	"github.com/labstack/echo/v4"
)

type TodoHandler struct {
	TodoUsecase usecase.TodoUsecase
}

func NewTodoHandler(tu usecase.TodoUsecase) *TodoHandler {
	return &TodoHandler{
		TodoUsecase: tu,
	}
}

func (h *TodoHandler) CreateTodo(c echo.Context) error {
	// Init struct for raw
	todo := new(model.Todo)
	if err := c.Bind(todo); err != nil {
		return err
	}

	// Get Context user_auth from auth.go
	data, ok := c.Get("user_auth").(*model.User)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	todo.UserID = data.ID
	if err := h.TodoUsecase.Create(todo); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, echo.Map{
		"Message": "Success",
		"Data":    todo,
	})
}

func (h *TodoHandler) GetAllTodo(c echo.Context) error {
	todo, err := h.TodoUsecase.FindAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "No Record")
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Success",
		"data":    todo,
	})
}

func (h *TodoHandler) DeleteTodo(c echo.Context) error {
	// Params
	id, _ := strconv.Atoi(c.Param("id"))

	// Get Context user_auth from auth.go
	data, ok := c.Get("user_auth").(*model.User)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}
	todo, err := h.TodoUsecase.FindByID(uint(id))
	if err != nil || todo.UserID != data.ID {
		return echo.NewHTTPError(http.StatusBadRequest, "No Record")
	}

	// Query
	if err := h.TodoUsecase.Delete(todo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "No Record")
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Success",
	})
}
