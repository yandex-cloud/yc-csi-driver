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
	"flag"
	"fmt"

	"github.com/yandex-cloud/yc-csi-driver/cmd"
	"github.com/yandex-cloud/yc-csi-driver/pkg/server/controller"

	"k8s.io/klog"
)

var (
	endpoint, cloudFolderID, saKeyFilePath, ycAPIEndpoint string
)

func init() {
	flag.StringVar(&endpoint, "endpoint", "unix://tmp/csi.sock", "CSI Endpoint")
	flag.StringVar(&ycAPIEndpoint, "yc-api-endpoint", "api.cloud.yandex.net:443",
		"cloud compute API endpoint")
	flag.StringVar(&cloudFolderID, "folder-id", "", "folder id to work in")
	flag.StringVar(&saKeyFilePath, "sa-key", "", "sa key file path for csi controller")
}

func main() {
	klog.InitFlags(nil)

	flag.Parse()

	err := run()
	if err != nil {
		klog.Fatalf("%+v", err)
	}
}

func run() error {
	opts, err := NewControllerOptions(endpoint, cloudFolderID, saKeyFilePath, ycAPIEndpoint)
	if err != nil {
		return err
	}

	srv, err := controller.New(*opts)
	if err != nil {
		return fmt.Errorf("server build: %w", err)
	}

	cmd.Execute(srv)

	return nil
}
