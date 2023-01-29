package plugin

import (
	"context"
	"encoding/base64"

	plugin_v1 "github.com/gatewayd-io/gatewayd-plugin-test/plugin/v1"
	"github.com/hashicorp/go-hclog"
	goplugin "github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/structpb"
)

var PluginID = plugin_v1.PluginID{
	Name:      "gatewayd-plugin-test",
	Version:   "0.0.1",
	RemoteUrl: "github.com/gatewayd-io/gatewayd-plugin-test",
}

var PluginMap = map[string]goplugin.Plugin{
	"gatewayd-plugin-test": &TestPlugin{},
}

type Plugin struct {
	goplugin.GRPCPlugin
	plugin_v1.GatewayDPluginServiceServer
	Logger hclog.Logger
}

type TestPlugin struct {
	goplugin.NetRPCUnsupportedPlugin
	Impl Plugin
}

// NewTestPlugin returns a new instance of the TestPlugin.
func NewTestPlugin(impl Plugin) *TestPlugin {
	return &TestPlugin{
		NetRPCUnsupportedPlugin: goplugin.NetRPCUnsupportedPlugin{},
		Impl:                    impl,
	}
}

// GRPCServer registers the plugin with the gRPC server.
func (p *TestPlugin) GRPCServer(b *goplugin.GRPCBroker, s *grpc.Server) error {
	plugin_v1.RegisterGatewayDPluginServiceServer(s, &p.Impl)
	return nil
}

// GRPCClient returns the plugin client.
func (p *TestPlugin) GRPCClient(ctx context.Context, b *goplugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return plugin_v1.NewGatewayDPluginServiceClient(c), nil
}

// GetPluginConfig returns the plugin config.
func (p *Plugin) GetPluginConfig(
	ctx context.Context, req *structpb.Struct) (*structpb.Struct, error) {
	return structpb.NewStruct(map[string]interface{}{
		"id": map[string]interface{}{
			"name":      PluginID.Name,
			"version":   PluginID.Version,
			"remoteUrl": PluginID.RemoteUrl,
		},
		"description": "Test plugin",
		"authors": []interface{}{
			"Mostafa Moradian <mstfmoradian@gmail.com>",
		},
		"license":    "Apache-2.0",
		"projectUrl": "https://github.com/gatewayd-io/gatewayd-plugin-test",
		"config": map[string]interface{}{
			"key": "value",
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
	})
}

// OnConfigLoaded is called when the global config is loaded by GatewayD.
func (p *Plugin) OnConfigLoaded(
	ctx context.Context, req *structpb.Struct) (*structpb.Struct, error) {
	if req.Fields == nil {
		req.Fields = make(map[string]*structpb.Value)
	}
	req.Fields["loggers.default.level"] = &structpb.Value{
		Kind: &structpb.Value_StringValue{
			StringValue: "debug",
		},
	}
	req.Fields["loggers.default.noColor"] = &structpb.Value{
		Kind: &structpb.Value_BoolValue{
			BoolValue: false,
		},
	}
	return req, nil
}

// OnTrafficFromClient is called when a request is received by GatewayD from the client.
func (p *Plugin) OnTrafficFromClient(
	ctx context.Context, req *structpb.Struct) (*structpb.Struct, error) {
	request := req.Fields["request"].GetStringValue()
	if reqBytes, err := base64.StdEncoding.DecodeString(request); err == nil {
		p.Logger.Debug("OnIngressTraffic", "request", string(reqBytes))
	}

	return req, nil
}

// OnTrafficFromServer is called when a response is received by GatewayD from the server.
func (p *Plugin) OnTrafficFromServer(
	ctx context.Context, resp *structpb.Struct) (*structpb.Struct, error) {
	response := resp.Fields["response"].GetStringValue()
	if respBytes, err := base64.StdEncoding.DecodeString(response); err == nil {
		p.Logger.Debug("OnEgressTraffic", "response", string(respBytes))
	}
	return resp, nil
}
