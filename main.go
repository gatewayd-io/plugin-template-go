package main

import (
	"os"

	"github.com/gatewayd-io/gatewayd-plugin-test/plugin"
	"github.com/hashicorp/go-hclog"
	goplugin "github.com/hashicorp/go-plugin"
)

func main() {
	logger := hclog.New(&hclog.LoggerOptions{
		Level:      hclog.Debug,
		Output:     os.Stderr,
		JSONFormat: true,
		Color:      hclog.ColorOff,
	})
	pluginInstance := plugin.NewTestPlugin(plugin.Plugin{
		Logger: logger,
	})
	goplugin.Serve(&goplugin.ServeConfig{
		HandshakeConfig: goplugin.HandshakeConfig{
			ProtocolVersion:  1,
			MagicCookieKey:   "GATEWAYD_PLUGIN",
			MagicCookieValue: "5712b87aa5d7e9f9e9ab643e6603181c5b796015cb1c09d6f5ada882bf2a1872",
		},
		Plugins: goplugin.PluginSet{
			"gateway-plugin-test": pluginInstance,
		},
		GRPCServer: goplugin.DefaultGRPCServer,
		Logger:     logger,
	})
}
