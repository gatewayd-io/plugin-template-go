package plugin

import (
	"context"

	"github.com/davecgh/go-spew/spew"
	v1 "github.com/gatewayd-io/gatewayd-plugin-test/plugin/v1"
	goplugin "github.com/hashicorp/go-plugin"
)

type GDPServiceServerImpl struct {
	v1.UnimplementedGDPServiceServer
}

type TestPlugin struct {
	goplugin.Plugin
	Impl GDPServiceServerImpl
}

var _ v1.GDPServiceServer = &GDPServiceServerImpl{}

func (s *GDPServiceServerImpl) PluginConfig(
	ctx context.Context, req *v1.PluginConfigRequest) (*v1.PluginConfigResponse, error) {
	return &v1.PluginConfigResponse{}, nil
}

func (s *GDPServiceServerImpl) OnConfigLoaded(
	ctx context.Context, req *v1.OnConfigLoadedRequest) (*v1.OnConfigLoadedResponse, error) {
	spew.Dump(req.Config)
	return &v1.OnConfigLoadedResponse{}, nil
}

func (s *GDPServiceServerImpl) OnNewLogger(
	ctx context.Context, req *v1.OnNewLoggerRequest) (*v1.OnNewLoggerResponse, error) {
	return &v1.OnNewLoggerResponse{}, nil
}

func (s *GDPServiceServerImpl) OnNewPool(
	ctx context.Context, req *v1.OnNewPoolRequest) (*v1.OnNewPoolResponse, error) {
	return &v1.OnNewPoolResponse{}, nil
}

func (s *GDPServiceServerImpl) OnNewProxy(
	ctx context.Context, req *v1.OnNewProxyRequest) (*v1.OnNewProxyResponse, error) {
	return &v1.OnNewProxyResponse{}, nil
}

func (s *GDPServiceServerImpl) OnNewServer(
	ctx context.Context, req *v1.OnNewServerRequest) (*v1.OnNewServerResponse, error) {
	return &v1.OnNewServerResponse{}, nil
}

func (s *GDPServiceServerImpl) OnSignal(
	ctx context.Context, req *v1.OnSignalRequest) (*v1.OnSignalResponse, error) {
	return &v1.OnSignalResponse{}, nil
}

func (s *GDPServiceServerImpl) OnRun(
	ctx context.Context, req *v1.OnRunRequest) (*v1.OnRunResponse, error) {
	return &v1.OnRunResponse{}, nil
}

func (s *GDPServiceServerImpl) OnBooting(
	ctx context.Context, req *v1.OnBootingRequest) (*v1.OnBootingResponse, error) {
	return &v1.OnBootingResponse{}, nil
}

func (s *GDPServiceServerImpl) OnBooted(
	ctx context.Context, req *v1.OnBootedRequest) (*v1.OnBootedResponse, error) {
	return &v1.OnBootedResponse{}, nil
}

func (s *GDPServiceServerImpl) OnOpening(
	ctx context.Context, req *v1.OnOpeningRequest) (*v1.OnOpeningResponse, error) {
	return &v1.OnOpeningResponse{}, nil
}

func (s *GDPServiceServerImpl) OnOpened(
	ctx context.Context, req *v1.OnOpenedRequest) (*v1.OnOpenedResponse, error) {
	return &v1.OnOpenedResponse{}, nil
}

func (s *GDPServiceServerImpl) OnClosing(
	ctx context.Context, req *v1.OnClosingRequest) (*v1.OnClosingResponse, error) {
	return &v1.OnClosingResponse{}, nil
}

func (s *GDPServiceServerImpl) OnClosed(
	ctx context.Context, req *v1.OnClosedRequest) (*v1.OnClosedResponse, error) {
	return &v1.OnClosedResponse{}, nil
}

func (s *GDPServiceServerImpl) OnTraffic(
	ctx context.Context, req *v1.OnTrafficRequest) (*v1.OnTrafficResponse, error) {
	return &v1.OnTrafficResponse{}, nil
}

func (s *GDPServiceServerImpl) OnIngressTraffic(
	ctx context.Context, req *v1.OnIngressTrafficRequest) (*v1.OnIngressTrafficResponse, error) {
	return &v1.OnIngressTrafficResponse{}, nil
}

func (s *GDPServiceServerImpl) OnEgressTraffic(
	ctx context.Context, req *v1.OnEgressTrafficRequest) (*v1.OnEgressTrafficResponse, error) {
	return &v1.OnEgressTrafficResponse{}, nil
}

func (s *GDPServiceServerImpl) OnShutdown(
	ctx context.Context, req *v1.OnShutdownRequest) (*v1.OnShutdownResponse, error) {
	return &v1.OnShutdownResponse{}, nil
}

func (s *GDPServiceServerImpl) OnTick(
	ctx context.Context, req *v1.OnTickRequest) (*v1.OnTickResponse, error) {
	return &v1.OnTickResponse{}, nil
}

func (s *GDPServiceServerImpl) OnNewClient(
	ctx context.Context, req *v1.OnNewClientRequest) (*v1.OnNewClientResponse, error) {
	return &v1.OnNewClientResponse{}, nil
}
