package envoygo

import "github.com/envoyproxy/envoy/contrib/golang/common/go/api"

type LogLevel api.LogType

const (
	DebugLevel LogLevel = iota + 1
	InfoLevel
	WarnLevel
	ErrorLevel
	CriticalLevel
)

type ConfigParser interface {
	api.StreamFilterConfigParser
}

func DefaultErrorHandler(c Context, err error) {
}
