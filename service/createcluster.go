package service

import (
	"context"

	"grpc-pixiu/cmd"
)

var CreateClusterService = &createClusterService{}

type createClusterService struct {
}

func (p *createClusterService) CreateCluster(ctx context.Context, request *ClusterRequest) (*ClusterResponse, error) {
	clusterName := request.ClusterName

	if err := cmd.CheckKubez(request); err != nil {
		return &ClusterResponse{ClusterName: clusterName}, err
	}

	return &ClusterResponse{ClusterName: clusterName}, nil
}
