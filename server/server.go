package server

import (
	"os"

	"github.com/mesxx/Echo_Todo_API/auth"
	"github.com/mesxx/Echo_Todo_API/config"
	"github.com/mesxx/Echo_Todo_API/handler"
	"github.com/mesxx/Echo_Todo_API/model"
	"github.com/mesxx/Echo_Todo_API/repository"
	"github.com/mesxx/Echo_Todo_API/usecase"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewServer() *echo.Echo {
	godotenv.Load()
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Database initialization
	db, err := config.NewDB()
	if err != nil {
		e.Logger.Fatal(err)
	}
	db.AutoMigrate(&model.User{}, &model.Todo{})

	// Repository initialization
	ur := repository.NewUserRepository(db)
	tr := repository.NewTodoRepository(db)

	// Usecase initialization
	uu := usecase.NewUserUsecase(*ur)
	tu := usecase.NewTodoUsecase(*tr)

	// Handler initialization
	hu := handler.NewUserHandler(*uu)
	ht := handler.NewTodoHandler(*tu)

	// Auth initialization
	au := auth.NewAuthMiddleware(*uu)

	// Public Route
	e.POST("/user", hu.RegisterUser)
	e.POST("/user/login", hu.Login)
	e.GET("/users", hu.GetAllUsers)
	e.GET("/todo", ht.GetAllTodo)

	// Auth Route
	ua := e.Group("")
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(model.CustomClaims)
		},
		SigningKey: []byte(os.Getenv("SECRET_KEY")),
	}
	ua.Use(echojwt.WithConfig(config))
	ua.GET("/user", hu.GetUser, au.Authenticate)
	ua.GET("/user/todo", hu.GetTodo, au.Authenticate)
	ua.PUT("/user", hu.UpdateUser, au.Authenticate)
	ua.DELETE("/user", hu.DeleteUser, au.Authenticate)
	ua.POST("/todo", ht.CreateTodo, au.Authenticate)
	ua.DELETE("/todo/:id", ht.DeleteTodo, au.Authenticate)

	return e
}
