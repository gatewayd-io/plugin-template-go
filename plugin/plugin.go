package plugin

import (
	"context"

	v1 "github.com/gatewayd-io/gatewayd-plugin-test/plugin/v1"
	goplugin "github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/structpb"
)

var PluginID = v1.PluginID{
	Name:      "gatewayd-plugin-test",
	Version:   "0.0.1",
	RemoteUrl: "github.com/gatewayd-io/gatewayd-plugin-test",
}

var PluginMap = map[string]goplugin.Plugin{
	"gatewayd-plugin-test": &TestPlugin{},
}

type Plugin struct {
	v1.GatewayDPluginServiceServer
}
type TestPlugin struct {
	goplugin.NetRPCUnsupportedPlugin
	Impl Plugin
}

func NewTestPlugin(impl Plugin) *TestPlugin {
	return &TestPlugin{
		NetRPCUnsupportedPlugin: goplugin.NetRPCUnsupportedPlugin{},
		Impl:                    impl,
	}
}

func (p *TestPlugin) GRPCServer(b *goplugin.GRPCBroker, s *grpc.Server) error {
	v1.RegisterGatewayDPluginServiceServer(s, &p.Impl)
	return nil
}

func (p *TestPlugin) GRPCClient(ctx context.Context, b *goplugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return v1.NewGatewayDPluginServiceClient(c), nil
}

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
		// TODO: Use enum/constant for hooks
		"hooks":      []interface{}{"onConfigLoaded", "onPluginConfigLoaded"},
		"tags":       []interface{}{"test", "plugin"},
		"categories": []interface{}{"test"},
	})
}

func (p *Plugin) OnConfigLoaded(
	ctx context.Context, req *structpb.Struct) (*structpb.Struct, error) {
	if req.Fields == nil {
		req.Fields = make(map[string]*structpb.Value)
	}
	req.Fields["loggers.logger.level"] = &structpb.Value{
		Kind: &structpb.Value_StringValue{
			StringValue: "debug",
		},
	}
	req.Fields["loggers.logger.noColor"] = &structpb.Value{
		Kind: &structpb.Value_BoolValue{
			BoolValue: false,
		},
	}
	return req, nil
}
