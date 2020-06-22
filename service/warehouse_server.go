package service

import (
	"context"

	api "github.com/craigericlewis/warehouse/api"
)

type WarehouseServer struct {
}

func NewWarehouseServer() *WarehouseServer {
	return &WarehouseServer{}
}

func (server *WarehouseServer) GetPackageInfoRequest(
	ctx context.Context,
	req *api.GetPackageInfoRequest,
) (*api.GetPackageInfoResponse, error) {
	packageName := req.PackageName()
}
