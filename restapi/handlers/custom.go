package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/etherlabsio/healthcheck/v2"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/yhsiang/f2pool-challenge/database"
)

func Metrics(handler http.Handler) http.Handler {
	return promhttp.Handler()
}

func Health(handler http.Handler) http.Handler {
	return healthcheck.Handler(
		// WithTimeout allows you to set a max overall timeout.
		healthcheck.WithTimeout(5*time.Second),

		healthcheck.WithChecker(
			"database", healthcheck.CheckerFunc(
				func(ctx context.Context) error {
					db := database.NewDB(0)
					return db.Ping()
				},
			),
		),
	)
}
