package prometheus

import (
	"fmt"

	"contrib.go.opencensus.io/exporter/prometheus"
	prom "github.com/prometheus/client_golang/prometheus"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/stats/view"
)

// NewExporter creates a new, configured Prometheus exporter.
func NewExporter(namespace string) (*prometheus.Exporter, error) {
	exporter, err := prometheus.NewExporter(prometheus.Options{
		Namespace: namespace,
		Registry:  prom.DefaultGatherer.(*prom.Registry),
	})
	if err != nil {
		return nil, fmt.Errorf("create prometheus exporter: %w", err)
	}

	view.RegisterExporter(exporter)

	// Register stat views
	if err := view.Register(ochttp.DefaultServerViews...); err != nil {
		return nil, fmt.Errorf("failed to register server views for HTTP metrics: %v", err)
	}
	return exporter, nil
}
