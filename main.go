package main

import (
	"fmt"

	"github.com/gatewayd-io/gatewayd-plugin-test/plugin"
	goplugin "github.com/hashicorp/go-plugin"
)

func init() {
	fmt.Println("init")
}

func main() {
	goplugin.Serve(&goplugin.ServeConfig{
		HandshakeConfig: goplugin.HandshakeConfig{
			ProtocolVersion:  1,
			MagicCookieKey:   "GATEWAYD_PLUGIN",
			MagicCookieValue: "gatewayd",
		},
		Plugins: map[string]goplugin.Plugin{
			"test": &plugin.TestPlugin{Impl: plugin.GDPServiceServerImpl{}},
		},

		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: goplugin.DefaultGRPCServer,
	})
}
