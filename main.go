package main

import (
	"github.com/gatewayd-io/gatewayd-plugin-test/plugin"
	goplugin "github.com/hashicorp/go-plugin"
)

func main() {
	goplugin.Serve(&goplugin.ServeConfig{
		HandshakeConfig: plugin.Handshake,
		Plugins: goplugin.PluginSet{
			"gateway-plugin-test": plugin.NewTestPlugin(plugin.Plugin{}),
		},
		GRPCServer: goplugin.DefaultGRPCServer,
	})
}
