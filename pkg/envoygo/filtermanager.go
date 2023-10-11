package envoygo

import (
	"github.com/envoyproxy/envoy/contrib/golang/common/go/api"
)

type filter struct {
	callback api.FilterCallbackHandler
	config   interface{}

	p Plugin
}

func newFilter(callback api.FilterCallbackHandler, config interface{}, p Plugin) api.StreamFilter {
	f := &filter{
		callback: callback,
		config:   config,
		p:        p,
	}

	/*
		start channel for request
		start channel for response
	*/

	return f
}

func (f *filter) DecodeHeaders(header api.RequestHeaderMap, endStream bool) api.StatusType {
	return f.p.handleRequestHeaderStep(header)
}

func (f *filter) EncodeHeaders(header api.ResponseHeaderMap, endStream bool) api.StatusType {
	return f.p.handleResponseHeaderStep(header)
}

func (f *filter) DecodeData(buffer api.BufferInstance, endStream bool) api.StatusType {
	return api.Continue
}

func (f *filter) EncodeData(buffer api.BufferInstance, endStream bool) api.StatusType {
	return api.Continue
}

func (f *filter) DecodeTrailers(trailers api.RequestTrailerMap) api.StatusType {
	return api.Continue
}

func (f *filter) EncodeTrailers(trailers api.ResponseTrailerMap) api.StatusType {
	return api.Continue
}

func (f *filter) OnDestroy(reason api.DestroyReason) {
}
