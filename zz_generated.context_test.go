// Code generated by mockery v2.36.1. DO NOT EDIT.

package gonvoy

import (
	http "net/http"

	api "github.com/envoyproxy/envoy/contrib/golang/common/go/api"

	logr "github.com/go-logr/logr"

	mock "github.com/stretchr/testify/mock"
)

// MockContext is an autogenerated mock type for the Context type
type MockContext struct {
	mock.Mock
}

type MockContext_Expecter struct {
	mock *mock.Mock
}

func (_m *MockContext) EXPECT() *MockContext_Expecter {
	return &MockContext_Expecter{mock: &_m.Mock}
}

// Committed provides a mock function with given fields:
func (_m *MockContext) Committed() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockContext_Committed_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Committed'
type MockContext_Committed_Call struct {
	*mock.Call
}

// Committed is a helper method to define mock.On call
func (_e *MockContext_Expecter) Committed() *MockContext_Committed_Call {
	return &MockContext_Committed_Call{Call: _e.mock.On("Committed")}
}

func (_c *MockContext_Committed_Call) Run(run func()) *MockContext_Committed_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockContext_Committed_Call) Return(_a0 bool) *MockContext_Committed_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockContext_Committed_Call) RunAndReturn(run func() bool) *MockContext_Committed_Call {
	_c.Call.Return(run)
	return _c
}

// GetCache provides a mock function with given fields:
func (_m *MockContext) GetCache() Cache {
	ret := _m.Called()

	var r0 Cache
	if rf, ok := ret.Get(0).(func() Cache); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Cache)
		}
	}

	return r0
}

// MockContext_GetCache_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCache'
type MockContext_GetCache_Call struct {
	*mock.Call
}

// GetCache is a helper method to define mock.On call
func (_e *MockContext_Expecter) GetCache() *MockContext_GetCache_Call {
	return &MockContext_GetCache_Call{Call: _e.mock.On("GetCache")}
}

func (_c *MockContext_GetCache_Call) Run(run func()) *MockContext_GetCache_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockContext_GetCache_Call) Return(_a0 Cache) *MockContext_GetCache_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockContext_GetCache_Call) RunAndReturn(run func() Cache) *MockContext_GetCache_Call {
	_c.Call.Return(run)
	return _c
}

// GetFilterConfig provides a mock function with given fields:
func (_m *MockContext) GetFilterConfig() interface{} {
	ret := _m.Called()

	var r0 interface{}
	if rf, ok := ret.Get(0).(func() interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// MockContext_GetFilterConfig_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetFilterConfig'
type MockContext_GetFilterConfig_Call struct {
	*mock.Call
}

// GetFilterConfig is a helper method to define mock.On call
func (_e *MockContext_Expecter) GetFilterConfig() *MockContext_GetFilterConfig_Call {
	return &MockContext_GetFilterConfig_Call{Call: _e.mock.On("GetFilterConfig")}
}

func (_c *MockContext_GetFilterConfig_Call) Run(run func()) *MockContext_GetFilterConfig_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockContext_GetFilterConfig_Call) Return(_a0 interface{}) *MockContext_GetFilterConfig_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockContext_GetFilterConfig_Call) RunAndReturn(run func() interface{}) *MockContext_GetFilterConfig_Call {
	_c.Call.Return(run)
	return _c
}

// GetProperty provides a mock function with given fields: name, defaultVal
func (_m *MockContext) GetProperty(name string, defaultVal string) (string, error) {
	ret := _m.Called(name, defaultVal)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (string, error)); ok {
		return rf(name, defaultVal)
	}
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(name, defaultVal)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(name, defaultVal)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockContext_GetProperty_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProperty'
type MockContext_GetProperty_Call struct {
	*mock.Call
}

// GetProperty is a helper method to define mock.On call
//   - name string
//   - defaultVal string
func (_e *MockContext_Expecter) GetProperty(name interface{}, defaultVal interface{}) *MockContext_GetProperty_Call {
	return &MockContext_GetProperty_Call{Call: _e.mock.On("GetProperty", name, defaultVal)}
}

func (_c *MockContext_GetProperty_Call) Run(run func(name string, defaultVal string)) *MockContext_GetProperty_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockContext_GetProperty_Call) Return(_a0 string, _a1 error) *MockContext_GetProperty_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockContext_GetProperty_Call) RunAndReturn(run func(string, string) (string, error)) *MockContext_GetProperty_Call {
	_c.Call.Return(run)
	return _c
}

// IsRequestBodyReadable provides a mock function with given fields:
func (_m *MockContext) IsRequestBodyReadable() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockContext_IsRequestBodyReadable_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsRequestBodyReadable'
type MockContext_IsRequestBodyReadable_Call struct {
	*mock.Call
}

// IsRequestBodyReadable is a helper method to define mock.On call
func (_e *MockContext_Expecter) IsRequestBodyReadable() *MockContext_IsRequestBodyReadable_Call {
	return &MockContext_IsRequestBodyReadable_Call{Call: _e.mock.On("IsRequestBodyReadable")}
}

func (_c *MockContext_IsRequestBodyReadable_Call) Run(run func()) *MockContext_IsRequestBodyReadable_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockContext_IsRequestBodyReadable_Call) Return(_a0 bool) *MockContext_IsRequestBodyReadable_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockContext_IsRequestBodyReadable_Call) RunAndReturn(run func() bool) *MockContext_IsRequestBodyReadable_Call {
	_c.Call.Return(run)
	return _c
}

// IsRequestBodyWriteable provides a mock function with given fields:
func (_m *MockContext) IsRequestBodyWriteable() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockContext_IsRequestBodyWriteable_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsRequestBodyWriteable'
type MockContext_IsRequestBodyWriteable_Call struct {
	*mock.Call
}

// IsRequestBodyWriteable is a helper method to define mock.On call
func (_e *MockContext_Expecter) IsRequestBodyWriteable() *MockContext_IsRequestBodyWriteable_Call {
	return &MockContext_IsRequestBodyWriteable_Call{Call: _e.mock.On("IsRequestBodyWriteable")}
}

func (_c *MockContext_IsRequestBodyWriteable_Call) Run(run func()) *MockContext_IsRequestBodyWriteable_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockContext_IsRequestBodyWriteable_Call) Return(_a0 bool) *MockContext_IsRequestBodyWriteable_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockContext_IsRequestBodyWriteable_Call) RunAndReturn(run func() bool) *MockContext_IsRequestBodyWriteable_Call {
	_c.Call.Return(run)
	return _c
}

// IsResponseBodyReadable provides a mock function with given fields:
func (_m *MockContext) IsResponseBodyReadable() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockContext_IsResponseBodyReadable_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsResponseBodyReadable'
type MockContext_IsResponseBodyReadable_Call struct {
	*mock.Call
}

// IsResponseBodyReadable is a helper method to define mock.On call
func (_e *MockContext_Expecter) IsResponseBodyReadable() *MockContext_IsResponseBodyReadable_Call {
	return &MockContext_IsResponseBodyReadable_Call{Call: _e.mock.On("IsResponseBodyReadable")}
}

func (_c *MockContext_IsResponseBodyReadable_Call) Run(run func()) *MockContext_IsResponseBodyReadable_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockContext_IsResponseBodyReadable_Call) Return(_a0 bool) *MockContext_IsResponseBodyReadable_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockContext_IsResponseBodyReadable_Call) RunAndReturn(run func() bool) *MockContext_IsResponseBodyReadable_Call {
	_c.Call.Return(run)
	return _c
}

// IsResponseBodyWriteable provides a mock function with given fields:
func (_m *MockContext) IsResponseBodyWriteable() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockContext_IsResponseBodyWriteable_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsResponseBodyWriteable'
type MockContext_IsResponseBodyWriteable_Call struct {
	*mock.Call
}

// IsResponseBodyWriteable is a helper method to define mock.On call
func (_e *MockContext_Expecter) IsResponseBodyWriteable() *MockContext_IsResponseBodyWriteable_Call {
	return &MockContext_IsResponseBodyWriteable_Call{Call: _e.mock.On("IsResponseBodyWriteable")}
}

func (_c *MockContext_IsResponseBodyWriteable_Call) Run(run func()) *MockContext_IsResponseBodyWriteable_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockContext_IsResponseBodyWriteable_Call) Return(_a0 bool) *MockContext_IsResponseBodyWriteable_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockContext_IsResponseBodyWriteable_Call) RunAndReturn(run func() bool) *MockContext_IsResponseBodyWriteable_Call {
	_c.Call.Return(run)
	return _c
}

// JSON provides a mock function with given fields: code, b, header, opts
func (_m *MockContext) JSON(code int, b []byte, header http.Header, opts ...ReplyOption) error {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, code, b, header)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, []byte, http.Header, ...ReplyOption) error); ok {
		r0 = rf(code, b, header, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockContext_JSON_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'JSON'
type MockContext_JSON_Call struct {
	*mock.Call
}

// JSON is a helper method to define mock.On call
//   - code int
//   - b []byte
//   - header http.Header
//   - opts ...ReplyOption
func (_e *MockContext_Expecter) JSON(code interface{}, b interface{}, header interface{}, opts ...interface{}) *MockContext_JSON_Call {
	return &MockContext_JSON_Call{Call: _e.mock.On("JSON",
		append([]interface{}{code, b, header}, opts...)...)}
}

func (_c *MockContext_JSON_Call) Run(run func(code int, b []byte, header http.Header, opts ...ReplyOption)) *MockContext_JSON_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]ReplyOption, len(args)-3)
		for i, a := range args[3:] {
			if a != nil {
				variadicArgs[i] = a.(ReplyOption)
			}
		}
		run(args[0].(int), args[1].([]byte), args[2].(http.Header), variadicArgs...)
	})
	return _c
}

func (_c *MockContext_JSON_Call) Return(_a0 error) *MockContext_JSON_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockContext_JSON_Call) RunAndReturn(run func(int, []byte, http.Header, ...ReplyOption) error) *MockContext_JSON_Call {
	_c.Call.Return(run)
	return _c
}

// Log provides a mock function with given fields:
func (_m *MockContext) Log() logr.Logger {
	ret := _m.Called()

	var r0 logr.Logger
	if rf, ok := ret.Get(0).(func() logr.Logger); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(logr.Logger)
	}

	return r0
}

// MockContext_Log_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Log'
type MockContext_Log_Call struct {
	*mock.Call
}

// Log is a helper method to define mock.On call
func (_e *MockContext_Expecter) Log() *MockContext_Log_Call {
	return &MockContext_Log_Call{Call: _e.mock.On("Log")}
}

func (_c *MockContext_Log_Call) Run(run func()) *MockContext_Log_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockContext_Log_Call) Return(_a0 logr.Logger) *MockContext_Log_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockContext_Log_Call) RunAndReturn(run func() logr.Logger) *MockContext_Log_Call {
	_c.Call.Return(run)
	return _c
}

// Metrics provides a mock function with given fields:
func (_m *MockContext) Metrics() Metrics {
	ret := _m.Called()

	var r0 Metrics
	if rf, ok := ret.Get(0).(func() Metrics); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Metrics)
		}
	}

	return r0
}

// MockContext_Metrics_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Metrics'
type MockContext_Metrics_Call struct {
	*mock.Call
}

// Metrics is a helper method to define mock.On call
func (_e *MockContext_Expecter) Metrics() *MockContext_Metrics_Call {
	return &MockContext_Metrics_Call{Call: _e.mock.On("Metrics")}
}

func (_c *MockContext_Metrics_Call) Run(run func()) *MockContext_Metrics_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockContext_Metrics_Call) Return(_a0 Metrics) *MockContext_Metrics_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockContext_Metrics_Call) RunAndReturn(run func() Metrics) *MockContext_Metrics_Call {
	_c.Call.Return(run)
	return _c
}

// Request provides a mock function with given fields:
func (_m *MockContext) Request() *http.Request {
	ret := _m.Called()

	var r0 *http.Request
	if rf, ok := ret.Get(0).(func() *http.Request); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Request)
		}
	}

	return r0
}

// MockContext_Request_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Request'
type MockContext_Request_Call struct {
	*mock.Call
}

// Request is a helper method to define mock.On call
func (_e *MockContext_Expecter) Request() *MockContext_Request_Call {
	return &MockContext_Request_Call{Call: _e.mock.On("Request")}
}

func (_c *MockContext_Request_Call) Run(run func()) *MockContext_Request_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockContext_Request_Call) Return(_a0 *http.Request) *MockContext_Request_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockContext_Request_Call) RunAndReturn(run func() *http.Request) *MockContext_Request_Call {
	_c.Call.Return(run)
	return _c
}

// RequestBody provides a mock function with given fields:
func (_m *MockContext) RequestBody() Body {
	ret := _m.Called()

	var r0 Body
	if rf, ok := ret.Get(0).(func() Body); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Body)
		}
	}

	return r0
}

// MockContext_RequestBody_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RequestBody'
type MockContext_RequestBody_Call struct {
	*mock.Call
}

// RequestBody is a helper method to define mock.On call
func (_e *MockContext_Expecter) RequestBody() *MockContext_RequestBody_Call {
	return &MockContext_RequestBody_Call{Call: _e.mock.On("RequestBody")}
}

func (_c *MockContext_RequestBody_Call) Run(run func()) *MockContext_RequestBody_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockContext_RequestBody_Call) Return(_a0 Body) *MockContext_RequestBody_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockContext_RequestBody_Call) RunAndReturn(run func() Body) *MockContext_RequestBody_Call {
	_c.Call.Return(run)
	return _c
}

// RequestHeader provides a mock function with given fields:
func (_m *MockContext) RequestHeader() Header {
	ret := _m.Called()

	var r0 Header
	if rf, ok := ret.Get(0).(func() Header); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Header)
		}
	}

	return r0
}

// MockContext_RequestHeader_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RequestHeader'
type MockContext_RequestHeader_Call struct {
	*mock.Call
}

// RequestHeader is a helper method to define mock.On call
func (_e *MockContext_Expecter) RequestHeader() *MockContext_RequestHeader_Call {
	return &MockContext_RequestHeader_Call{Call: _e.mock.On("RequestHeader")}
}

func (_c *MockContext_RequestHeader_Call) Run(run func()) *MockContext_RequestHeader_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockContext_RequestHeader_Call) Return(_a0 Header) *MockContext_RequestHeader_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockContext_RequestHeader_Call) RunAndReturn(run func() Header) *MockContext_RequestHeader_Call {
	_c.Call.Return(run)
	return _c
}

// Response provides a mock function with given fields:
func (_m *MockContext) Response() *http.Response {
	ret := _m.Called()

	var r0 *http.Response
	if rf, ok := ret.Get(0).(func() *http.Response); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	return r0
}

// MockContext_Response_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Response'
type MockContext_Response_Call struct {
	*mock.Call
}

// Response is a helper method to define mock.On call
func (_e *MockContext_Expecter) Response() *MockContext_Response_Call {
	return &MockContext_Response_Call{Call: _e.mock.On("Response")}
}

func (_c *MockContext_Response_Call) Run(run func()) *MockContext_Response_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockContext_Response_Call) Return(_a0 *http.Response) *MockContext_Response_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockContext_Response_Call) RunAndReturn(run func() *http.Response) *MockContext_Response_Call {
	_c.Call.Return(run)
	return _c
}

// ResponseBody provides a mock function with given fields:
func (_m *MockContext) ResponseBody() Body {
	ret := _m.Called()

	var r0 Body
	if rf, ok := ret.Get(0).(func() Body); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Body)
		}
	}

	return r0
}

// MockContext_ResponseBody_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ResponseBody'
type MockContext_ResponseBody_Call struct {
	*mock.Call
}

// ResponseBody is a helper method to define mock.On call
func (_e *MockContext_Expecter) ResponseBody() *MockContext_ResponseBody_Call {
	return &MockContext_ResponseBody_Call{Call: _e.mock.On("ResponseBody")}
}

func (_c *MockContext_ResponseBody_Call) Run(run func()) *MockContext_ResponseBody_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockContext_ResponseBody_Call) Return(_a0 Body) *MockContext_ResponseBody_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockContext_ResponseBody_Call) RunAndReturn(run func() Body) *MockContext_ResponseBody_Call {
	_c.Call.Return(run)
	return _c
}

// ResponseHeader provides a mock function with given fields:
func (_m *MockContext) ResponseHeader() Header {
	ret := _m.Called()

	var r0 Header
	if rf, ok := ret.Get(0).(func() Header); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Header)
		}
	}

	return r0
}

// MockContext_ResponseHeader_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ResponseHeader'
type MockContext_ResponseHeader_Call struct {
	*mock.Call
}

// ResponseHeader is a helper method to define mock.On call
func (_e *MockContext_Expecter) ResponseHeader() *MockContext_ResponseHeader_Call {
	return &MockContext_ResponseHeader_Call{Call: _e.mock.On("ResponseHeader")}
}

func (_c *MockContext_ResponseHeader_Call) Run(run func()) *MockContext_ResponseHeader_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockContext_ResponseHeader_Call) Return(_a0 Header) *MockContext_ResponseHeader_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockContext_ResponseHeader_Call) RunAndReturn(run func() Header) *MockContext_ResponseHeader_Call {
	_c.Call.Return(run)
	return _c
}

// SetRequestBody provides a mock function with given fields: buffer, endStream
func (_m *MockContext) SetRequestBody(buffer api.BufferInstance, endStream bool) {
	_m.Called(buffer, endStream)
}

// MockContext_SetRequestBody_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetRequestBody'
type MockContext_SetRequestBody_Call struct {
	*mock.Call
}

// SetRequestBody is a helper method to define mock.On call
//   - buffer api.BufferInstance
//   - endStream bool
func (_e *MockContext_Expecter) SetRequestBody(buffer interface{}, endStream interface{}) *MockContext_SetRequestBody_Call {
	return &MockContext_SetRequestBody_Call{Call: _e.mock.On("SetRequestBody", buffer, endStream)}
}

func (_c *MockContext_SetRequestBody_Call) Run(run func(buffer api.BufferInstance, endStream bool)) *MockContext_SetRequestBody_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(api.BufferInstance), args[1].(bool))
	})
	return _c
}

func (_c *MockContext_SetRequestBody_Call) Return() *MockContext_SetRequestBody_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockContext_SetRequestBody_Call) RunAndReturn(run func(api.BufferInstance, bool)) *MockContext_SetRequestBody_Call {
	_c.Call.Return(run)
	return _c
}

// SetRequestHeader provides a mock function with given fields: _a0
func (_m *MockContext) SetRequestHeader(_a0 api.RequestHeaderMap) {
	_m.Called(_a0)
}

// MockContext_SetRequestHeader_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetRequestHeader'
type MockContext_SetRequestHeader_Call struct {
	*mock.Call
}

// SetRequestHeader is a helper method to define mock.On call
//   - _a0 api.RequestHeaderMap
func (_e *MockContext_Expecter) SetRequestHeader(_a0 interface{}) *MockContext_SetRequestHeader_Call {
	return &MockContext_SetRequestHeader_Call{Call: _e.mock.On("SetRequestHeader", _a0)}
}

func (_c *MockContext_SetRequestHeader_Call) Run(run func(_a0 api.RequestHeaderMap)) *MockContext_SetRequestHeader_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(api.RequestHeaderMap))
	})
	return _c
}

func (_c *MockContext_SetRequestHeader_Call) Return() *MockContext_SetRequestHeader_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockContext_SetRequestHeader_Call) RunAndReturn(run func(api.RequestHeaderMap)) *MockContext_SetRequestHeader_Call {
	_c.Call.Return(run)
	return _c
}

// SetResponseBody provides a mock function with given fields: buffer, endStream
func (_m *MockContext) SetResponseBody(buffer api.BufferInstance, endStream bool) {
	_m.Called(buffer, endStream)
}

// MockContext_SetResponseBody_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetResponseBody'
type MockContext_SetResponseBody_Call struct {
	*mock.Call
}

// SetResponseBody is a helper method to define mock.On call
//   - buffer api.BufferInstance
//   - endStream bool
func (_e *MockContext_Expecter) SetResponseBody(buffer interface{}, endStream interface{}) *MockContext_SetResponseBody_Call {
	return &MockContext_SetResponseBody_Call{Call: _e.mock.On("SetResponseBody", buffer, endStream)}
}

func (_c *MockContext_SetResponseBody_Call) Run(run func(buffer api.BufferInstance, endStream bool)) *MockContext_SetResponseBody_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(api.BufferInstance), args[1].(bool))
	})
	return _c
}

func (_c *MockContext_SetResponseBody_Call) Return() *MockContext_SetResponseBody_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockContext_SetResponseBody_Call) RunAndReturn(run func(api.BufferInstance, bool)) *MockContext_SetResponseBody_Call {
	_c.Call.Return(run)
	return _c
}

// SetResponseHeader provides a mock function with given fields: _a0
func (_m *MockContext) SetResponseHeader(_a0 api.ResponseHeaderMap) {
	_m.Called(_a0)
}

// MockContext_SetResponseHeader_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetResponseHeader'
type MockContext_SetResponseHeader_Call struct {
	*mock.Call
}

// SetResponseHeader is a helper method to define mock.On call
//   - _a0 api.ResponseHeaderMap
func (_e *MockContext_Expecter) SetResponseHeader(_a0 interface{}) *MockContext_SetResponseHeader_Call {
	return &MockContext_SetResponseHeader_Call{Call: _e.mock.On("SetResponseHeader", _a0)}
}

func (_c *MockContext_SetResponseHeader_Call) Run(run func(_a0 api.ResponseHeaderMap)) *MockContext_SetResponseHeader_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(api.ResponseHeaderMap))
	})
	return _c
}

func (_c *MockContext_SetResponseHeader_Call) Return() *MockContext_SetResponseHeader_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockContext_SetResponseHeader_Call) RunAndReturn(run func(api.ResponseHeaderMap)) *MockContext_SetResponseHeader_Call {
	_c.Call.Return(run)
	return _c
}

// SkipNextPhase provides a mock function with given fields:
func (_m *MockContext) SkipNextPhase() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockContext_SkipNextPhase_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SkipNextPhase'
type MockContext_SkipNextPhase_Call struct {
	*mock.Call
}

// SkipNextPhase is a helper method to define mock.On call
func (_e *MockContext_Expecter) SkipNextPhase() *MockContext_SkipNextPhase_Call {
	return &MockContext_SkipNextPhase_Call{Call: _e.mock.On("SkipNextPhase")}
}

func (_c *MockContext_SkipNextPhase_Call) Run(run func()) *MockContext_SkipNextPhase_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockContext_SkipNextPhase_Call) Return(_a0 error) *MockContext_SkipNextPhase_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockContext_SkipNextPhase_Call) RunAndReturn(run func() error) *MockContext_SkipNextPhase_Call {
	_c.Call.Return(run)
	return _c
}

// StatusType provides a mock function with given fields:
func (_m *MockContext) StatusType() api.StatusType {
	ret := _m.Called()

	var r0 api.StatusType
	if rf, ok := ret.Get(0).(func() api.StatusType); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(api.StatusType)
	}

	return r0
}

// MockContext_StatusType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StatusType'
type MockContext_StatusType_Call struct {
	*mock.Call
}

// StatusType is a helper method to define mock.On call
func (_e *MockContext_Expecter) StatusType() *MockContext_StatusType_Call {
	return &MockContext_StatusType_Call{Call: _e.mock.On("StatusType")}
}

func (_c *MockContext_StatusType_Call) Run(run func()) *MockContext_StatusType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockContext_StatusType_Call) Return(_a0 api.StatusType) *MockContext_StatusType_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockContext_StatusType_Call) RunAndReturn(run func() api.StatusType) *MockContext_StatusType_Call {
	_c.Call.Return(run)
	return _c
}

// StreamInfo provides a mock function with given fields:
func (_m *MockContext) StreamInfo() api.StreamInfo {
	ret := _m.Called()

	var r0 api.StreamInfo
	if rf, ok := ret.Get(0).(func() api.StreamInfo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(api.StreamInfo)
		}
	}

	return r0
}

// MockContext_StreamInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StreamInfo'
type MockContext_StreamInfo_Call struct {
	*mock.Call
}

// StreamInfo is a helper method to define mock.On call
func (_e *MockContext_Expecter) StreamInfo() *MockContext_StreamInfo_Call {
	return &MockContext_StreamInfo_Call{Call: _e.mock.On("StreamInfo")}
}

func (_c *MockContext_StreamInfo_Call) Run(run func()) *MockContext_StreamInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockContext_StreamInfo_Call) Return(_a0 api.StreamInfo) *MockContext_StreamInfo_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockContext_StreamInfo_Call) RunAndReturn(run func() api.StreamInfo) *MockContext_StreamInfo_Call {
	_c.Call.Return(run)
	return _c
}

// String provides a mock function with given fields: code, s, header, opts
func (_m *MockContext) String(code int, s string, header http.Header, opts ...ReplyOption) error {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, code, s, header)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, string, http.Header, ...ReplyOption) error); ok {
		r0 = rf(code, s, header, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockContext_String_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'String'
type MockContext_String_Call struct {
	*mock.Call
}

// String is a helper method to define mock.On call
//   - code int
//   - s string
//   - header http.Header
//   - opts ...ReplyOption
func (_e *MockContext_Expecter) String(code interface{}, s interface{}, header interface{}, opts ...interface{}) *MockContext_String_Call {
	return &MockContext_String_Call{Call: _e.mock.On("String",
		append([]interface{}{code, s, header}, opts...)...)}
}

func (_c *MockContext_String_Call) Run(run func(code int, s string, header http.Header, opts ...ReplyOption)) *MockContext_String_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]ReplyOption, len(args)-3)
		for i, a := range args[3:] {
			if a != nil {
				variadicArgs[i] = a.(ReplyOption)
			}
		}
		run(args[0].(int), args[1].(string), args[2].(http.Header), variadicArgs...)
	})
	return _c
}

func (_c *MockContext_String_Call) Return(_a0 error) *MockContext_String_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockContext_String_Call) RunAndReturn(run func(int, string, http.Header, ...ReplyOption) error) *MockContext_String_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockContext creates a new instance of MockContext. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockContext(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockContext {
	mock := &MockContext{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
