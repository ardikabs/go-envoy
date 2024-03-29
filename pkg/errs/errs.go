package errs

import "errors"

var (
	ErrInternalServer = errors.New("Internal Server Error")
	ErrBadRequest     = errors.New("Bad Request")
	ErrUnauthorized   = errors.New("Unauthorized")
	ErrAccessDenied   = errors.New("RBAC: Access Denied")

	ErrPanic = errors.New("panic occurred")
)

func Unwrap(err error) error {
	if unwrapErr := errors.Unwrap(err); unwrapErr != nil {
		return unwrapErr
	}

	return err
}
