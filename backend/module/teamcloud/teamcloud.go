package teamcloud

import (
	"context"
	"errors"

	"github.com/golang/protobuf/ptypes/any"
	"github.com/uber-go/tally/v4"
	"go.uber.org/zap"

	teamcloudv1 "github.com/colbylwilliams/clutch-teamcloud-gateway/backend/api/teamcloud/v1"
	teamcloudservice "github.com/colbylwilliams/clutch-teamcloud-gateway/backend/service/teamcloud"
	"github.com/lyft/clutch/backend/module"
	"github.com/lyft/clutch/backend/service"
)

const Name = "colbylwilliams.module.teamcloud"

func New(*any.Any, *zap.Logger, tally.Scope) (module.Module, error) {
	svc, ok := service.Registry["colbylwilliams.service.teamcloud"]
	if !ok {
		return nil, errors.New("no teamcloud service was registered")
	}

	client, ok := svc.(teamcloudservice.Client)
	if !ok {
		return nil, errors.New("teamcloud service in registry was the wrong type")
	}

	return &mod{client: client}, nil
}

type mod struct {
	client teamcloudservice.Client
}

func (m *mod) Register(r module.Registrar) error {
	teamcloudv1.RegisterTeamCloudAPIServer(r.GRPCServer(), m)
	return r.RegisterJSONGateway(teamcloudv1.RegisterTeamCloudAPIHandler)
}

func (m *mod) GetTeamCloudInfo(ctx context.Context, request *teamcloudv1.GetTeamCloudInfoRequest) (*teamcloudv1.GetTeamCloudInfoResponse, error) {
	// a, err := m.client.GetTeamCloudInfo(ctx, request.Name)
	a, err := m.client.GetTeamCloudInfo(ctx)
	if err != nil {
		return nil, err
	}
	return a, nil
	// return &teamcloudv1.GetAmiiboResponse{Amiibo: a}, nil
}
