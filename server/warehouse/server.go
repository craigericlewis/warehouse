package warehouse

import (
	"context"
	"errors"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/craigericlewis/warehouse/server/api"
	"github.com/craigericlewis/warehouse/server/builder"
)

type grpcServer struct {
	service             Service
	pythonBuilderClient *builder.Client
}

func ListenGRPC(s Service, pythonBuilderUrl string, port int) error {
	pythonBuilderClient, err := builder.NewClient(pythonBuilderUrl)

	if err != nil {
		return err
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		pythonBuilderClient.Close()
		return err
	}

	serv := grpc.NewServer()
	api.RegisterPackageServiceServer(serv, &grpcServer{s, pythonBuilderClient})
	reflection.Register(serv)
	return serv.Serve(lis)
}

func (s *grpcServer) GetPackageInfo(ctx context.Context, req *api.GetPackageInfoRequest) (*api.PackageInfo, error) {
	if req.GetPackageName() == "" {
		return nil, errors.New("Invalid or empty package name")
	}

	if req.GetLanguage() == api.Language_NONE {
		return nil, errors.New("Empty language specification")
	}

	var packageSize *api.PackageSize
	var err error

	if req.GetLanguage() == api.Language_PYTHON2 {
		packageSize, err = s.getPythonPackageInfo(ctx, req.GetPackageName())
	}

	if err != nil {
		return nil, err
	}

	resp, err := s.service.GetPackageInfo(ctx, req.GetPackageName(), packageSize.Size, req.GetLanguage())

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *grpcServer) getPythonPackageInfo(ctx context.Context, packageName string) (*api.PackageSize, error) {
	packageSize, err := s.pythonBuilderClient.GetPackageSize(ctx, packageName)

	if err != nil {
		return nil, err
	}

	return packageSize, nil
}
