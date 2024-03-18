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

package services

import (
	"crypto/sha1"
	"fmt"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func VerifyVolumeCapabilities(volCaps []*csi.VolumeCapability, volumeCaps []*csi.VolumeCapability_AccessMode) error {
	var modes []string

	hasSupport := func(cap *csi.VolumeCapability) bool {
		for _, c := range volumeCaps {
			if c.GetMode() == cap.AccessMode.GetMode() {
				return true
			}

			modes = append(modes, cap.GetAccessMode().GetMode().String())
		}

		return false
	}

	foundAll := true

	for _, c := range volCaps {
		if !hasSupport(c) {
			foundAll = false
		}
	}

	if !foundAll {
		return status.Errorf(codes.InvalidArgument, "Volume AccessModes (%+v) not supported", modes)
	}

	return nil
}

func CheckVolumeAccessType(volCap *csi.VolumeCapability) error {
	switch volCap.GetAccessType().(type) {
	case *csi.VolumeCapability_Block, *csi.VolumeCapability_Mount:
		return nil
	default:
		return status.Error(codes.InvalidArgument,
			"Only VolumeCapability_Block or VolumeCapability_Mount can be used")
	}
}

func CSIToYCName(volumeName string) string {
	// CSI spec requirements for names differ from those of cloud compute.
	// For example, csi-sanity generates names with capital chars forbidden by cloud compute.
	// It also checks "creating volume with maximum-length name", this is equal to 128, according to
	// https://github.com/container-storage-interface/spec/blob/master/spec.md#size-limits
	// while cloud compute allows a maximum of 63 characters.
	h := sha1.New()
	h.Write([]byte(volumeName))
	// SHA1 hash is 20 bytes, we'll fit into 63 chars.
	return fmt.Sprintf("k8s-csi-%x", h.Sum(nil))
}
