package main

import (
	"net/http"

	"github.com/ardikabs/go-envoy"
)

type HandlerThree struct {
	RequestHeaders map[string]string
}

func (h *HandlerThree) RequestHandler(next envoy.HandlerFunc) envoy.HandlerFunc {
	return func(c envoy.Context) error {
		for k, v := range h.RequestHeaders {
			c.RequestHeader().Set(k, v)
		}

		return next(c)
	}
}

func (h *HandlerThree) ResponseHandler(next envoy.HandlerFunc) envoy.HandlerFunc {
	return func(c envoy.Context) error {
		switch sc := c.Response().StatusCode; sc {
		case http.StatusUnauthorized:
			return c.JSON(sc, envoy.CreateSimpleJSONBody("UNAUTHORIZED", "UNAUTHORIZED"), envoy.ToFlatHeader(c.Response().Header))
		case http.StatusTooManyRequests:
			return c.JSON(sc, envoy.CreateSimpleJSONBody("TOO_MANY_REQUESTS", "TOO_MANY_REQUESTS"), envoy.ToFlatHeader(c.Response().Header), envoy.WithResponseCodeDetails("rate limit exceeded"))
		case http.StatusServiceUnavailable:
			return c.JSON(sc, envoy.CreateSimpleJSONBody("SERVICE_UNAVAILABLE", "SERVICE_UNAVAILABLE"), envoy.ToFlatHeader(c.Response().Header), envoy.WithResponseCodeDetails("service unavailable"))
		}

		return next(c)
	}
}
