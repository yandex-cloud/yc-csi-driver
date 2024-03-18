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

package sanitytests

import (
	"context"
	"path/filepath"
	"runtime"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-csi/csi-test/v4/pkg/sanity"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"k8s.io/mount-utils"

	"github.com/yandex-cloud/yc-csi-driver/pkg/services/node"
)

var (
	mounter = node.NewFakeSafeFormatAndMounter(mount.NewFakeMounter([]mount.MountPoint{}))
)

func runControllerTest(
	sc *sanity.TestContext,
	r *sanity.Resources,
	controllerPublishSupported bool,
	nodeStageSupported bool,
	nodeVolumeStatsSupported bool,
	count int,
) {
	name := sanity.UniqueString("sanity-node-full")

	ginkgo.By("getting node information")
	ni, err := r.NodeGetInfo(
		context.Background(),
		&csi.NodeGetInfoRequest{})
	gomega.Expect(err).NotTo(gomega.HaveOccurred())
	gomega.Expect(ni).NotTo(gomega.BeNil())
	gomega.Expect(ni.GetNodeId()).NotTo(gomega.BeEmpty())

	var accReqs *csi.TopologyRequirement

	if ni.AccessibleTopology != nil {
		// Topology requirements are honored if provided by the driver
		accReqs = &csi.TopologyRequirement{
			Requisite: []*csi.Topology{ni.AccessibleTopology},
		}
	}

	// Create Volume First
	ginkgo.By("creating a single node writer volume")
	vol := r.MustCreateVolume(
		context.Background(),
		&csi.CreateVolumeRequest{
			Name: name,
			VolumeCapabilities: []*csi.VolumeCapability{
				sanity.TestVolumeCapabilityWithAccessType(sc, csi.VolumeCapability_AccessMode_SINGLE_NODE_WRITER),
			},
			CapacityRange: &csi.CapacityRange{
				RequiredBytes: sanity.TestVolumeSize(sc),
			},
			Secrets:                   sc.Secrets.CreateVolumeSecret,
			Parameters:                sc.Config.TestVolumeParameters,
			AccessibilityRequirements: accReqs,
		},
	)

	var conpubvol *csi.ControllerPublishVolumeResponse

	if controllerPublishSupported {
		ginkgo.By("controller publishing volume")

		conpubvol, err = r.ControllerPublishVolume(
			context.Background(),
			&csi.ControllerPublishVolumeRequest{
				VolumeId:         vol.GetVolume().GetVolumeId(),
				NodeId:           ni.GetNodeId(),
				VolumeCapability: sanity.TestVolumeCapabilityWithAccessType(sc, csi.VolumeCapability_AccessMode_SINGLE_NODE_WRITER),
				VolumeContext:    vol.GetVolume().GetVolumeContext(),
				Readonly:         false,
				Secrets:          sc.Secrets.ControllerPublishVolumeSecret,
			},
		)
		gomega.Expect(err).NotTo(gomega.HaveOccurred())
		gomega.Expect(conpubvol).NotTo(gomega.BeNil())
	}

	// NodeStageVolume
	if nodeStageSupported {
		for i := 0; i < count; i++ {
			ginkgo.By("node staging volume")
			nodestagevol, err := r.NodeStageVolume(
				context.Background(),
				&csi.NodeStageVolumeRequest{
					VolumeId:          vol.GetVolume().GetVolumeId(),
					VolumeCapability:  sanity.TestVolumeCapabilityWithAccessType(sc, csi.VolumeCapability_AccessMode_SINGLE_NODE_WRITER),
					StagingTargetPath: sc.StagingPath,
					VolumeContext:     vol.GetVolume().GetVolumeContext(),
					PublishContext:    conpubvol.GetPublishContext(),
					Secrets:           sc.Secrets.NodeStageVolumeSecret,
				},
			)
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(nodestagevol).NotTo(gomega.BeNil())
		}
	}

	// NodePublishVolume
	var stagingPath string

	if nodeStageSupported {
		stagingPath = sc.StagingPath
	}

	for i := 0; i < count; i++ {
		targetMps := 0

		ginkgo.By("publishing the volume on a node")

		targetPath := sc.TargetPath + "/target"

		nodepubvol, err := r.NodePublishVolume(
			context.Background(),
			&csi.NodePublishVolumeRequest{
				VolumeId:          vol.GetVolume().GetVolumeId(),
				TargetPath:        targetPath,
				StagingTargetPath: stagingPath,
				VolumeCapability:  sanity.TestVolumeCapabilityWithAccessType(sc, csi.VolumeCapability_AccessMode_SINGLE_NODE_WRITER),
				VolumeContext:     vol.GetVolume().GetVolumeContext(),
				PublishContext:    conpubvol.GetPublishContext(),
				Secrets:           sc.Secrets.NodePublishVolumeSecret,
			},
		)

		checkTargetPath := targetPath
		if runtime.GOOS == "darwin" {
			checkTargetPath = filepath.Join("/private/", targetPath)
		}

		gomega.Expect(err).NotTo(gomega.HaveOccurred())
		gomega.Expect(nodepubvol).NotTo(gomega.BeNil())

		mps, err := mounter.List()

		gomega.Expect(err).NotTo(gomega.HaveOccurred())

		for _, mp := range mps {
			if mp.Path == checkTargetPath {
				targetMps++
			}
		}
		gomega.Expect(targetMps).Should(gomega.BeIdenticalTo(1))
	}

	// NodeGetVolumeStats
	if nodeVolumeStatsSupported {
		ginkgo.By("Get node volume stats")
		statsResp, err := r.NodeGetVolumeStats(
			context.Background(),
			&csi.NodeGetVolumeStatsRequest{
				VolumeId:   vol.GetVolume().GetVolumeId(),
				VolumePath: sc.TargetPath + "/target",
			},
		)
		gomega.Expect(err).ToNot(gomega.HaveOccurred())
		gomega.Expect(statsResp.GetUsage()).ToNot(gomega.BeNil())
	}
}

var _ = sanity.DescribeSanity("YC Node Service", func(sc *sanity.TestContext) {
	var (
		r *sanity.Resources
	)

	ginkgo.BeforeEach(func() {
		cl := csi.NewControllerClient(sc.ControllerConn)
		n := csi.NewNodeClient(sc.Conn)

		i := csi.NewIdentityClient(sc.Conn)
		req := &csi.GetPluginCapabilitiesRequest{}
		res, err := i.GetPluginCapabilities(context.Background(), req)
		gomega.Expect(err).NotTo(gomega.HaveOccurred())
		gomega.Expect(res).NotTo(gomega.BeNil())
		r = &sanity.Resources{
			Context:          sc,
			ControllerClient: cl,
			NodeClient:       n,
		}
	})

	ginkgo.AfterEach(func() {
		r.Cleanup()
	})

	ginkgo.It("YC Disk Idempotency Test", func() {
		count := sc.Config.IdempotentCount
		ginkgo.By("runControllerTest with Idempotent count")
		runControllerTest(sc, r, true, true, false, count)
	})

})
