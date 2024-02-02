package auth

import (
	"net/http"

	"github.com/mesxx/Echo_Todo_API/model"
	"github.com/mesxx/Echo_Todo_API/usecase"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type AuthMiddleware struct {
	UserUsecase usecase.UserUsecase
}

func NewAuthMiddleware(uu usecase.UserUsecase) *AuthMiddleware {
	return &AuthMiddleware{
		UserUsecase: uu,
	}
}

func (u *AuthMiddleware) Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		auth := c.Get("user").(*jwt.Token)
		claims := auth.Claims.(*model.CustomClaims)

		user, err := u.UserUsecase.FindByUsername(claims.Username)
		if err != nil || claims.Password != user.Password {
			return echo.NewHTTPError(http.StatusUnauthorized, "You dont have credentials!")
		}

		c.Set("user_auth", user)
		return next(c)
	}
}
