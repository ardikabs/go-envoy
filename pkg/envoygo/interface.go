package envoygo

import (
	"net/http"

	"github.com/envoyproxy/envoy/contrib/golang/common/go/api"
)

type Request interface {
	HeaderWriter() HeaderWriter

	Http() *http.Request

	Buffer() api.BufferInstance

	Trailer() http.Header
}

type Response interface {
	HeaderWriter() HeaderWriter

	Http() *http.Response

	Buffer() api.BufferInstance

	Trailer() http.Header

	Status() int
}
