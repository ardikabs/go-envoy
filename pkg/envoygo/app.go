package envoygo

import (
	"github.com/envoyproxy/envoy/contrib/golang/common/go/api"
	"github.com/envoyproxy/envoy/contrib/golang/filters/http/source/go/pkg/http"
)

func filterFactoryConstructor(p Plugin) api.StreamFilterConfigFactory {
	return func(config interface{}) api.StreamFilterFactory {
		return func(callbacks api.FilterCallbackHandler) api.StreamFilter {
			return &filter{
				callback: callbacks,
				config:   config,
				p:        p,
			}
		}
	}
}

type Plugin interface {
	Use(handler Handler)

	ErrorHandler() ErrorHandler

	HandlerChain() HandlerFunc

	Register()

	handleRequestHeaderStep(header api.RequestHeaderMap) api.StatusType

	handleResponseHeaderStep(header api.ResponseHeaderMap) api.StatusType
}

type plugin struct {
	name         string
	configParser ConfigParser
	errorHandler ErrorHandler

	handlers     []Handler
	handlerChain HandlerFunc

	ctx Context
}

func New(name string, configParser ConfigParser) Plugin {
	p := &plugin{
		name:         name,
		errorHandler: DefaultErrorHandler,
		configParser: configParser,
	}

	return p
}

func (p *plugin) Use(handler Handler) {
	if handler == nil {
		return
	}

	p.handlers = append(p.handlers, handler)
}

func (p *plugin) HandlerChain() HandlerFunc {
	return p.handlerChain
}

func (p *plugin) ErrorHandler() ErrorHandler {
	return p.errorHandler
}

func (p *plugin) Register() {
	p.handlerChain = func(c Context) error { return nil }

	for i := len(p.handlers) - 1; i >= 0; i-- {
		p.handlerChain = HandlerDecorator(p.handlers[i](p.handlerChain))
	}

	http.RegisterHttpFilterConfigFactoryAndParser(p.name, filterFactoryConstructor(p), p.configParser)
}

func (p *plugin) handleRequestHeaderStep(header api.RequestHeaderMap) api.StatusType {
	return api.Continue
}

func (p *plugin) handleResponseHeaderStep(header api.ResponseHeaderMap) api.StatusType {
	return api.Continue
}
