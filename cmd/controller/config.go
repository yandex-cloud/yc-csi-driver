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

package main

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"

	"github.com/yandex-cloud/yc-csi-driver/pkg/diskapi/public"
	"github.com/yandex-cloud/yc-csi-driver/pkg/server/controller"
	"github.com/yandex-cloud/yc-csi-driver/pkg/services"

	ycsdk "github.com/yandex-cloud/go-sdk"
	"github.com/yandex-cloud/go-sdk/iamkey"
)

func NewControllerOptions(endpoint, cloudFolderID, saKeyFilePath, apiEndpoint string) (*controller.Options, error) {
	if len(saKeyFilePath) == 0 {
		return nil, errors.New("no saKeyFilePath provided")
	}

	iamKey, err := iamkey.ReadFromJSONFile(saKeyFilePath)
	if err != nil {
		return nil, fmt.Errorf("read iamKey from file failed: %w", err)
	}

	saKey, err := ycsdk.ServiceAccountKey(iamKey)
	if err != nil {
		return nil, fmt.Errorf("saKey from iamKey failed: %w", err)
	}

	sdk, err := ycsdk.Build(
		context.Background(),
		ycsdk.Config{
			Credentials: saKey,
			Endpoint:    apiEndpoint,
			TLSConfig:   &tls.Config{},
		},
	)

	if err != nil {
		return nil, err
	}

	diskAPI := public.NewPublicDiskAPI(sdk, cloudFolderID)

	return &controller.Options{
		Endpoint:     endpoint,
		DriverName:   "", // Not required.
		DiskAPI:      diskAPI,
		Capabilities: services.ProductionCapsForPublicAPIController,
	}, nil
}
