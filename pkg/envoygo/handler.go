package envoygo

import (
	"fmt"
	"net/http"
	"time"
)

type Response interface {
	http.ResponseWriter

	Trailer() http.Header

	Status() int
}

type ErrorHandler func(c Context, err error)

type Handler func(next HandlerFunc) HandlerFunc

type HandlerFunc func(c Context) error

func requestHeaderModifier(headers map[string]string) Handler {
	return func(next HandlerFunc) HandlerFunc {
		return func(c Context) error {
			for k, v := range headers {
				c.Request().Header.Set(k, v)
			}

			return next(c)
		}
	}
}

func responseEnforcer() Handler {
	return func(next HandlerFunc) HandlerFunc {
		return func(c Context) error {
			start := time.Now()
			err := next(c)
			end := time.Since(start)
			c.Log(InfoLevel, fmt.Sprintf("execution requires: %s", end))

			switch sc := c.Response().Status(); sc {
			case http.StatusUnauthorized:
				return c.JSON(sc, nil, WithReplyResponseCodeDetails("unauthorized access"))
			}

			return err
		}
	}
}
