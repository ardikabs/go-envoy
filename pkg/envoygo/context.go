package envoygo

import (
	"github.com/envoyproxy/envoy/contrib/golang/common/go/api"
)

type Context interface {
	Config() interface{}

	Log(lvl LogLevel, msg string)

	Request() Request

	SetRequest()

	Response() Response

	SetResponse()

	StreamInfo() api.StreamInfo

	String(code int, s string, opts ...ContextOption) error

	JSON(code int, i interface{}, opts ...ContextOption) error

	NoContent(code int, opts ...ContextOption) error

	StatusType() api.StatusType

	Committed() bool
}

type ContextOption func(o *ContextOptions)

type ContextOptions struct {
	ResponseCodeDetail string
	GrpcStatusCode     int64
}

func WithReplyResponseCodeDetails(s string) ContextOption {
	return func(o *ContextOptions) {
		o.ResponseCodeDetail = s
	}
}

func WithReplyRpcStatusCode(code int64) ContextOption {
	return func(o *ContextOptions) {
		o.GrpcStatusCode = code
	}
}
