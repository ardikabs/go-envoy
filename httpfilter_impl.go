package gonvoy

import "github.com/envoyproxy/envoy/contrib/golang/common/go/api"

var _ api.StreamFilter = &httpFilterImpl{}

// httpFilterImpl is an HTTP Filter implementation for Envoy.
type httpFilterImpl struct {
	srv HttpFilterServer
}

func (f *httpFilterImpl) OnLog() { f.srv.Finalize() }

func (f *httpFilterImpl) OnDestroy(reason api.DestroyReason) { f.srv = nil }

func (f *httpFilterImpl) DecodeHeaders(header api.RequestHeaderMap, endStream bool) api.StatusType {
	decoder := f.decodeHeaders(header)
	res := f.srv.ServeDecodeFilter(decoder)
	return res.Status
}

func (f *httpFilterImpl) DecodeData(buffer api.BufferInstance, endStream bool) api.StatusType {
	decoder := f.decodeData(buffer, endStream)
	res := f.srv.ServeDecodeFilter(decoder)
	return res.Status
}

func (f *httpFilterImpl) EncodeHeaders(header api.ResponseHeaderMap, endStream bool) api.StatusType {
	encoder := f.encodeHeaders(header)
	res := f.srv.ServeEncodeFilter(encoder)
	return res.Status
}

func (f *httpFilterImpl) EncodeData(buffer api.BufferInstance, endStream bool) api.StatusType {
	encoder := f.encodeData(buffer, endStream)
	res := f.srv.ServeEncodeFilter(encoder)
	return res.Status
}

func (f *httpFilterImpl) decodeHeaders(header api.RequestHeaderMap) HttpFilterDecoderFunc {
	return func(c Context, p HttpFilterDecodeProcessor) (HttpFilterAction, error) {
		c.SetRequestHeader(header)

		if err := p.HandleOnRequestHeader(c); err != nil {
			return ActionContinue, err
		}

		if c.IsRequestBodyWriteable() {
			return ActionPause, nil
		}

		return ActionContinue, nil
	}
}

func (f *httpFilterImpl) decodeData(buffer api.BufferInstance, endStream bool) HttpFilterDecoderFunc {
	return func(c Context, p HttpFilterDecodeProcessor) (HttpFilterAction, error) {
		if !isRequestBodyAccessible(c) {
			return ActionSkip, nil
		}

		if buffer.Len() > 0 {
			c.SetRequestBody(buffer)
		}

		if endStream {
			return ActionContinue, p.HandleOnRequestBody(c)
		}

		return ActionPause, nil
	}
}

func (f *httpFilterImpl) encodeHeaders(header api.ResponseHeaderMap) HttpFilterEncoderFunc {
	return func(c Context, p HttpFilterEncodeProcessor) (HttpFilterAction, error) {
		c.SetResponseHeader(header)

		if err := p.HandleOnResponseHeader(c); err != nil {
			return ActionContinue, err
		}

		// During the Encode phases or HTTP Response flows,
		// if a user needs access to the HTTP Response Body, whether for reading or writing,
		// the EncodeHeaders phase should return with ActionPause (StopAndBuffer status) status.
		// This is necessary because the Response Header must be buffered in Envoy's filter-manager.
		// If this buffering is not done, the Response Header might be sent to the downstream client prematurely,
		// preventing the filter from returning a custom error response in case of unexpected events during processing.
		//
		// Hence, we opted for the IsResponseBodyReadable check instead of IsResponseBodyWritable.
		// It's worth noting that the behavior differs in the Decode phase because the stream flow is directed towards the upstream.
		// This means that even if DecodeHeaders has returned with ActionContinue (Continue status),
		// DecodeData is still under supervision within Envoy's filter-manager state.
		if c.IsResponseBodyReadable() {
			return ActionPause, nil
		}

		return ActionContinue, nil
	}
}

// Attention! Please be mindful of the Listener or Cluster per_connection_buffer_limit_bytes value
// when enabling the response body access on ConfigOptions (EnableResponseBodyRead or EnableResponseBodyWrite).
// The default value set by Envoy is 1MB. If the response body size exceeds this limit, the process will be halted.
// Although it's unclear whether this is considered a bug or a limitation at present,
// the Envoy Golang HTTP Filter library currently returns a 413 status code with a PayloadTooLarge message in such cases.
// Code references: https://github.com/envoyproxy/envoy/blob/v1.29.4/contrib/golang/filters/http/source/processor_state.cc#L362-L371.
func (f *httpFilterImpl) encodeData(buffer api.BufferInstance, endStream bool) HttpFilterEncoderFunc {
	return func(c Context, p HttpFilterEncodeProcessor) (HttpFilterAction, error) {
		if !isResponseBodyAccessible(c) {
			return ActionSkip, nil
		}

		if buffer.Len() > 0 {
			c.SetResponseBody(buffer)
		}

		if endStream {
			return ActionContinue, p.HandleOnResponseBody(c)
		}

		return ActionPause, nil
	}
}

func (*httpFilterImpl) DecodeTrailers(api.RequestTrailerMap) api.StatusType  { return api.Continue }
func (*httpFilterImpl) EncodeTrailers(api.ResponseTrailerMap) api.StatusType { return api.Continue }
func (*httpFilterImpl) OnLogDownstreamPeriodic()                             {}
func (*httpFilterImpl) OnLogDownstreamStart()                                {}