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

package public

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1"
	ycsdk "github.com/yandex-cloud/go-sdk"
	"github.com/yandex-cloud/yc-csi-driver/pkg/diskapi"
	"github.com/yandex-cloud/yc-csi-driver/pkg/inflight"
	"github.com/yandex-cloud/yc-csi-driver/pkg/services"
	"github.com/yandex-cloud/yc-csi-driver/pkg/services/controller"
	"github.com/yandex-cloud/yc-csi-driver/pkg/testutil/pubapimock"
)

const (
	zoneID               = "ru-central1-a"
	diskName             = "NewVolume"
	diskTypeID           = "mytype"
	instanceID           = "NewInstance"
	defaultDiskSizeBytes = 4 * 1024 * 1024 * 1024
)

type testResources struct {
	t                    *testing.T
	folderID             string
	networkFolderID      string
	clusterFolderID      string
	networkID            string
	zoneID               string
	diskTypeID           string
	diskName             string
	testData             *testData
	defaultDiskSizeBytes int64
}

type controllerTest struct {
	diskapi    diskapi.DiskAPI
	sdk        *ycsdk.SDK
	controller csi.ControllerServer
	cloudRes   *testResources
	cleanup    func()
}

func TestCreateDisk(t *testing.T) {
	test := initializeTest(t)
	defer test.cleanup()
	disk, err := test.diskapi.CreateDisk(
		context.Background(),
		&diskapi.CreateDiskRequest{
			Name:   test.cloudRes.diskName,
			ZoneID: test.cloudRes.zoneID,
			TypeID: test.cloudRes.diskTypeID,
			Size:   test.cloudRes.defaultDiskSizeBytes,
		})
	require.NoError(t, err)
	assert.Equal(t, disk.Name, test.cloudRes.diskName)
}

func TestExpandDisk(t *testing.T) {
	test := initializeTest(t)
	defer test.cleanup()
	disk, err := test.diskapi.CreateDisk(
		context.Background(),
		&diskapi.CreateDiskRequest{
			Name:   test.cloudRes.diskName,
			ZoneID: test.cloudRes.zoneID,
			TypeID: test.cloudRes.diskTypeID,
			Size:   test.cloudRes.defaultDiskSizeBytes,
		})
	require.NoError(t, err)
	newSize := disk.Size + 4*1024*1024*1024
	err = test.diskapi.ExpandDisk(
		context.Background(),
		&diskapi.ExpandDiskRequest{
			ID:   disk.ID,
			Size: newSize,
		},
	)
	require.NoError(t, err)
	d, err := test.diskapi.GetDisk(context.Background(), &diskapi.GetDiskRequest{ID: disk.ID})
	require.NoError(t, err)
	require.Equal(t, newSize, d.Size)
}

func TestDeleteDisk(t *testing.T) {
	test := initializeTest(t)
	defer test.cleanup()
	// Delete nonexistent disk shoud return no error.
	err := test.diskapi.DeleteDisk(
		context.Background(),
		&diskapi.DeleteDiskRequest{
			DiskID: "nonexistent",
		},
	)
	require.NoError(t, err)

	disk, err := test.diskapi.CreateDisk(
		context.Background(),
		&diskapi.CreateDiskRequest{
			Name:   test.cloudRes.diskName,
			ZoneID: test.cloudRes.zoneID,
			TypeID: test.cloudRes.diskTypeID,
			Size:   test.cloudRes.defaultDiskSizeBytes,
		})
	require.NoError(t, err)

	err = test.diskapi.DeleteDisk(
		context.Background(),
		&diskapi.DeleteDiskRequest{
			DiskID: disk.ID,
		},
	)
	require.NoError(t, err)

	_, err = test.diskapi.GetDisk(
		context.Background(),
		&diskapi.GetDiskRequest{
			ID: disk.ID,
		},
	)
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))
}

func TestGetDisk(t *testing.T) {
	test := initializeTest(t)
	defer test.cleanup()

	_, err := test.diskapi.GetDisk(
		context.Background(),
		&diskapi.GetDiskRequest{
			ID: "nonexistent",
		},
	)
	require.Error(t, err)
	assert.Error(t, status.Error(codes.NotFound, "Disk not found"), err)

	disk, err := test.diskapi.CreateDisk(
		context.Background(),
		&diskapi.CreateDiskRequest{
			Name:   test.cloudRes.diskName,
			ZoneID: test.cloudRes.zoneID,
			TypeID: test.cloudRes.diskTypeID,
			Size:   test.cloudRes.defaultDiskSizeBytes,
		})
	require.NoError(t, err)

	d, err := test.diskapi.GetDisk(
		context.Background(),
		&diskapi.GetDiskRequest{
			ID: disk.ID,
		},
	)
	require.NoError(t, err)
	assert.Equal(t, disk.ID, d.ID)
}

func TestGetDiskByName(t *testing.T) {
	test := initializeTest(t)
	defer test.cleanup()

	_, err := test.diskapi.GetDiskByName(
		context.Background(),
		&diskapi.GetDiskByNameRequest{
			Name: "nonexistent",
		},
	)
	require.Error(t, err)
	assert.Equal(t, codes.NotFound, status.Code(err))

	disk, err := test.diskapi.CreateDisk(
		context.Background(),
		&diskapi.CreateDiskRequest{
			Name:   test.cloudRes.diskName,
			ZoneID: test.cloudRes.zoneID,
			TypeID: test.cloudRes.diskTypeID,
			Size:   test.cloudRes.defaultDiskSizeBytes,
		})

	require.NoError(t, err)

	d, err := test.diskapi.GetDiskByName(
		context.Background(),
		&diskapi.GetDiskByNameRequest{
			Name: disk.Name,
		},
	)
	require.NoError(t, err)
	assert.Equal(t, disk.Name, d.Name)
}

func TestListDisks(t *testing.T) {
	test := initializeTest(t)
	defer test.cleanup()
	disk, err := test.diskapi.CreateDisk(
		context.Background(),
		&diskapi.CreateDiskRequest{
			Name:   test.cloudRes.diskName,
			ZoneID: test.cloudRes.zoneID,
			TypeID: test.cloudRes.diskTypeID,
			Size:   test.cloudRes.defaultDiskSizeBytes,
		})
	require.NoError(t, err)

	d, err := test.diskapi.ListDisks(
		context.Background(),
		&diskapi.ListDisksRequest{
			Filter: fmt.Sprintf("name = \"%s\"", disk.Name),
		},
	)
	require.NoError(t, err)
	assert.True(t, len(d) > 0)
	assert.Equal(t, disk.ID, d[0].ID)
}

func TestGetInstance(t *testing.T) {
	test := initializeTest(t)
	defer test.cleanup()
	instance, err := test.diskapi.GetInstance(
		context.Background(),
		&diskapi.GetInstanceRequest{
			InstanceID: instanceID,
		},
	)
	require.NoError(t, err)
	assert.Equal(t, instanceID, instance.ID)
}

func TestAttachDetachDisk(t *testing.T) {
	test := initializeTest(t)
	defer test.cleanup()
	instanceID := instanceID
	disk := computeDisk(test.cloudRes.folderID)
	test.cloudRes.testData.disk = append(test.cloudRes.testData.disk, disk)
	err := test.diskapi.AttachDisk(
		context.Background(),
		&diskapi.AttachDiskRequest{
			DiskID:     disk.Id,
			InstanceID: instanceID,
		},
	)
	require.NoError(t, err)
	diskID := disk.Id
	err = test.diskapi.AttachDisk(
		context.Background(),
		&diskapi.AttachDiskRequest{
			DiskID:     diskID,
			InstanceID: instanceID,
		},
	)
	require.NoError(t, err)
	diskAPIInstance, err := test.diskapi.GetInstance(
		context.Background(),
		&diskapi.GetInstanceRequest{
			InstanceID: instanceID,
		},
	)
	require.NoError(t, err)
	assert.Contains(t, diskAPIInstance.DiskAttachments, diskapi.DiskAttachment{DiskID: diskID})

	err = test.diskapi.DetachDisk(
		context.Background(),
		&diskapi.DetachDiskRequest{
			InstanceID: instanceID,
			DiskID:     diskID,
		},
	)
	require.NoError(t, err)
	diskAPIInstance, err = test.diskapi.GetInstance(
		context.Background(),
		&diskapi.GetInstanceRequest{
			InstanceID: instanceID,
		},
	)
	require.NoError(t, err)
	assert.NotContains(t, diskAPIInstance.DiskAttachments, diskapi.DiskAttachment{DiskID: diskID})

	// Detach non-attached disk should return no error.
	err = test.diskapi.DetachDisk(
		context.Background(),
		&diskapi.DetachDiskRequest{
			InstanceID: instanceID,
			DiskID:     diskID,
		},
	)
	require.NoError(t, err)

	// Detach non-existent disk should return no error.
	err = test.diskapi.DetachDisk(
		context.Background(),
		&diskapi.DetachDiskRequest{
			InstanceID: instanceID,
			DiskID:     `nonexistent`,
		},
	)
	require.NoError(t, err)
}

func initializeTest(t *testing.T) *controllerTest {
	api, baseResources := pubapimock.New(t)

	test := &controllerTest{
		sdk: api.SDK(t),
		cloudRes: &testResources{
			t:                    t,
			folderID:             baseResources.FolderID,
			networkFolderID:      baseResources.FolderID,
			clusterFolderID:      baseResources.FolderID,
			networkID:            uuid.New().String(),
			zoneID:               zoneID,
			diskTypeID:           diskTypeID,
			diskName:             diskName,
			defaultDiskSizeBytes: defaultDiskSizeBytes,
			testData:             &testData{disk: []*compute.Disk{}, instance: []*compute.Instance{}},
		},
	}

	test.diskapi = NewPublicDiskAPI(test.sdk, baseResources.FolderID)
	test.controller = controller.New(test.diskapi, inflight.NewWithTTL(), services.ProductionCapsForPublicAPIController)
	test.cleanup = func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		require.NoError(t, test.sdk.Shutdown(ctx))
	}
	test.cloudRes.testData.instance = append(test.cloudRes.testData.instance, computeInstance(test.cloudRes.folderID))
	initializeMocks(api, test.cloudRes)
	return test
}

func initializeMocks(api *pubapimock.API, testResources *testResources) {
	diskCreateMock(api, testResources.folderID, testResources.testData)
	diskUpdateMock(api, testResources.folderID, testResources.testData)
	diskDeleteMock(api, testResources.testData)
	diskGetMock(api, testResources.testData)
	diskListMock(api, testResources.testData)
	instanceGetMock(api, testResources.testData)
	instanceAttachDiskMock(api, testResources.testData)
	instanceDetachDiskMock(api, testResources.testData)
	instanceUpdateLabelsMock(api, testResources.testData)
}
