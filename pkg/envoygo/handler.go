package envoygo

type ErrorHandler func(c Context, err error)

type Handler func(next HandlerFunc) HandlerFunc

type HandlerFunc func(c Context) error

func HandlerDecorator(next HandlerFunc) HandlerFunc {
	return func(c Context) error {
		if c.Committed() {
			// if committed, it means current chaining needs to be halted
			// and return status to the Envoy immediately
			return nil
		}

		return next(c)
	}
}
