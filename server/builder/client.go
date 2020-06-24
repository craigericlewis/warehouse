//go:generate protoc ./builder.proto --go_out=plugins=grpc:../api
// protoc --python_out=./python builder.proto

package builder

import (
	"context"

	"google.golang.org/grpc"

	api "github.com/craigericlewis/warehouse/server/api"
)

type Client struct {
	conn    *grpc.ClientConn
	service api.BuilderServiceClient
}

func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	c := api.NewBuilderServiceClient(conn)
	return &Client{conn, c}, nil
}

func (c *Client) Close() {
	c.conn.Close()
}

func (c *Client) GetPackageSize(ctx context.Context, packageName string) (*api.PackageSize, error) {
	resp, err := c.service.GetPackageSize(
		ctx,
		&api.GetPackageSizeRequest{
			PackageName: packageName,
		},
	)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
