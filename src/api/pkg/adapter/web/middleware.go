package web

import (
	"log"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
		TokenLookupFuncs: []middleware.ValuesExtractor{
			func(c echo.Context) ([]string, error) {
				log.Println("TokenLookupFuncs")
				cookie, err := c.Cookie("auth_token")
				if err != nil {
					return nil, err
				}
				return []string{cookie.Value}, nil
			},
		},
		/*
			ParseTokenFunc: func(c echo.Context, auth string) (interface{}, error) {
				log.Println("ParseTokenFunc")
				return jwt.Parse(auth, func(token *jwt.Token) (interface{}, error) {
					return []byte("secret"), nil
				})
			},
			KeyFunc: func(t *jwt.Token) (interface{}, error) {
				log.Println("KeyFunc")
				return []byte("secret"), nil
			},
			NewClaimsFunc: func(c echo.Context) jwt.Claims {
				log.Println("NewClaimsFunc")
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
		*/
	})
}
