package handler

import (
	"net/http"
	"os"
	"time"

	"todo_api/model"
	"todo_api/usecase"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserUsecase usecase.UserUsecase
}

func NewUserHandler(uu usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		UserUsecase: uu,
	}
}

func (h *UserHandler) RegisterUser(c echo.Context) error {
	// Init struct for raw
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return err
	}

	hashedPassword, err := h.UserUsecase.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	if err := h.UserUsecase.Create(user); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, echo.Map{
		"message": "Success",
		"data":    user,
	})
}

func (h *UserHandler) GetAllUsers(c echo.Context) error {
	users, err := h.UserUsecase.FindAll()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Success",
		"data":    users,
	})
}

func (h *UserHandler) GetUser(c echo.Context) error {
	// Get Context user_auth from auth.go
	data, ok := c.Get("user_auth").(*model.User)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	user, err := h.UserUsecase.FindByID(uint(data.ID))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Success",
		"data":    user,
	})
}

func (h *UserHandler) GetTodo(c echo.Context) error {
	// Get Context user_auth from auth.go
	data, ok := c.Get("user_auth").(*model.User)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	user, err := h.UserUsecase.FindByID(uint(data.ID))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Success",
		"data":    user.Todo,
	})
}

func (h *UserHandler) UpdateUser(c echo.Context) error {
	// Get Context user_auth from auth.go
	data, ok := c.Get("user_auth").(*model.User)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	user, err := h.UserUsecase.FindByID(uint(data.ID))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "No Record")
	}

	// Init raw
	if err := c.Bind(user); err != nil {
		return err
	}

	// Query
	if err := h.UserUsecase.Update(user); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Success",
		"data":    user,
	})
}

func (h *UserHandler) DeleteUser(c echo.Context) error {
	// Get Context user_auth from auth.go
	data, ok := c.Get("user_auth").(*model.User)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	// Query
	if err := h.UserUsecase.Delete(uint(data.ID)); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Success",
	})
}

func (h *UserHandler) Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	user, err := h.UserUsecase.FindByUsername(username)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Wrong Username!")
	}
	if err := h.UserUsecase.CheckPasswordHash(password, user.Password); err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Wrong Password!")
	}

	// Set custom claims
	claims := &model.CustomClaims{
		user.ID,
		user.Username,
		user.Password,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Success",
		"token":   t,
	})
}
