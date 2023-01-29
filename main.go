package main

import (
	"flag"
	"os"
	"strconv"

	sdkConfig "github.com/gatewayd-io/gatewayd-plugin-sdk/config"
	"github.com/gatewayd-io/gatewayd-plugin-sdk/logging"
	"github.com/gatewayd-io/gatewayd-plugin-test/plugin"
	"github.com/hashicorp/go-hclog"
	goplugin "github.com/hashicorp/go-plugin"
	"github.com/mitchellh/mapstructure"
)

func main() {
	// Parse command line flags, passed by GatewayD via the plugin config
	logLevel := flag.String("log-level", "debug", "Log level")
	flag.Parse()

	logger := hclog.New(&hclog.LoggerOptions{
		Level:      logging.GetLogLevel(*logLevel),
		Output:     os.Stderr,
		JSONFormat: true,
		Color:      hclog.ColorOff,
	})

	pluginInstance := plugin.NewTestPlugin(plugin.Plugin{
		Logger: logger,
	})

	var config map[string]interface{}
	mapstructure.Decode(plugin.PluginConfig["config"], &config)
	if metricsEnabled, err := strconv.ParseBool(config["metricsEnabled"].(string)); err == nil {
		metricsConfig := plugin.MetricsConfig{
			Enabled:          metricsEnabled,
			UnixDomainSocket: config["metricsUnixDomainSocket"].(string),
			Endpoint:         config["metricsEndpoint"].(string),
		}
		go plugin.ExposeMetrics(metricsConfig, logger)
	}

	goplugin.Serve(&goplugin.ServeConfig{
		HandshakeConfig: goplugin.HandshakeConfig{
			ProtocolVersion:  1,
			MagicCookieKey:   sdkConfig.GetEnv("MAGIC_COOKIE_KEY", ""),
			MagicCookieValue: sdkConfig.GetEnv("MAGIC_COOKIE_VALUE", ""),
		},
		Plugins: goplugin.PluginSet{
			"gateway-plugin-test": pluginInstance,
		},
		GRPCServer: goplugin.DefaultGRPCServer,
		Logger:     logger,
	})
}
