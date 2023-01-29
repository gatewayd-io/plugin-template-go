package plugin

import (
	"github.com/gatewayd-io/gatewayd-plugin-sdk/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// The following metrics are defined in the plugin and are used to
// track the number of times the plugin methods are called. These
// metrics are used to test the plugin metrics functionality.
var (
	OnConfigLoaded = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: metrics.Namespace,
		Help:      "The total number of calls to the onConfigLoaded method",
		Name:      "on_config_loaded_total",
	})
	OnTrafficFromClient = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: metrics.Namespace,
		Name:      "on_traffic_from_client_total",
		Help:      "The total number of of calls to the onTrafficFromClient method",
	})
	OnTrafficFromServer = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: metrics.Namespace,
		Name:      "on_traffic_from_server_total",
		Help:      "The total number of calls to the onTrafficFromServer method",
	})
)
