package web

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

const (
	AuthJwtKey = "authJwtClaims"
)

type AuthJwtClaims struct {
	AccountId string
	jwt.RegisteredClaims
}

func NewAuthJwtMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		// xxx: ハードコーディング
		SigningKey: []byte("secret"),
		ContextKey: "authJwtClaims",
		NewClaimsFunc: func(c echo.Context) jwt.Claims {

			cookie, err := c.Cookie("authJwtClaims")
			if err != nil {
				return nil
			}

			cookieValue := cookie.Value

			claims := &AuthJwtClaims{}

			jwt.ParseWithClaims(cookieValue, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte("secret"), nil
			})

			return claims
		},
	})
}
