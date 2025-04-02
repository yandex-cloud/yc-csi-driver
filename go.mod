module github.com/yandex-cloud/yc-csi-driver

go 1.23.0

toolchain go1.23.7

require (
	github.com/container-storage-interface/spec v1.11.0
	github.com/google/uuid v1.6.0
	github.com/kubernetes-csi/csi-test/v4 v4.4.0
	github.com/mitchellh/go-testing-interface v1.14.1
	github.com/onsi/ginkgo v1.16.5
	github.com/onsi/gomega v1.36.3
	github.com/stretchr/testify v1.10.0
	github.com/yandex-cloud/go-genproto v0.0.0-20250325081613-cd85d9003939
	github.com/yandex-cloud/go-sdk v0.0.0-20250325134853-dcb34ef70818
	golang.org/x/net v0.38.0
	google.golang.org/grpc v1.71.0
	google.golang.org/protobuf v1.36.6
	k8s.io/apimachinery v0.32.3
	k8s.io/klog v1.0.0
	k8s.io/kubernetes v1.32.3
	k8s.io/mount-utils v0.32.3
	k8s.io/utils v0.0.0-20250321185631-1f6e0b77f77e
)

require (
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/fsnotify/fsnotify v1.8.0 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/golang-jwt/jwt/v4 v4.5.2 // indirect
	github.com/google/go-cmp v0.7.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/moby/sys/mountinfo v0.7.2 // indirect
	github.com/moby/sys/userns v0.1.0 // indirect
	github.com/nxadm/tail v1.4.11 // indirect
	github.com/opencontainers/selinux v1.12.0 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/stretchr/objx v0.5.2 // indirect
	golang.org/x/sys v0.31.0 // indirect
	golang.org/x/text v0.23.0 // indirect
	google.golang.org/genproto v0.0.0-20250324211829-b45e905df463 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20250324211829-b45e905df463 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250324211829-b45e905df463 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/klog/v2 v2.130.1 // indirect
)

replace (
	k8s.io/api => k8s.io/api v0.28.4
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.28.4
	k8s.io/apiserver => k8s.io/apiserver v0.28.4
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.28.4
	k8s.io/client-go => k8s.io/client-go v0.28.4
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.28.4
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.28.4
	k8s.io/code-generator => k8s.io/code-generator v0.28.4
	k8s.io/component-base => k8s.io/component-base v0.28.4
	k8s.io/component-helpers => k8s.io/component-helpers v0.28.4
	k8s.io/controller-manager => k8s.io/controller-manager v0.28.4
	k8s.io/cri-api => k8s.io/cri-api v0.28.4
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.28.4
	k8s.io/dynamic-resource-allocation => k8s.io/dynamic-resource-allocation v0.28.4
	k8s.io/endpointslice => k8s.io/endpointslice v0.28.4
	k8s.io/kms => k8s.io/kms v0.28.4
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.28.4
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.28.4
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.28.4
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.28.4
	k8s.io/kubectl => k8s.io/kubectl v0.28.4
	k8s.io/kubelet => k8s.io/kubelet v0.28.4
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.28.4
	k8s.io/metrics => k8s.io/metrics v0.28.4
	k8s.io/pod-security-admission => k8s.io/pod-security-admission v0.28.4
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.28.4
)
