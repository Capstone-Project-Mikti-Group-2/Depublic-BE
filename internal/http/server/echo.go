package server

import (
	"net/http"

	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/common"
	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/internal/config"
	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/internal/http/binder"
	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/internal/http/router"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	*echo.Echo
}

func NewServer(
	cfg *config.Config,
	binder *binder.Binder,
	publicRoutes, privateRoutes []*router.Route) *Server {

	e := echo.New()
	e.Binder = binder

	e.Use(
		middleware.Logger(),
		middleware.Recover(),
		middleware.CORS(),
	)
	//group API baru
	v1 := e.Group("/api/v1")
	for _, public := range publicRoutes {
		v1.Add(public.Method, public.Path, public.Handler)
	}

	for _, private := range privateRoutes {
		v1.Add(private.Method, private.Path, private.Handler, JWTProtected(cfg.JWT.SecretKey), RBACMiddleware(private.Roles...)) //JWTProtected(cfg.JWT.SecretKey)
	}
	e.GET("/ping", func(c echo.Context) error {
		return c.String(200, "pong")
	})
	return &Server{e}
}

func JWTProtected(secretKey string) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(common.JwtCustomClaims)
		},
		SigningKey: []byte(secretKey),
	})
}

func RBACMiddleware(roles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user, ok := c.Get("user").(*jwt.Token)
			if !ok {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"message": "Silahkan login terlebih dahulu",
				})
			}
			claims := user.Claims.(*common.JwtCustomClaims)

			//cek apakah user mempunyai role yang diinginkan
			if !contains(roles, claims.Role) {
				return c.JSON(http.StatusForbidden, map[string]interface{}{
					"message": "Anda tidak diizinkan untuk mengakses resource ini",
				})
			}
			return next(c)
		}
	}
}

func contains(slice []string, s string) bool {
	for _, value := range slice {
		if value == s {
			return true
		}
	}
	return false
}
