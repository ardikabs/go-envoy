package envoygo

import (
	"net/http"
)

type Context interface {
	Config() interface{}

	Log(lvl LogLevel, msg string)

	Request() *http.Request

	Response() Response

	String(code int, s string, opts ...ReplyOption) error

	JSON(code int, i interface{}, opts ...ReplyOption) error

	NoContent(code int, opts ...ReplyOption) error
}

type ReplyOption func(o *ReplyOptions)

type ReplyOptions struct {
	ResponseCodeDetail string
	GrpcStatusCode     int64
}

func WithReplyResponseCodeDetails(s string) ReplyOption {
	return func(o *ReplyOptions) {
		o.ResponseCodeDetail = s
	}
}

func WithReplyRpcStatusCode(code int64) ReplyOption {
	return func(o *ReplyOptions) {
		o.GrpcStatusCode = code
	}
}
