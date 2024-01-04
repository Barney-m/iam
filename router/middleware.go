package router

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func RequestLogging(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		t := time.Now()

		latency := time.Since(t)

		log.Info().Msg(fmt.Sprintf("%s %s %s %s\n",
			c.Request().Method,
			c.Request().RequestURI,
			c.Request().Proto,
			latency,
		))

		return next(c)
	}
}

func ResponseLogging(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Writer.Header().Set("X-Content-Type-Options", "nosniff")

		log.Info().Msg(fmt.Sprintf("%d %s %s\n",
			c.Response().Status,
			c.Request().Method,
			c.Request().RequestURI,
		))

		return next(c)
	}
}

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return errors.New("No token")
		}

		if strings.HasPrefix(authHeader, "Bearer ") {
			return errors.New("No token formed")
		}

		token := authHeader[7 : len(authHeader)-1]
		fmt.Println("test")
		fmt.Println(token)
		return next(c)
	}
}
