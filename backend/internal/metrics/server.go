// Package metrics provides Prometheus metrics server
package metrics

import (
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Listen starts metrics HTTP server
func Listen(address string) error {
	if os.Getenv("CI") == "true" {
		// В CI не запускаем метрики или на другом порту
		return nil
	}
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	return http.ListenAndServe(address, mux)
}
