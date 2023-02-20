package plugin

import (
	v1 "github.com/gatewayd-io/gatewayd-plugin-sdk/plugin/v1"
	goplugin "github.com/hashicorp/go-plugin"
)

var (
	PluginID = v1.PluginID{
		Name:      "gatewayd-plugin-template",
		Version:   "0.0.1",
		RemoteUrl: "github.com/gatewayd-io/gatewayd-plugin-template",
	}
	PluginMap = map[string]goplugin.Plugin{
		"gatewayd-plugin-template": &TemplatePlugin{},
	}
	// TODO: Handle this in a better way
	// https://github.com/gatewayd-io/gatewayd-plugin-sdk/issues/3
	PluginConfig = map[string]interface{}{
		"id": map[string]interface{}{
			"name":      PluginID.Name,
			"version":   PluginID.Version,
			"remoteUrl": PluginID.RemoteUrl,
		},
		"description": "Template plugin",
		"authors": []interface{}{
			"Mostafa Moradian <mostafa@gatewayd.io>",
		},
		"license":    "Apache-2.0",
		"projectUrl": "https://github.com/gatewayd-io/gatewayd-plugin-template",
		// Compile-time configuration
		"config": map[string]interface{}{
			"metricsEnabled":          "true",
			"metricsUnixDomainSocket": "/tmp/gatewayd-plugin-template.sock",
			"metricsEndpoint":         "/metrics",
		},
		// "requires": []interface{}{
		// 	map[string]interface{}{
		// 		"name":      "gatewayd-plugin-non-existing",
		// 		"version":   "0.0.1",
		// 		"remoteUrl": "github.com/gatewayd-io/gatewayd-plugin-non-existing",
		// 	},
		// },
		// TODO: Use enum/constant for hooks
		"hooks": []interface{}{
			"onConfigLoaded",
			"onNewLogger",
			"onNewPool",
			"onNewClient",
			"onNewProxy",
			"onNewServer",
			"onSignal",
			"onRun",
			"onBooting",
			"onBooted",
			"onOpening",
			"onOpened",
			"onClosing",
			"onClosed",
			"onTraffic",
			"onShutdown",
			"onTick",
			"onTrafficFromClient",
			"onTrafficToServer",
			"onTrafficFromServer",
			"onTrafficToClient",
			// The following hook is invalid, and leads to an error in GatewayD,
			// but it'll be ignored. This is used for testing purposes. Feel free
			// to remove it in your plugin.
			"onPluginConfigLoaded",
		},
		"tags":       []interface{}{"template", "plugin"},
		"categories": []interface{}{"template"},
	}
)
