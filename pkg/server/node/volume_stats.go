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

package node

import (
	"syscall"

	"github.com/yandex-cloud/yc-csi-driver/pkg/services/node"
)

type VolumeStats struct{}

func (v *VolumeStats) GetVolumeStats(target string) (byteStats, inodesStats *node.VolStats, err error) {
	var fs syscall.Statfs_t

	err = syscall.Statfs(target, &fs)
	if err == nil {
		byteStats = &node.VolStats{
			Total:     int64(fs.Blocks * uint64(fs.Bsize)),
			Used:      int64((fs.Blocks - fs.Bfree) * uint64(fs.Bsize)),
			Available: int64(fs.Bavail * uint64(fs.Bsize)),
		}

		inodesStats = &node.VolStats{
			Total:     int64(fs.Files),
			Used:      int64(fs.Files - fs.Ffree),
			Available: int64(fs.Ffree),
		}
	}

	return
}
