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
	"context"
	"io"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"k8s.io/kubernetes/pkg/volume/util/hostutil"
	"k8s.io/mount-utils"
	"k8s.io/utils/exec"

	"github.com/yandex-cloud/yc-csi-driver/pkg/instancemeta"
)

type VolStats struct {
	Total     int64
	Used      int64
	Available int64
}
type VolumeStats interface {
	GetVolumeStats(string) (*VolStats, *VolStats, error)
}

type FakeVolumeStats struct {
}

func (f FakeVolumeStats) GetVolumeStats(s string) (bytesStats, inodeStats *VolStats, err error) {
	bytesStats = &VolStats{}
	inodeStats = &VolStats{}
	return
}

func NewFakeMetadataGetter(instID, zoneID string) *FakeMetadataGetter {
	return &FakeMetadataGetter{instID: instID, zoneID: zoneID}
}

type FakeMetadataGetter struct {
	instID string
	zoneID string
}

func (f *FakeMetadataGetter) InstanceID() (string, error) {
	return f.instID, nil
}

func (f *FakeMetadataGetter) ZoneID() (string, error) {
	return f.zoneID, nil
}

func newFakeMounter() *mount.FakeMounter {
	return &mount.FakeMounter{
		MountPoints: []mount.MountPoint{},
	}
}

func NewFakeSafeFormatAndMounter(fakeMounter mount.Interface) *mount.SafeFormatAndMount {
	return &mount.SafeFormatAndMount{
		Interface: fakeMounter,
		Exec:      &fakeExec{runHook: fakeRunHook},
	}
}

func NewTestNodeService(mounter mount.Interface, hostUtil hostutil.HostUtils, fileXattrs FileXattrs,
	mdGetter instancemeta.MetadataGetter, caps []*csi.NodeServiceCapability) *Service {
	return New(
		NewFakeSafeFormatAndMounter(mounter),
		hostUtil,
		FakeVolumeStats{},
		fileXattrs,
		mdGetter,
		caps,
		DefaultOptions(),
	)
}

var _ exec.Interface = &fakeExec{}

// fakeExec for testing.
type fakeExec struct {
	runHook runHook
}

func (f *fakeExec) Command(cmd string, args ...string) exec.Cmd {
	return &fakeCmd{cmd: cmd, args: args, runHook: f.runHook}
}

func (f *fakeExec) CommandContext(ctx context.Context, cmd string, args ...string) exec.Cmd {
	panic("implement me")
}
func (f *fakeExec) LookPath(file string) (string, error) { panic("implement me") }

type runHook func(cmd string, args ...string) ([]byte, error)

type fakeCmd struct {
	cmd     string
	args    []string
	runHook runHook
}

func (f *fakeCmd) Run() error {
	_, err := f.CombinedOutput()
	return err
}

func (f *fakeCmd) CombinedOutput() ([]byte, error) {
	if f.runHook != nil {
		return f.runHook(f.cmd, f.args...)
	}
	return nil, nil
}

func fakeRunHook(cmd string, args ...string) ([]byte, error) {
	switch cmd {
	// Fake blkid cmd output.
	case "blkid":
		return []byte(`DEVNAME=/dev/vdb` + "\n" + `TYPE=ext4`), nil
	}

	return nil, nil
}

func (f *fakeCmd) Output() ([]byte, error)            { panic("implement me") }
func (f *fakeCmd) SetDir(dir string)                  { panic("implement me") }
func (f *fakeCmd) SetStdin(in io.Reader)              { panic("implement me") }
func (f *fakeCmd) SetStdout(out io.Writer)            { panic("implement me") }
func (f *fakeCmd) SetStderr(out io.Writer)            { panic("implement me") }
func (f *fakeCmd) SetEnv(env []string)                { panic("implement me") }
func (f *fakeCmd) StdoutPipe() (io.ReadCloser, error) { panic("implement me") }
func (f *fakeCmd) StderrPipe() (io.ReadCloser, error) { panic("implement me") }
func (f *fakeCmd) Start() error                       { panic("implement me") }
func (f *fakeCmd) Wait() error                        { panic("implement me") }
func (f *fakeCmd) Stop()                              { panic("implement me") }

var _ exec.Cmd = &fakeCmd{}
