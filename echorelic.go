package echorelic

import (
	"github.com/labstack/echo"
	newrelic "github.com/newrelic/go-agent"
)

type (
	EchoRelic struct {
		app newrelic.Application
	}
)

func New(app newrelic.Application) *EchoRelic {
	return &EchoRelic{
		app: app,
	}
}

func (er *EchoRelic) Process(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tx := er.app.StartTransaction(c.Request().URL.Path, c.Response(), c.Request())

		defer tx.End()

		c.Set("newrelic", tx)

		return next(c)
	}
}
