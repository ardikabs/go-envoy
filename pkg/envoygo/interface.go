package envoygo

import (
	"net/http"

	"github.com/ardikabs/go-envoy/pkg/types"
	"github.com/envoyproxy/envoy/contrib/golang/common/go/api"
)

type Request interface {
	HeaderWriter() HeaderWriter

	Http() *http.Request

	Buffer() api.BufferInstance
}

type request struct {
	headerWriter HeaderWriter

	httpReq *http.Request
}

func NewHttpRequest(header api.RequestHeaderMap) Request {
	httpReq, err := types.NewRequest(
		header.Method(), header.Host(),
		types.WithRequestURI(header.Path()),
		types.WithRequestHeaderRangeSetter(header),
	)

	if err != nil {
		panic(err)
	}

	r := &request{
		headerWriter: &headerWriter{header},
		httpReq:      httpReq,
	}
	return r
}

func (r *request) HeaderWriter() HeaderWriter {
	return r.headerWriter
}

func (r *request) Http() *http.Request {
	return r.httpReq
}

func (r *request) Buffer() api.BufferInstance {
	panic("NOT IMPLEMENTED")
}

type Response interface {
	HeaderWriter() HeaderWriter

	Http() *http.Response

	Buffer() api.BufferInstance
}

type response struct {
	headerWriter HeaderWriter

	httpResp *http.Response
}

func NewHttpResponse(header api.ResponseHeaderMap) Response {
	code, ok := header.Status()
	if !ok {
		return nil
	}

	httpResp, err := types.NewResponse(code, types.WithResponseHeaderRangeSetter(header))
	if err != nil {
		return nil
	}
	if err != nil {
		panic(err)
	}

	r := &response{
		headerWriter: &headerWriter{header},
		httpResp:     httpResp,
	}
	return r
}

func (r *response) HeaderWriter() HeaderWriter {
	return r.headerWriter
}

func (r *response) Http() *http.Response {
	return r.httpResp
}

func (r *response) Buffer() api.BufferInstance {
	panic("NOT IMPLEMENTED")
}
