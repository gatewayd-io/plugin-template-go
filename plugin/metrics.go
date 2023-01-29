package plugin

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const (
	Namespace = "gatewayd_plugin_test"
)

// The following metrics are defined in the plugin and are used to
// track the number of times the plugin methods are called. These
// metrics are used to test the plugin metrics functionality.
var (
	OnConfigLoaded = promauto.NewCounter(prometheus.CounterOpts{
		Name:      "on_config_loaded",
		Help:      "The total number of calls to the onConfigLoaded method",
		Namespace: Namespace,
	})
	OnTrafficFromClient = promauto.NewCounter(prometheus.CounterOpts{
		Name:      "on_traffic_from_client",
		Help:      "The total number of of calls to the onTrafficFromClient method",
		Namespace: Namespace,
	})
	OnTrafficFromServer = promauto.NewCounter(prometheus.CounterOpts{
		Name:      "on_traffic_from_server",
		Help:      "The total number of calls to the onTrafficFromServer method",
		Namespace: Namespace,
	})
)
