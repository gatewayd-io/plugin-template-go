package plugin

import (
	"context"
	"encoding/base64"

	v1 "github.com/gatewayd-io/gatewayd-plugin-sdk/plugin/v1"
	"github.com/hashicorp/go-hclog"
	goplugin "github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/structpb"
)

type Plugin struct {
	goplugin.GRPCPlugin
	v1.GatewayDPluginServiceServer
	Logger hclog.Logger
}

type TestPlugin struct {
	goplugin.NetRPCUnsupportedPlugin
	Impl Plugin
}

// GRPCServer registers the plugin with the gRPC server.
func (p *TestPlugin) GRPCServer(b *goplugin.GRPCBroker, s *grpc.Server) error {
	v1.RegisterGatewayDPluginServiceServer(s, &p.Impl)
	return nil
}

// GRPCClient returns the plugin client.
func (p *TestPlugin) GRPCClient(ctx context.Context, b *goplugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return v1.NewGatewayDPluginServiceClient(c), nil
}

// NewTestPlugin returns a new instance of the TestPlugin.
func NewTestPlugin(impl Plugin) *TestPlugin {
	return &TestPlugin{
		NetRPCUnsupportedPlugin: goplugin.NetRPCUnsupportedPlugin{},
		Impl:                    impl,
	}
}

// GetPluginConfig returns the plugin config.
func (p *Plugin) GetPluginConfig(
	ctx context.Context, req *structpb.Struct) (*structpb.Struct, error) {
	return structpb.NewStruct(PluginConfig)
}

// OnConfigLoaded is called when the global config is loaded by GatewayD.
func (p *Plugin) OnConfigLoaded(
	ctx context.Context, req *structpb.Struct) (*structpb.Struct, error) {
	OnConfigLoaded.Inc()

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
	OnTrafficFromClient.Inc()

	request := req.Fields["request"].GetStringValue()
	if reqBytes, err := base64.StdEncoding.DecodeString(request); err == nil {
		p.Logger.Debug("OnIngressTraffic", "request", string(reqBytes))
	}

	return req, nil
}

// OnTrafficFromServer is called when a response is received by GatewayD from the server.
func (p *Plugin) OnTrafficFromServer(
	ctx context.Context, resp *structpb.Struct) (*structpb.Struct, error) {
	OnTrafficFromServer.Inc()

	response := resp.Fields["response"].GetStringValue()
	if respBytes, err := base64.StdEncoding.DecodeString(response); err == nil {
		p.Logger.Debug("OnEgressTraffic", "response", string(respBytes))
	}
	return resp, nil
}
