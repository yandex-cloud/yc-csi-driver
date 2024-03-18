// Code generated by sdkgen. DO NOT EDIT.

// nolint
package datasphere

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/access"
	datasphere "github.com/yandex-cloud/go-genproto/yandex/cloud/datasphere/v2"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
)

//revive:disable

// ProjectServiceClient is a datasphere.ProjectServiceClient with
// lazy GRPC connection initialization.
type ProjectServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// Create implements datasphere.ProjectServiceClient
func (c *ProjectServiceClient) Create(ctx context.Context, in *datasphere.CreateProjectRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return datasphere.NewProjectServiceClient(conn).Create(ctx, in, opts...)
}

// Delete implements datasphere.ProjectServiceClient
func (c *ProjectServiceClient) Delete(ctx context.Context, in *datasphere.DeleteProjectRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return datasphere.NewProjectServiceClient(conn).Delete(ctx, in, opts...)
}

// Execute implements datasphere.ProjectServiceClient
func (c *ProjectServiceClient) Execute(ctx context.Context, in *datasphere.ProjectExecutionRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return datasphere.NewProjectServiceClient(conn).Execute(ctx, in, opts...)
}

// Get implements datasphere.ProjectServiceClient
func (c *ProjectServiceClient) Get(ctx context.Context, in *datasphere.GetProjectRequest, opts ...grpc.CallOption) (*datasphere.Project, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return datasphere.NewProjectServiceClient(conn).Get(ctx, in, opts...)
}

// GetCellOutputs implements datasphere.ProjectServiceClient
func (c *ProjectServiceClient) GetCellOutputs(ctx context.Context, in *datasphere.CellOutputsRequest, opts ...grpc.CallOption) (*datasphere.CellOutputsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return datasphere.NewProjectServiceClient(conn).GetCellOutputs(ctx, in, opts...)
}

// GetStateVariables implements datasphere.ProjectServiceClient
func (c *ProjectServiceClient) GetStateVariables(ctx context.Context, in *datasphere.GetStateVariablesRequest, opts ...grpc.CallOption) (*datasphere.GetStateVariablesResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return datasphere.NewProjectServiceClient(conn).GetStateVariables(ctx, in, opts...)
}

// GetUnitBalance implements datasphere.ProjectServiceClient
func (c *ProjectServiceClient) GetUnitBalance(ctx context.Context, in *datasphere.GetUnitBalanceRequest, opts ...grpc.CallOption) (*datasphere.GetUnitBalanceResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return datasphere.NewProjectServiceClient(conn).GetUnitBalance(ctx, in, opts...)
}

// List implements datasphere.ProjectServiceClient
func (c *ProjectServiceClient) List(ctx context.Context, in *datasphere.ListProjectsRequest, opts ...grpc.CallOption) (*datasphere.ListProjectsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return datasphere.NewProjectServiceClient(conn).List(ctx, in, opts...)
}

type ProjectIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *ProjectServiceClient
	request *datasphere.ListProjectsRequest

	items []*datasphere.Project
}

func (c *ProjectServiceClient) ProjectIterator(ctx context.Context, req *datasphere.ListProjectsRequest, opts ...grpc.CallOption) *ProjectIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &ProjectIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *ProjectIterator) Next() bool {
	if it.err != nil {
		return false
	}
	if len(it.items) > 1 {
		it.items[0] = nil
		it.items = it.items[1:]
		return true
	}
	it.items = nil // consume last item, if any

	if it.started && it.request.PageToken == "" {
		return false
	}
	it.started = true

	if it.requestedSize == 0 || it.requestedSize > it.pageSize {
		it.request.PageSize = it.pageSize
	} else {
		it.request.PageSize = it.requestedSize
	}

	response, err := it.client.List(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.Projects
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *ProjectIterator) Take(size int64) ([]*datasphere.Project, error) {
	if it.err != nil {
		return nil, it.err
	}

	if size == 0 {
		size = 1 << 32 // something insanely large
	}
	it.requestedSize = size
	defer func() {
		// reset iterator for future calls.
		it.requestedSize = 0
	}()

	var result []*datasphere.Project

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *ProjectIterator) TakeAll() ([]*datasphere.Project, error) {
	return it.Take(0)
}

func (it *ProjectIterator) Value() *datasphere.Project {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *ProjectIterator) Error() error {
	return it.err
}

// ListAccessBindings implements datasphere.ProjectServiceClient
func (c *ProjectServiceClient) ListAccessBindings(ctx context.Context, in *access.ListAccessBindingsRequest, opts ...grpc.CallOption) (*access.ListAccessBindingsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return datasphere.NewProjectServiceClient(conn).ListAccessBindings(ctx, in, opts...)
}

type ProjectAccessBindingsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *ProjectServiceClient
	request *access.ListAccessBindingsRequest

	items []*access.AccessBinding
}

func (c *ProjectServiceClient) ProjectAccessBindingsIterator(ctx context.Context, req *access.ListAccessBindingsRequest, opts ...grpc.CallOption) *ProjectAccessBindingsIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &ProjectAccessBindingsIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *ProjectAccessBindingsIterator) Next() bool {
	if it.err != nil {
		return false
	}
	if len(it.items) > 1 {
		it.items[0] = nil
		it.items = it.items[1:]
		return true
	}
	it.items = nil // consume last item, if any

	if it.started && it.request.PageToken == "" {
		return false
	}
	it.started = true

	if it.requestedSize == 0 || it.requestedSize > it.pageSize {
		it.request.PageSize = it.pageSize
	} else {
		it.request.PageSize = it.requestedSize
	}

	response, err := it.client.ListAccessBindings(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.AccessBindings
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *ProjectAccessBindingsIterator) Take(size int64) ([]*access.AccessBinding, error) {
	if it.err != nil {
		return nil, it.err
	}

	if size == 0 {
		size = 1 << 32 // something insanely large
	}
	it.requestedSize = size
	defer func() {
		// reset iterator for future calls.
		it.requestedSize = 0
	}()

	var result []*access.AccessBinding

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *ProjectAccessBindingsIterator) TakeAll() ([]*access.AccessBinding, error) {
	return it.Take(0)
}

func (it *ProjectAccessBindingsIterator) Value() *access.AccessBinding {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *ProjectAccessBindingsIterator) Error() error {
	return it.err
}

// Open implements datasphere.ProjectServiceClient
func (c *ProjectServiceClient) Open(ctx context.Context, in *datasphere.OpenProjectRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return datasphere.NewProjectServiceClient(conn).Open(ctx, in, opts...)
}

// SetAccessBindings implements datasphere.ProjectServiceClient
func (c *ProjectServiceClient) SetAccessBindings(ctx context.Context, in *access.SetAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return datasphere.NewProjectServiceClient(conn).SetAccessBindings(ctx, in, opts...)
}

// SetUnitBalance implements datasphere.ProjectServiceClient
func (c *ProjectServiceClient) SetUnitBalance(ctx context.Context, in *datasphere.SetUnitBalanceRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return datasphere.NewProjectServiceClient(conn).SetUnitBalance(ctx, in, opts...)
}

// Update implements datasphere.ProjectServiceClient
func (c *ProjectServiceClient) Update(ctx context.Context, in *datasphere.UpdateProjectRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return datasphere.NewProjectServiceClient(conn).Update(ctx, in, opts...)
}

// UpdateAccessBindings implements datasphere.ProjectServiceClient
func (c *ProjectServiceClient) UpdateAccessBindings(ctx context.Context, in *access.UpdateAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return datasphere.NewProjectServiceClient(conn).UpdateAccessBindings(ctx, in, opts...)
}