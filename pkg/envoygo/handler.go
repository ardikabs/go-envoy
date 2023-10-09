package envoygo

type ErrorHandler func(c Context, err error)

type Handler func(next HandlerFunc) HandlerFunc

type HandlerFunc func(c Context) error
