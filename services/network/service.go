package network

import (
	"context"
	"net"

	"github.com/ehazlett/element"
	"github.com/ehazlett/stellar"
	datastoreapi "github.com/ehazlett/stellar/api/services/datastore/v1"
	api "github.com/ehazlett/stellar/api/services/network/v1"
	"github.com/ehazlett/stellar/client"
	ptypes "github.com/gogo/protobuf/types"
	"google.golang.org/grpc"
)

const (
	serviceID = "stellar.services.network.v1"
	// TODO: make configurable
	// default max subnets (max nodes)
	maxSubnets          = 1024
	dsNetworkBucketName = "stellar." + stellar.APIVersion + ".services.network"
)

var (
	empty = &ptypes.Empty{}
)

type service struct {
	network *net.IPNet
	agent   *element.Agent
	ds      datastoreapi.DatastoreServer
	config  *stellar.Config
}

func New(cfg *stellar.Config, agent *element.Agent, ds datastoreapi.DatastoreServer) (*service, error) {
	return &service{
		network: cfg.Subnet,
		agent:   agent,
		config:  cfg,
		ds:      ds,
	}, nil
}

func (s *service) Register(server *grpc.Server) error {
	api.RegisterNetworkServer(server, s)
	return nil
}

func (s *service) ID() string {
	return serviceID
}

func (s *service) Info(ctx context.Context, req *api.InfoRequest) (*api.InfoResponse, error) {
	return &api.InfoResponse{
		ID: serviceID,
	}, nil
}

func (s *service) Start() error {
	return nil
}

func (s *service) client(address string) (*client.Client, error) {
	opts, err := client.DialOptionsFromConfig(s.config)
	if err != nil {
		return nil, err
	}
	return client.NewClient(address, opts...)
}

func (s *service) Stop() error {
	return nil
}
