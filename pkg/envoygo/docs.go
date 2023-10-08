package envoygo

/*

plugin.Register(r Registerer)

http.RegisterHttpFilterConfigFactoryAndParser(r.Name(), factoryConstructor(r.Handlers()...), r.ConfigParser())

func factoryConstructor(handlers ...Handler) api.StreamFilterFactory {
	return func(c interface{}) api.StreamFilterFactory {
		return func(callbacks api.FilterCallbackHandler) api.StreamFilter {
			return &filter{
				callbacks: callbacks,
				config:    c,
				handlers: handlers,
			}
		}
	}
}

type Registerer interface {
	Name() string
	ConfigParser() api.StreamFilterConfigParser
	Handlers() []Handler
}

type Handler func(next HandlerFunc) HandlerFunc

type HandlerFunc func(c Context) error

type Context interface {
	Config() interface{}

	Log(lvl LogLevel, msg string)

	Request() *http.Request

	Response() http.ResponseWriter

	String(code int, s string, opts ...EnvoyReplyOption) error

	JSON(code int, i interface{}, opts ...EnvoyReplyOption) error

	NoContent(code int, opts ...EnvoyReplyOption) error
}

type EnvoyReplyOption func (o *EnvoyReplyOptions) error
func WithReplyResponseCodeDetails("suceeded_response") EnvoyReplyOption
func WithReplyGrpcStatusCode(-1) EnvoyReplyOption

func init() {
	app := plugin.New("go_http_plugin")

	app.SetErrorHandler(plugin.DefaultErrorHandler)
	app.SetConfigParser(&configParser{})

	app.Use(responseEnforcer)       // flow: incoming request 1		outcoming response 7
	app.Use(revocation.Middleware) 	// flow: incoming request 2		outcoming response 6
	app.Use(ipallowlistChecker)		// flow: incoming request 3		outcoming response 5
	app.Use(authn.Middleware)		// flow: incoming request 4		outcoming response 4
	app.Use(authz.Middleware)		// flow: incoming request 5		outcoming response 3
	app.Use(requestHeaderModifier)	// flow: incoming request 6		outcoming response 2
	app.Use(responseHeaderModifier)	// flow: incoming request 7		outcoming response 1

	app.Register()
}
*/
