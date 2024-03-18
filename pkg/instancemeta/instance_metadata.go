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

package instancemeta

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"

	ycsdk "github.com/yandex-cloud/go-sdk"
)

type MetadataGetter interface {
	InstanceID() (string, error)
	ZoneID() (string, error)
}

func New() *InstanceMetadata {
	return &InstanceMetadata{}
}

type InstanceMetadata struct{}

func (a *InstanceMetadata) InstanceID() (string, error) {
	return getMetadata("instance-id")
}

func (a *InstanceMetadata) ZoneID() (string, error) {
	return getMetadata("placement/availability-zone")
}

func getMetadata(pathSuffix string) (string, error) {
	u := url.URL{
		Scheme: "http",
		Host:   ycsdk.InstanceMetadataAddr,
		Path:   path.Join("/latest/meta-data", pathSuffix),
	}

	resp, err := http.Get(u.String())
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to acquire metadata '%s': %d", pathSuffix, resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read metadata '%s' body: %w", pathSuffix, err)
	}

	return string(bytes), nil
}
