package main

import (
	"github.com/labstack/echo/v4"
	"github.com/newrelic/go-agent/v3/integrations/nrecho-v4"
	"github.com/newrelic/go-agent/v3/newrelic"
	"log"
	"net/http"
	"os"
)

func main() {
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("test_app"),
		newrelic.ConfigLicense(os.Getenv("NEWRELIC_LICENSE_KEY")),
		newrelic.ConfigDistributedTracerEnabled(true),
		func(cfg *newrelic.Config) {
			cfg.ErrorCollector.IgnoreStatusCodes = []int{
				http.StatusBadRequest,
				http.StatusNotFound,
			}
		},
	)
	if err != nil {
		log.Fatalf("newrelic init app error: %s", err.Error())
	}

	e := echo.New()

	e.Use(nrecho.Middleware(app))
	e.GET("/", handler)

	e.Logger.Fatal(e.Start(":1323"))
}

func handler(c echo.Context) error {
	query := c.QueryParam("q")
	switch query {
	case "500":
		return c.String(http.StatusInternalServerError, "XD")
	case "200":
		return c.String(http.StatusOK, "ok")
	case "func":
		if err := BrokenFunc(); err != nil {
			nrecho.FromContext(c).NoticeError(err)
			return c.String(http.StatusInternalServerError, "XD")
		}
		return c.String(http.StatusOK, "ok")
	}
	return c.String(http.StatusBadRequest, "query is not passed e.g. q?=200")
}

func BrokenFunc() error {
	return newrelic.Error{
		Message:    "BrokenFunc is broken",
		Class:      "FuncError",
		Attributes: map[string]interface{}{},
		Stack:      newrelic.NewStackTrace(),
	}
}
