// Code generated by sdkgen. DO NOT EDIT.

// nolint
package text_generation

import (
	"context"

	"google.golang.org/grpc"

	text_generation "github.com/yandex-cloud/go-genproto/yandex/cloud/ai/foundation_models/v1/text_generation"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
)

//revive:disable

// TextGenerationAsyncServiceClient is a text_generation.TextGenerationAsyncServiceClient with
// lazy GRPC connection initialization.
type TextGenerationAsyncServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// Completion implements text_generation.TextGenerationAsyncServiceClient
func (c *TextGenerationAsyncServiceClient) Completion(ctx context.Context, in *text_generation.CompletionRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return text_generation.NewTextGenerationAsyncServiceClient(conn).Completion(ctx, in, opts...)
}
