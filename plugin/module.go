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
			v1.HookName_HOOK_NAME_ON_CONFIG_LOADED,
			v1.HookName_HOOK_NAME_ON_NEW_LOGGER,
			v1.HookName_HOOK_NAME_ON_NEW_POOL,
			v1.HookName_HOOK_NAME_ON_NEW_CLIENT,
			v1.HookName_HOOK_NAME_ON_NEW_PROXY,
			v1.HookName_HOOK_NAME_ON_NEW_SERVER,
			v1.HookName_HOOK_NAME_ON_SIGNAL,
			v1.HookName_HOOK_NAME_ON_RUN,
			v1.HookName_HOOK_NAME_ON_BOOTING,
			v1.HookName_HOOK_NAME_ON_BOOTED,
			v1.HookName_HOOK_NAME_ON_OPENING,
			v1.HookName_HOOK_NAME_ON_OPENED,
			v1.HookName_HOOK_NAME_ON_CLOSING,
			v1.HookName_HOOK_NAME_ON_CLOSED,
			v1.HookName_HOOK_NAME_ON_TRAFFIC,
			v1.HookName_HOOK_NAME_ON_TRAFFIC_FROM_CLIENT,
			v1.HookName_HOOK_NAME_ON_TRAFFIC_TO_SERVER,
			v1.HookName_HOOK_NAME_ON_TRAFFIC_FROM_SERVER,
			v1.HookName_HOOK_NAME_ON_TRAFFIC_TO_CLIENT,
			v1.HookName_HOOK_NAME_ON_SHUTDOWN,
			v1.HookName_HOOK_NAME_ON_TICK,
			// The following hook is invalid, and leads to an error in GatewayD,
			// but it'll be ignored. This is used for testing purposes. Feel free
			// to remove it in your plugin.
			v1.HookName(1000),
		},
		"tags":       []interface{}{"template", "plugin"},
		"categories": []interface{}{"template"},
	}
)
