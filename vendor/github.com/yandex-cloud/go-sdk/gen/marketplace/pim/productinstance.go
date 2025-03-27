// Code generated by sdkgen. DO NOT EDIT.

// nolint
package saas

import (
	"context"

	"google.golang.org/grpc"

	saas "github.com/yandex-cloud/go-genproto/yandex/cloud/marketplace/pim/v1/saas"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
)

//revive:disable

// ProductInstanceServiceClient is a saas.ProductInstanceServiceClient with
// lazy GRPC connection initialization.
type ProductInstanceServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// Claim implements saas.ProductInstanceServiceClient
func (c *ProductInstanceServiceClient) Claim(ctx context.Context, in *saas.ClaimProductInstanceRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return saas.NewProductInstanceServiceClient(conn).Claim(ctx, in, opts...)
}

// Get implements saas.ProductInstanceServiceClient
func (c *ProductInstanceServiceClient) Get(ctx context.Context, in *saas.GetProductInstanceRequest, opts ...grpc.CallOption) (*saas.ProductInstance, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return saas.NewProductInstanceServiceClient(conn).Get(ctx, in, opts...)
}
