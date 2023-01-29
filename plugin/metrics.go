package plugin

import (
	"net"
	"net/http"
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	Namespace = "gatewayd"
)

// The following metrics are defined in the plugin and are used to
// track the number of times the plugin methods are called. These
// metrics are used to test the plugin metrics functionality.
var (
	OnConfigLoaded = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: Namespace,
		Help:      "The total number of calls to the onConfigLoaded method",
		Name:      "on_config_loaded_total",
	})
	OnTrafficFromClient = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: Namespace,
		Name:      "on_traffic_from_client_total",
		Help:      "The total number of of calls to the onTrafficFromClient method",
	})
	OnTrafficFromServer = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: Namespace,
		Name:      "on_traffic_from_server_total",
		Help:      "The total number of calls to the onTrafficFromServer method",
	})
)

type MetricsConfig struct {
	Enabled          bool
	UnixDomainSocket string
	Endpoint         string
}

func ExposeMetrics(metricsConfig MetricsConfig, logger hclog.Logger) {
	logger.Info(
		"Starting metrics server via HTTP over Unix domain socket",
		"unixDomainSocket", metricsConfig.UnixDomainSocket,
		"endpoint", metricsConfig.Endpoint)

	if err := os.Remove(metricsConfig.UnixDomainSocket); err != nil {
		logger.Error("Failed to remove unix domain socket")
	}

	listener, err := net.Listen("unix", metricsConfig.UnixDomainSocket)
	if err != nil {
		logger.Error("Failed to start metrics server")
	}

	if err := http.Serve(listener, promhttp.Handler()); err != nil {
		logger.Error("Failed to start metrics server")
	}
}
