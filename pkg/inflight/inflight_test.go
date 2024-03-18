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

package inflight

import (
	"sync"
	"testing"
	"time"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/stretchr/testify/require"

	"github.com/yandex-cloud/yc-csi-driver/pkg/util"
)

const (
	GiB = 1024 * 1024 * 1024
)

type testRequest struct {
	request *csi.CreateVolumeRequest
	expResp bool
	delete  bool
}

var stdVolCap = []*csi.VolumeCapability{
	{
		AccessType: &csi.VolumeCapability_Mount{
			Mount: &csi.VolumeCapability_MountVolume{},
		},
		AccessMode: &csi.VolumeCapability_AccessMode{
			Mode: csi.VolumeCapability_AccessMode_SINGLE_NODE_WRITER,
		},
	},
}

var (
	stdVolSize  = int64(5 * GiB)
	stdCapRange = &csi.CapacityRange{RequiredBytes: stdVolSize}
	stdParams   = map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
)

// NewTestInFlight instantiates a InFlightTTL structures with fake time.
func NewTestInFlight() *InFlightTTL {
	now := time.Now()

	return &InFlightTTL{
		lock:     new(sync.Mutex),
		inFlight: make(map[string]time.Time),
		ttl:      DefaultEntryTTL,
		now: func() time.Time {
			return now
		},
	}
}

func TestInFlight(t *testing.T) {
	testCases := []struct {
		name     string
		requests []testRequest
	}{
		{
			name: "success normal",
			requests: []testRequest{
				{
					request: &csi.CreateVolumeRequest{
						Name:               "random-vol-name",
						CapacityRange:      stdCapRange,
						VolumeCapabilities: stdVolCap,
						Parameters:         stdParams,
					},
					expResp: true,
				},
			},
		},
		{
			name: "success adding request with different name",
			requests: []testRequest{
				{
					request: &csi.CreateVolumeRequest{
						Name:               "random-vol-foobar",
						CapacityRange:      stdCapRange,
						VolumeCapabilities: stdVolCap,
						Parameters:         stdParams,
					},
					expResp: true,
				},
				{
					request: &csi.CreateVolumeRequest{
						Name:               "random-vol-name-foobar",
						CapacityRange:      stdCapRange,
						VolumeCapabilities: stdVolCap,
						Parameters:         stdParams,
					},
					expResp: true,
				},
			},
		},
		{
			name: "success adding request with different parameters",
			requests: []testRequest{
				{
					request: &csi.CreateVolumeRequest{
						Name:               "random-vol-name-foobar",
						CapacityRange:      stdCapRange,
						VolumeCapabilities: stdVolCap,
						Parameters:         map[string]string{"foo": "bar"},
					},
					expResp: true,
				},
				{
					request: &csi.CreateVolumeRequest{
						Name:               "random-vol-name-foobar",
						CapacityRange:      stdCapRange,
						VolumeCapabilities: stdVolCap,
					},
					expResp: true,
				},
			},
		},
		{
			name: "success adding request with different parameters",
			requests: []testRequest{
				{
					request: &csi.CreateVolumeRequest{
						Name:               "random-vol-name-foobar",
						CapacityRange:      stdCapRange,
						VolumeCapabilities: stdVolCap,
						Parameters:         map[string]string{"foo": "bar"},
					},
					expResp: true,
				},
				{
					request: &csi.CreateVolumeRequest{
						Name:               "random-vol-name-foobar",
						CapacityRange:      stdCapRange,
						VolumeCapabilities: stdVolCap,
						Parameters:         map[string]string{"foo": "baz"},
					},
					expResp: true,
				},
			},
		},
		{
			name: "failure adding copy of request",
			requests: []testRequest{
				{
					request: &csi.CreateVolumeRequest{
						Name:               "random-vol-name",
						CapacityRange:      stdCapRange,
						VolumeCapabilities: stdVolCap,
						Parameters:         stdParams,
					},
					expResp: true,
				},
				{
					request: &csi.CreateVolumeRequest{
						Name:               "random-vol-name",
						CapacityRange:      stdCapRange,
						VolumeCapabilities: stdVolCap,
						Parameters:         stdParams,
					},
					expResp: false,
				},
			},
		},
		{
			name: "success add, delete, add copy",
			requests: []testRequest{
				{
					request: &csi.CreateVolumeRequest{
						Name:               "random-vol-name",
						CapacityRange:      stdCapRange,
						VolumeCapabilities: stdVolCap,
						Parameters:         stdParams,
					},
					expResp: true,
				},
				{
					request: &csi.CreateVolumeRequest{
						Name:               "random-vol-name",
						CapacityRange:      stdCapRange,
						VolumeCapabilities: stdVolCap,
						Parameters:         stdParams,
					},
					expResp: false,
					delete:  true,
				},
				{
					request: &csi.CreateVolumeRequest{
						Name:               "random-vol-name",
						CapacityRange:      stdCapRange,
						VolumeCapabilities: stdVolCap,
						Parameters:         stdParams,
					},
					expResp: true,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db := NewTestInFlight()
			for _, r := range tc.requests {
				var resp bool
				if r.delete {
					db.Delete(r.request)
				} else {
					resp = db.Insert(r.request)
				}
				if r.expResp != resp {
					t.Fatalf("expected insert to be %+v, got %+v", r.expResp, resp)
				}
			}
		})

	}
}

func TestInFlight_GetNotExpired(t *testing.T) {
	db := NewTestInFlight()
	hash := util.StringToStringer("random-vol-name").String()
	db.inFlight[hash] = db.now().Add(-9 * time.Minute)
	require.True(t, db.Get(util.StringToStringer("random-vol-name")))
}

func TestInFlight_GetExpired(t *testing.T) {
	db := NewTestInFlight()
	hash := util.StringToStringer("random-vol-name").String()
	db.inFlight[hash] = db.now().Add(-11 * time.Minute)
	require.False(t, db.Get(util.StringToStringer("random-vol-name")))
}
