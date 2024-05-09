package gonvoy

import (
	"net/http"
	"testing"

	"github.com/envoyproxy/envoy/contrib/golang/common/go/api"
	"github.com/go-logr/logr"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHttpFilterHandlerManager_SetErrorHandler(t *testing.T) {
	t.Run("nil error handler will be ignored", func(t *testing.T) {
		mgr := &httpFilterHandlerManager{}

		mgr.SetErrorHandler(nil)
		assert.Nil(t, mgr.errorHandler)
	})

	t.Run("a custom error handler will be used", func(t *testing.T) {
		mgr := &httpFilterHandlerManager{}

		mgr.SetErrorHandler(DefaultHttpFilterErrorHandler)
		assert.NotNil(t, mgr.errorHandler)
	})
}

func TestHttpFilterHandlerManager_RegisterHandler(t *testing.T) {

	t.Run("execution order of all registerd handlers must follow a FIFO sequence", func(t *testing.T) {
		mockHandlerFirst := NewMockHttpFilterHandler(t)
		mockHandlerSecond := NewMockHttpFilterHandler(t)
		mockHandlerThird := NewMockHttpFilterHandler(t)

		mockHandlerFirst.EXPECT().Disable().Return(false)
		mockHandlerSecond.EXPECT().Disable().Return(false)
		mockHandlerThird.EXPECT().Disable().Return(false)

		t.Run("within OnRequestHeader", func(t *testing.T) {
			firstHandlerOnRequestHeader := mockHandlerFirst.EXPECT().OnRequestHeader(mock.Anything, mock.Anything).Return(nil)
			secondHandlerOnRequestHeader := mockHandlerSecond.EXPECT().OnRequestHeader(mock.Anything, mock.Anything).Return(nil)
			thirdHandlerOnRequestHeader := mockHandlerThird.EXPECT().OnRequestHeader(mock.Anything, mock.Anything).Return(nil)

			secondHandlerOnRequestHeader.NotBefore(firstHandlerOnRequestHeader.Call)
			thirdHandlerOnRequestHeader.NotBefore(firstHandlerOnRequestHeader.Call, secondHandlerOnRequestHeader.Call)

			mockContext := NewMockContext(t)
			mockContext.EXPECT().Request().Return(&http.Request{}).Maybe()
			mockContext.EXPECT().Committed().Return(false).Maybe()

			mockCtrl := NewMockHttpFilterPhaseController(t)
			mockCtrl.EXPECT().Handle(mock.Anything, mock.Anything).Run(func(c Context, proc HttpFilterProcessor) {
				_ = proc.HandleOnRequestHeader(c)
			}).Return(ActionPause, nil)

			mgr := &httpFilterHandlerManager{}
			mgr.RegisterHandler(mockHandlerFirst)
			mgr.RegisterHandler(mockHandlerSecond)
			mgr.RegisterHandler(mockHandlerThird)

			status := mgr.Serve(mockContext, mockCtrl)
			assert.Equal(t, api.StopAndBuffer, status)
		})

		t.Run("within OnRequestBody", func(t *testing.T) {
			firstHandlerOnRequestBody := mockHandlerFirst.EXPECT().OnRequestBody(mock.Anything, mock.Anything).Return(nil)
			secondHandlerOnRequestBody := mockHandlerSecond.EXPECT().OnRequestBody(mock.Anything, mock.Anything).Return(nil)
			thirdHandlerOnRequestBody := mockHandlerThird.EXPECT().OnRequestBody(mock.Anything, mock.Anything).Return(nil)

			secondHandlerOnRequestBody.NotBefore(firstHandlerOnRequestBody.Call)
			thirdHandlerOnRequestBody.NotBefore(firstHandlerOnRequestBody.Call, secondHandlerOnRequestBody.Call)

			mockContext := NewMockContext(t)
			mockContext.EXPECT().RequestBody().Return(&bodyWriter{}).Maybe()
			mockContext.EXPECT().Committed().Return(false).Maybe()

			mockCtrl := NewMockHttpFilterPhaseController(t)
			mockCtrl.EXPECT().Handle(mock.Anything, mock.Anything).Run(func(c Context, proc HttpFilterProcessor) {
				_ = proc.HandleOnRequestBody(c)
			}).Return(ActionPause, nil)

			mgr := &httpFilterHandlerManager{}
			mgr.RegisterHandler(mockHandlerFirst)
			mgr.RegisterHandler(mockHandlerSecond)
			mgr.RegisterHandler(mockHandlerThird)

			status := mgr.Serve(mockContext, mockCtrl)
			assert.Equal(t, api.StopAndBuffer, status)
		})

		t.Run("within OnResponseHeader", func(t *testing.T) {
			firstHandlerOnResponseHeader := mockHandlerFirst.EXPECT().OnResponseHeader(mock.Anything, mock.Anything).Return(nil)
			secondHandlerOnResponseHeader := mockHandlerSecond.EXPECT().OnResponseHeader(mock.Anything, mock.Anything).Return(nil)
			thirdHandlerOnResponseHeader := mockHandlerThird.EXPECT().OnResponseHeader(mock.Anything, mock.Anything).Return(nil)

			secondHandlerOnResponseHeader.NotBefore(firstHandlerOnResponseHeader.Call)
			thirdHandlerOnResponseHeader.NotBefore(firstHandlerOnResponseHeader.Call, secondHandlerOnResponseHeader.Call)

			mockContext := NewMockContext(t)
			mockContext.EXPECT().Response().Return(&http.Response{}).Maybe()
			mockContext.EXPECT().Committed().Return(false).Maybe()

			mockCtrl := NewMockHttpFilterPhaseController(t)
			mockCtrl.EXPECT().Handle(mock.Anything, mock.Anything).Run(func(c Context, proc HttpFilterProcessor) {
				_ = proc.HandleOnResponseHeader(c)
			}).Return(ActionPause, nil)

			mgr := &httpFilterHandlerManager{}
			mgr.RegisterHandler(mockHandlerFirst)
			mgr.RegisterHandler(mockHandlerSecond)
			mgr.RegisterHandler(mockHandlerThird)

			status := mgr.Serve(mockContext, mockCtrl)
			assert.Equal(t, api.StopAndBuffer, status)
		})

		t.Run("within OnResponseBody", func(t *testing.T) {
			firstHandlerOnResponseBody := mockHandlerFirst.EXPECT().OnResponseBody(mock.Anything, mock.Anything).Return(nil)
			secondHandlerOnResponseBody := mockHandlerSecond.EXPECT().OnResponseBody(mock.Anything, mock.Anything).Return(nil)
			thirdHandlerOnResponseBody := mockHandlerThird.EXPECT().OnResponseBody(mock.Anything, mock.Anything).Return(nil)

			secondHandlerOnResponseBody.NotBefore(firstHandlerOnResponseBody.Call)
			thirdHandlerOnResponseBody.NotBefore(firstHandlerOnResponseBody.Call, secondHandlerOnResponseBody.Call)

			mockContext := NewMockContext(t)
			mockContext.EXPECT().ResponseBody().Return(&bodyWriter{}).Maybe()
			mockContext.EXPECT().Committed().Return(false).Maybe()

			mockCtrl := NewMockHttpFilterPhaseController(t)
			mockCtrl.EXPECT().Handle(mock.Anything, mock.Anything).Run(func(c Context, proc HttpFilterProcessor) {
				_ = proc.HandleOnResponseBody(c)
			}).Return(ActionPause, nil)

			mgr := &httpFilterHandlerManager{}
			mgr.RegisterHandler(mockHandlerFirst)
			mgr.RegisterHandler(mockHandlerSecond)
			mgr.RegisterHandler(mockHandlerThird)

			status := mgr.Serve(mockContext, mockCtrl)
			assert.Equal(t, api.StopAndBuffer, status)
		})
	})

	t.Run("a nil handler won't be registered", func(t *testing.T) {
		mgr := &httpFilterHandlerManager{}
		createBadHandlerFn := func() *PassthroughHttpFilterHandler {
			return nil
		}

		mgr.RegisterHandler(createBadHandlerFn())
		assert.Nil(t, mgr.entrypoint)
	})

	t.Run("a disabled handler won't be registered", func(t *testing.T) {
		mgr := &httpFilterHandlerManager{}

		mockHandler := NewMockHttpFilterHandler(t)
		mockHandler.EXPECT().Disable().Return(true)

		mgr.RegisterHandler(mockHandler)
		assert.Nil(t, mgr.entrypoint)
	})
}

func TestHttpFilterHandlerManager_Serve(t *testing.T) {

	t.Run("serve and catch a panic", func(t *testing.T) {
		mockContext := NewMockContext(t)
		mockContext.EXPECT().Log().Return(logr.Logger{})
		mockContext.EXPECT().GetProperty(mock.Anything, mock.Anything).Return("", nil)
		mockContext.EXPECT().JSON(
			mock.MatchedBy(func(code int) bool {
				return assert.Equal(t, http.StatusInternalServerError, code)
			}),
			mock.MatchedBy(func(body []byte) bool {
				return assert.Equal(t, ResponseInternalServerError, body)
			}),
			mock.Anything,
			mock.Anything,
		).Return(nil)
		mockCtrl := NewMockHttpFilterPhaseController(t)
		mockCtrl.EXPECT().Handle(mock.Anything, mock.Anything).Panic("unexpected action")

		mgr := &httpFilterHandlerManager{
			errorHandler: DefaultHttpFilterErrorHandler,
			entrypoint:   newHttpFilterProcessor(PassthroughHttpFilterHandler{}),
		}

		status := mgr.Serve(mockContext, mockCtrl)
		assert.Equal(t, api.LocalReply, status)
	})

	t.Run("serve and got pause action", func(t *testing.T) {
		mockContext := NewMockContext(t)
		mockCtrl := NewMockHttpFilterPhaseController(t)
		mockCtrl.EXPECT().Handle(mock.Anything, mock.Anything).Return(ActionPause, nil)

		mgr := &httpFilterHandlerManager{
			errorHandler: DefaultHttpFilterErrorHandler,
			entrypoint:   newHttpFilterProcessor(PassthroughHttpFilterHandler{}),
		}

		status := mgr.Serve(mockContext, mockCtrl)
		assert.Equal(t, api.StopAndBuffer, status)
	})

	t.Run("serve with continue action", func(t *testing.T) {
		mockContext := NewMockContext(t)
		mockContext.EXPECT().StatusType().Return(api.Continue)
		mockCtrl := NewMockHttpFilterPhaseController(t)
		mockCtrl.EXPECT().Handle(mock.Anything, mock.Anything).Return(ActionContinue, nil)

		mgr := &httpFilterHandlerManager{
			errorHandler: DefaultHttpFilterErrorHandler,
			entrypoint:   newHttpFilterProcessor(PassthroughHttpFilterHandler{}),
		}

		status := mgr.Serve(mockContext, mockCtrl)
		assert.Equal(t, api.Continue, status)
	})

	t.Run("serve with no handler being registered", func(t *testing.T) {
		mockContext := NewMockContext(t)
		mockCtrl := NewMockHttpFilterPhaseController(t)

		mgr := &httpFilterHandlerManager{
			errorHandler: DefaultHttpFilterErrorHandler,
		}

		status := mgr.Serve(mockContext, mockCtrl)
		assert.Equal(t, api.Continue, status)
	})
}