// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.17.3
// source: yandex/cloud/datasphere/v2/project_service.proto

package datasphere

import (
	context "context"
	access "github.com/yandex-cloud/go-genproto/yandex/cloud/access"
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ProjectService_Create_FullMethodName               = "/yandex.cloud.datasphere.v2.ProjectService/Create"
	ProjectService_Update_FullMethodName               = "/yandex.cloud.datasphere.v2.ProjectService/Update"
	ProjectService_Delete_FullMethodName               = "/yandex.cloud.datasphere.v2.ProjectService/Delete"
	ProjectService_Open_FullMethodName                 = "/yandex.cloud.datasphere.v2.ProjectService/Open"
	ProjectService_Get_FullMethodName                  = "/yandex.cloud.datasphere.v2.ProjectService/Get"
	ProjectService_List_FullMethodName                 = "/yandex.cloud.datasphere.v2.ProjectService/List"
	ProjectService_GetUnitBalance_FullMethodName       = "/yandex.cloud.datasphere.v2.ProjectService/GetUnitBalance"
	ProjectService_SetUnitBalance_FullMethodName       = "/yandex.cloud.datasphere.v2.ProjectService/SetUnitBalance"
	ProjectService_Execute_FullMethodName              = "/yandex.cloud.datasphere.v2.ProjectService/Execute"
	ProjectService_GetCellOutputs_FullMethodName       = "/yandex.cloud.datasphere.v2.ProjectService/GetCellOutputs"
	ProjectService_GetStateVariables_FullMethodName    = "/yandex.cloud.datasphere.v2.ProjectService/GetStateVariables"
	ProjectService_ListAccessBindings_FullMethodName   = "/yandex.cloud.datasphere.v2.ProjectService/ListAccessBindings"
	ProjectService_SetAccessBindings_FullMethodName    = "/yandex.cloud.datasphere.v2.ProjectService/SetAccessBindings"
	ProjectService_UpdateAccessBindings_FullMethodName = "/yandex.cloud.datasphere.v2.ProjectService/UpdateAccessBindings"
)

// ProjectServiceClient is the client API for ProjectService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProjectServiceClient interface {
	// Creates a project in the specified folder.
	Create(ctx context.Context, in *CreateProjectRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates the specified project.
	Update(ctx context.Context, in *UpdateProjectRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Deletes the specified project.
	Delete(ctx context.Context, in *DeleteProjectRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Opens the specified project.
	Open(ctx context.Context, in *OpenProjectRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Returns the specified project.
	Get(ctx context.Context, in *GetProjectRequest, opts ...grpc.CallOption) (*Project, error)
	// Lists projects for the specified community.
	List(ctx context.Context, in *ListProjectsRequest, opts ...grpc.CallOption) (*ListProjectsResponse, error)
	// Returns the unit balance of the specified project.
	GetUnitBalance(ctx context.Context, in *GetUnitBalanceRequest, opts ...grpc.CallOption) (*GetUnitBalanceResponse, error)
	// Sets the unit balance of the specified project.
	SetUnitBalance(ctx context.Context, in *SetUnitBalanceRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Executes code in the specified cell or notebook.
	Execute(ctx context.Context, in *ProjectExecutionRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Returns outputs of the specified cell.
	GetCellOutputs(ctx context.Context, in *CellOutputsRequest, opts ...grpc.CallOption) (*CellOutputsResponse, error)
	// Returns state variables of the specified notebook.
	GetStateVariables(ctx context.Context, in *GetStateVariablesRequest, opts ...grpc.CallOption) (*GetStateVariablesResponse, error)
	// Lists access bindings for the project.
	ListAccessBindings(ctx context.Context, in *access.ListAccessBindingsRequest, opts ...grpc.CallOption) (*access.ListAccessBindingsResponse, error)
	// Sets access bindings for the project.
	SetAccessBindings(ctx context.Context, in *access.SetAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates access bindings for the project.
	UpdateAccessBindings(ctx context.Context, in *access.UpdateAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error)
}

type projectServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProjectServiceClient(cc grpc.ClientConnInterface) ProjectServiceClient {
	return &projectServiceClient{cc}
}

func (c *projectServiceClient) Create(ctx context.Context, in *CreateProjectRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, ProjectService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) Update(ctx context.Context, in *UpdateProjectRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, ProjectService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) Delete(ctx context.Context, in *DeleteProjectRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, ProjectService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) Open(ctx context.Context, in *OpenProjectRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, ProjectService_Open_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) Get(ctx context.Context, in *GetProjectRequest, opts ...grpc.CallOption) (*Project, error) {
	out := new(Project)
	err := c.cc.Invoke(ctx, ProjectService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) List(ctx context.Context, in *ListProjectsRequest, opts ...grpc.CallOption) (*ListProjectsResponse, error) {
	out := new(ListProjectsResponse)
	err := c.cc.Invoke(ctx, ProjectService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) GetUnitBalance(ctx context.Context, in *GetUnitBalanceRequest, opts ...grpc.CallOption) (*GetUnitBalanceResponse, error) {
	out := new(GetUnitBalanceResponse)
	err := c.cc.Invoke(ctx, ProjectService_GetUnitBalance_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) SetUnitBalance(ctx context.Context, in *SetUnitBalanceRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, ProjectService_SetUnitBalance_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) Execute(ctx context.Context, in *ProjectExecutionRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, ProjectService_Execute_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) GetCellOutputs(ctx context.Context, in *CellOutputsRequest, opts ...grpc.CallOption) (*CellOutputsResponse, error) {
	out := new(CellOutputsResponse)
	err := c.cc.Invoke(ctx, ProjectService_GetCellOutputs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) GetStateVariables(ctx context.Context, in *GetStateVariablesRequest, opts ...grpc.CallOption) (*GetStateVariablesResponse, error) {
	out := new(GetStateVariablesResponse)
	err := c.cc.Invoke(ctx, ProjectService_GetStateVariables_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) ListAccessBindings(ctx context.Context, in *access.ListAccessBindingsRequest, opts ...grpc.CallOption) (*access.ListAccessBindingsResponse, error) {
	out := new(access.ListAccessBindingsResponse)
	err := c.cc.Invoke(ctx, ProjectService_ListAccessBindings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) SetAccessBindings(ctx context.Context, in *access.SetAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, ProjectService_SetAccessBindings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) UpdateAccessBindings(ctx context.Context, in *access.UpdateAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, ProjectService_UpdateAccessBindings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProjectServiceServer is the server API for ProjectService service.
// All implementations should embed UnimplementedProjectServiceServer
// for forward compatibility
type ProjectServiceServer interface {
	// Creates a project in the specified folder.
	Create(context.Context, *CreateProjectRequest) (*operation.Operation, error)
	// Updates the specified project.
	Update(context.Context, *UpdateProjectRequest) (*operation.Operation, error)
	// Deletes the specified project.
	Delete(context.Context, *DeleteProjectRequest) (*operation.Operation, error)
	// Opens the specified project.
	Open(context.Context, *OpenProjectRequest) (*operation.Operation, error)
	// Returns the specified project.
	Get(context.Context, *GetProjectRequest) (*Project, error)
	// Lists projects for the specified community.
	List(context.Context, *ListProjectsRequest) (*ListProjectsResponse, error)
	// Returns the unit balance of the specified project.
	GetUnitBalance(context.Context, *GetUnitBalanceRequest) (*GetUnitBalanceResponse, error)
	// Sets the unit balance of the specified project.
	SetUnitBalance(context.Context, *SetUnitBalanceRequest) (*operation.Operation, error)
	// Executes code in the specified cell or notebook.
	Execute(context.Context, *ProjectExecutionRequest) (*operation.Operation, error)
	// Returns outputs of the specified cell.
	GetCellOutputs(context.Context, *CellOutputsRequest) (*CellOutputsResponse, error)
	// Returns state variables of the specified notebook.
	GetStateVariables(context.Context, *GetStateVariablesRequest) (*GetStateVariablesResponse, error)
	// Lists access bindings for the project.
	ListAccessBindings(context.Context, *access.ListAccessBindingsRequest) (*access.ListAccessBindingsResponse, error)
	// Sets access bindings for the project.
	SetAccessBindings(context.Context, *access.SetAccessBindingsRequest) (*operation.Operation, error)
	// Updates access bindings for the project.
	UpdateAccessBindings(context.Context, *access.UpdateAccessBindingsRequest) (*operation.Operation, error)
}

// UnimplementedProjectServiceServer should be embedded to have forward compatible implementations.
type UnimplementedProjectServiceServer struct {
}

func (UnimplementedProjectServiceServer) Create(context.Context, *CreateProjectRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedProjectServiceServer) Update(context.Context, *UpdateProjectRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedProjectServiceServer) Delete(context.Context, *DeleteProjectRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedProjectServiceServer) Open(context.Context, *OpenProjectRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Open not implemented")
}
func (UnimplementedProjectServiceServer) Get(context.Context, *GetProjectRequest) (*Project, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedProjectServiceServer) List(context.Context, *ListProjectsRequest) (*ListProjectsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedProjectServiceServer) GetUnitBalance(context.Context, *GetUnitBalanceRequest) (*GetUnitBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUnitBalance not implemented")
}
func (UnimplementedProjectServiceServer) SetUnitBalance(context.Context, *SetUnitBalanceRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetUnitBalance not implemented")
}
func (UnimplementedProjectServiceServer) Execute(context.Context, *ProjectExecutionRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Execute not implemented")
}
func (UnimplementedProjectServiceServer) GetCellOutputs(context.Context, *CellOutputsRequest) (*CellOutputsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCellOutputs not implemented")
}
func (UnimplementedProjectServiceServer) GetStateVariables(context.Context, *GetStateVariablesRequest) (*GetStateVariablesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStateVariables not implemented")
}
func (UnimplementedProjectServiceServer) ListAccessBindings(context.Context, *access.ListAccessBindingsRequest) (*access.ListAccessBindingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAccessBindings not implemented")
}
func (UnimplementedProjectServiceServer) SetAccessBindings(context.Context, *access.SetAccessBindingsRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAccessBindings not implemented")
}
func (UnimplementedProjectServiceServer) UpdateAccessBindings(context.Context, *access.UpdateAccessBindingsRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccessBindings not implemented")
}

// UnsafeProjectServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProjectServiceServer will
// result in compilation errors.
type UnsafeProjectServiceServer interface {
	mustEmbedUnimplementedProjectServiceServer()
}

func RegisterProjectServiceServer(s grpc.ServiceRegistrar, srv ProjectServiceServer) {
	s.RegisterService(&ProjectService_ServiceDesc, srv)
}

func _ProjectService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).Create(ctx, req.(*CreateProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).Update(ctx, req.(*UpdateProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).Delete(ctx, req.(*DeleteProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_Open_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OpenProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).Open(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_Open_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).Open(ctx, req.(*OpenProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).Get(ctx, req.(*GetProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProjectsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).List(ctx, req.(*ListProjectsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_GetUnitBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUnitBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).GetUnitBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_GetUnitBalance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).GetUnitBalance(ctx, req.(*GetUnitBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_SetUnitBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetUnitBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).SetUnitBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_SetUnitBalance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).SetUnitBalance(ctx, req.(*SetUnitBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectExecutionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_Execute_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).Execute(ctx, req.(*ProjectExecutionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_GetCellOutputs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CellOutputsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).GetCellOutputs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_GetCellOutputs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).GetCellOutputs(ctx, req.(*CellOutputsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_GetStateVariables_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStateVariablesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).GetStateVariables(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_GetStateVariables_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).GetStateVariables(ctx, req.(*GetStateVariablesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_ListAccessBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(access.ListAccessBindingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).ListAccessBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_ListAccessBindings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).ListAccessBindings(ctx, req.(*access.ListAccessBindingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_SetAccessBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(access.SetAccessBindingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).SetAccessBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_SetAccessBindings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).SetAccessBindings(ctx, req.(*access.SetAccessBindingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_UpdateAccessBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(access.UpdateAccessBindingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).UpdateAccessBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_UpdateAccessBindings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).UpdateAccessBindings(ctx, req.(*access.UpdateAccessBindingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProjectService_ServiceDesc is the grpc.ServiceDesc for ProjectService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProjectService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.datasphere.v2.ProjectService",
	HandlerType: (*ProjectServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ProjectService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ProjectService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ProjectService_Delete_Handler,
		},
		{
			MethodName: "Open",
			Handler:    _ProjectService_Open_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ProjectService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ProjectService_List_Handler,
		},
		{
			MethodName: "GetUnitBalance",
			Handler:    _ProjectService_GetUnitBalance_Handler,
		},
		{
			MethodName: "SetUnitBalance",
			Handler:    _ProjectService_SetUnitBalance_Handler,
		},
		{
			MethodName: "Execute",
			Handler:    _ProjectService_Execute_Handler,
		},
		{
			MethodName: "GetCellOutputs",
			Handler:    _ProjectService_GetCellOutputs_Handler,
		},
		{
			MethodName: "GetStateVariables",
			Handler:    _ProjectService_GetStateVariables_Handler,
		},
		{
			MethodName: "ListAccessBindings",
			Handler:    _ProjectService_ListAccessBindings_Handler,
		},
		{
			MethodName: "SetAccessBindings",
			Handler:    _ProjectService_SetAccessBindings_Handler,
		},
		{
			MethodName: "UpdateAccessBindings",
			Handler:    _ProjectService_UpdateAccessBindings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/datasphere/v2/project_service.proto",
}
