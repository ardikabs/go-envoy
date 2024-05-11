package httpfilter

import (
	"fmt"
	"go-simple-extension/myfilter/handler"

	"github.com/ardikabs/gonvoy"
)

func init() {
	gonvoy.RunHttpFilter(Filter{}, gonvoy.ConfigOptions{
		BaseConfig:              new(Config),
		MetricPrefix:            "myfilter_",
		DisableStrictBodyAccess: true,
		EnableRequestBodyRead:   true,
		EnableResponseBodyRead:  true,
	})
}

type Filter struct{}

var _ gonvoy.HttpFilter = &Filter{}

func (f Filter) Name() string {
	return "myfilter"
}

func (f Filter) OnBegin(c gonvoy.RuntimeContext) error {
	fcfg := c.GetFilterConfig()
	cfg, ok := fcfg.(*Config)
	if !ok {
		return fmt.Errorf("unexpected configuration type %T, expecting %T", fcfg, cfg)
	}

	c.RegisterHTTPFilterHandler(&handler.HandlerOne{})
	c.RegisterHTTPFilterHandler(&handler.HandlerTwo{})
	c.RegisterHTTPFilterHandler(&handler.HandlerThree{RequestHeaders: cfg.RequestHeaders})
	return nil
}

func (f Filter) OnComplete(c gonvoy.RuntimeContext) error {
	c.Metrics().Counter("requests_total",
		"host", gonvoy.MustGetProperty(c, "request.host", "-"),
		"method", gonvoy.MustGetProperty(c, "request.method", "-"),
		"status_code", gonvoy.MustGetProperty(c, "response.code", "-"),
	).Increment(1)

	return nil
}
