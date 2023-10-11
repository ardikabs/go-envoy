package envoygo

import (
	"fmt"
	"net/http"
	"time"
)

func requestHeaderModifier(headers map[string]string) Handler {
	return func(next HandlerFunc) HandlerFunc {
		return func(c Context) error {
			// c.Request()
			// 	-> c.Request().HeaderWriter().Get(k)
			// 	-> c.Request().HeaderWriter().Add(k,v)
			// 	-> c.Request().HeaderWriter().Set(k,v)
			// 	-> c.Request().HeaderWriter().Del(k)
			// 	-> c.Request().HeaderWriter().Values(k)
			//
			// 	-> c.Request().Http() -> *http.Request
			//
			// RequestPipeline <- FilterRequest
			// during request pipeline, if some process decide to commit, then it will halt the request pipeline
			// and immediately return back to the envoy
			// c.Response()
			// 	-> c.Response().HeaderWriter().Get(k)
			// 	-> c.Response().HeaderWriter().Add(k,v)
			// 	-> c.Response().HeaderWriter().Set(k,v)
			// 	-> c.Response().HeaderWriter().Del(k)
			// 	-> c.Response().HeaderWriter().Values(k)
			//
			// 	-> c.Response().Http() -> *http.Response
			//
			// ResponsePipeline <- FilterResponse
			// during response pipeline, if some process decide to commit, then it will halt the response pipeline
			// and immediately return back to the envoy

			for k, v := range headers {
				c.Request().HeaderWriter().Set(k, v)
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

			switch sc := c.Response().Http().StatusCode; sc {
			case http.StatusUnauthorized:
				return c.JSON(sc, nil, WithReplyResponseCodeDetails("unauthorized access"))
			}

			return err
		}
	}
}
