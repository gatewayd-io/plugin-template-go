package plugin

import (
	pluginV1 "github.com/gatewayd-io/gatewayd-plugin-test/plugin/v1"
	goplugin "github.com/hashicorp/go-plugin"
)

var (
	PluginID = pluginV1.PluginID{
		Name:      "gatewayd-plugin-test",
		Version:   "0.0.1",
		RemoteUrl: "github.com/gatewayd-io/gatewayd-plugin-test",
	}
	PluginMap = map[string]goplugin.Plugin{
		"gatewayd-plugin-test": &TestPlugin{},
	}
	// TODO: Handle this in a better way
	// https://github.com/gatewayd-io/gatewayd-plugin-sdk/issues/3
	PluginConfig = map[string]interface{}{
		"id": map[string]interface{}{
			"name":      PluginID.Name,
			"version":   PluginID.Version,
			"remoteUrl": PluginID.RemoteUrl,
		},
		"description": "Test plugin",
		"authors": []interface{}{
			"Mostafa Moradian <mostafa@gatewayd.io>",
		},
		"license":    "Apache-2.0",
		"projectUrl": "https://github.com/gatewayd-io/gatewayd-plugin-test",
		// Compile-time configuration
		"config": map[string]interface{}{
			"metricsEnabled":          "true",
			"metricsUnixDomainSocket": "/tmp/gatewayd-plugin-test.sock",
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
			"onPluginConfigLoaded", // This leads to an error and will be ignored
			"onTrafficFromClient",
			"onTrafficFromServer",
		},
		"tags":       []interface{}{"test", "plugin"},
		"categories": []interface{}{"test"},
	}
)
