package echorelic

import (
	"github.com/labstack/echo"
	newrelic "github.com/newrelic/go-agent"
)

func Monitor(app newrelic.Application) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tx := app.StartTransaction(c.Request().URL.Path, c.Response(), c.Request())
			defer tx.End()

			c.Set("newrelic", tx)

			return next(c)
		}
	}
}
