package web

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	AuthJwtKey         = "authJwtClaims"
	JwtTokenCookieName = "auth_token"
)

type AuthJwtClaims struct {
	AccountId string `json:"accountId"`
	jwt.RegisteredClaims
}

func NewAuthJwtMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		// xxx: ハードコーディング
		SigningKey: []byte("secret"),
		TokenLookupFuncs: []middleware.ValuesExtractor{
			func(c echo.Context) ([]string, error) {
				cookie, err := c.Cookie(JwtTokenCookieName)
				if err != nil {
					return nil, err
				}
				return []string{cookie.Value}, nil
			},
		},
		ContextKey: AuthJwtKey,
		NewClaimsFunc: func(c echo.Context) jwt.Claims {

			return new(AuthJwtClaims)
		},
	})
}
