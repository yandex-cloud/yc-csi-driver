/*
Copyright 2024 YANDEX LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package pubapimock

import (
	"context"
	"reflect"
	"time"

	tiface "github.com/mitchellh/go-testing-interface"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"k8s.io/apimachinery/pkg/util/rand"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/endpoint"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/iam/v1"
	ycsdk "github.com/yandex-cloud/go-sdk"
	"github.com/yandex-cloud/yc-csi-driver/pkg/testutil/grpctest"
	"github.com/yandex-cloud/yc-csi-driver/pkg/testutil/pubapimock/mock/computemocks"
	"github.com/yandex-cloud/yc-csi-driver/pkg/testutil/pubapimock/mock/endpointmocks"
	"github.com/yandex-cloud/yc-csi-driver/pkg/testutil/pubapimock/mock/iammocks"
	"github.com/yandex-cloud/yc-csi-driver/pkg/testutil/pubapimock/mock/resourcemanagermocks"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/resourcemanager/v1"
)

// Remember that API:
// * Should be configurable.
// * Should afford configuration shortcuts for happy paths.
// * Should not be complex.
// * Should not contain configuration specific for tests from single package -
//   in such case keep such configuration in this package.

// API is helper to mock public API on server side.
type API struct {
	server *grpctest.Server
	// Should be set on one of methods.
	credentials     ycsdk.Credentials
	IAM             IAM
	ResourceManager ResourceManager
	// Just add generated mock as a field: it will be initialized and served automagically.
	Compute     Compute
	APIEndpoint APIEndpoint
}

type APIEndpoint struct {
	APIEndpoint *endpointmocks.ApiEndpointServiceServer
}

type Compute struct {
	Instance *computemocks.InstanceServiceServer
	Zone     *computemocks.ZoneServiceServer
	Disk     *computemocks.DiskServiceServer
}

type BaseResources struct {
	CloudID  string
	FolderID string
}

type IAM struct {
	IAMToken       *iammocks.IamTokenServiceServer
	ServiceAccount *iammocks.ServiceAccountServiceServer
}

type ResourceManager struct {
	Cloud  *resourcemanagermocks.CloudServiceServer
	Folder *resourcemanagermocks.FolderServiceServer
}

// New returns mock configured for most happy path tests.
// * Endpoints server returns mock server addr for all ids.
// * Credentials set to some OAuth token, and IAMToken server returns some IAM Token for it.
// * Resource manager configured to serve Get/List requests for one Cloud with one Folder.
func New(t tiface.T) (*API, BaseResources) {
	api := NewRaw(t)

	api.MockIAMToken()
	api.MockEndpoint()

	cloudID := api.ResourceManager.MockCloud()
	folderID := api.ResourceManager.MockFolder(cloudID)

	return api, BaseResources{
		CloudID:  cloudID,
		FolderID: folderID,
	}
}

// NewRaw return API without any mocks configuration.
func NewRaw(t tiface.T) *API {
	api := new(API)
	mocks := initMocks(t, reflect.ValueOf(api).Elem())
	api.server = grpctest.NewServer().Serve(mocks...)
	return api
}

func (api *API) Addr() string {
	return api.server.Addr()
}

func (api *API) Close() {
	api.server.Close()
}

func (api *API) SDK(t tiface.T) *ycsdk.SDK {
	sdk, err := ycsdk.Build(
		context.Background(),
		ycsdk.Config{
			Credentials: api.credentials,
			Endpoint:    api.Addr(),
			Plaintext:   true,
		},
	)
	require.NoError(t, err)
	return sdk
}

func (api *API) MockEndpoint() {
	api.APIEndpoint.APIEndpoint.On("Get", mock.Anything, mock.Anything).Return(
		&endpoint.ApiEndpoint{Id: string(ycsdk.ComputeServiceID), Address: api.Addr()}, nil,
	)

	api.APIEndpoint.APIEndpoint.On("List", mock.Anything, mock.Anything).Return(
		&endpoint.ListApiEndpointsResponse{
			Endpoints: []*endpoint.ApiEndpoint{
				{Id: string(ycsdk.IAMServiceID), Address: api.Addr()},
				{Id: string(ycsdk.OperationServiceID), Address: api.Addr()},
				{Id: string(ycsdk.ResourceManagementServiceID), Address: api.Addr()},
				{Id: string(ycsdk.ApiEndpointServiceID), Address: api.Addr()},
				{Id: string(ycsdk.ComputeServiceID), Address: api.Addr()},
			},
		},
		nil,
	)
}

func (api *API) MockIAMToken() {
	oauthToken := rand.String(20)
	iamToken := rand.String(20)

	request := &iam.CreateIamTokenRequest{
		Identity: &iam.CreateIamTokenRequest_YandexPassportOauthToken{
			YandexPassportOauthToken: oauthToken,
		},
	}

	response := &iam.CreateIamTokenResponse{
		IamToken: iamToken,
		ExpiresAt: &timestamppb.Timestamp{
			Seconds: time.Now().Add(100500 * time.Second).Unix(),
			Nanos:   0,
		}}

	api.IAM.IAMToken.On("Create", mock.Anything, MatchesProto(request)).Return(response, nil)
	api.credentials = ycsdk.OAuthToken(oauthToken)
}

func (rm *ResourceManager) MockCloud() string {
	cloudID := RandID()

	rm.Cloud.On("Get", mock.Anything, resourcemanager.GetCloudRequest{
		CloudId: cloudID,
	}).Return(&resourcemanager.Cloud{
		Id: cloudID,
	}, nil)

	rm.Cloud.On("List", mock.Anything, mock.Anything).Return(&resourcemanager.ListCloudsResponse{
		Clouds: []*resourcemanager.Cloud{{Id: cloudID}},
	})

	return cloudID
}

func (rm *ResourceManager) MockFolder(cloudID string) string {
	folderID := RandID()

	rm.Folder.On("Get", mock.Anything, resourcemanager.GetFolderRequest{
		FolderId: folderID,
	}).Return(&resourcemanager.Folder{
		Id: folderID,
	}, nil)

	rm.Folder.On("List", mock.Anything, MatchesProto(&resourcemanager.ListFoldersRequest{
		CloudId: cloudID,
	})).Return(&resourcemanager.ListFoldersResponse{
		Folders: []*resourcemanager.Folder{{Id: folderID}},
	})

	return folderID
}

func RandID() string {
	return rand.String(20)
}

// initMocks recursively visit struct fields, find grpctest.Registrant fields, set them if they are nil,
// and returns all of them.
func initMocks(t tiface.T, v reflect.Value) []grpctest.Registrant {
	require.Equal(t, reflect.Struct.String(), v.Kind().String())
	var rs []grpctest.Registrant
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)

		if f.Type().Implements(registrantIface) {
			if f.IsNil() {
				f.Set(reflect.New(f.Type().Elem()))
			}

			r := f.Interface().(grpctest.Registrant)
			rs = append(rs, r)

			continue
		}

		f = reflect.Indirect(f)

		if f.Kind() != reflect.Struct {
			continue
		}

		rs = append(rs, initMocks(t, f)...)
	}

	return rs
}

func MatchesProto(expected proto.Message) interface{} {
	return mock.MatchedBy(func(actual proto.Message) bool {
		return proto.Equal(actual, expected)
	})
}

var registrantIface = reflect.TypeOf((*grpctest.Registrant)(nil)).Elem()
