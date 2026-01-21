# Yandex Cloud CSI Driver

## About
A Container Storage Interface Driver for Yandex Compute Disks.

## Kubernetes compatibility

| Driver Version | Supported k8s version |
|----------------|-----------------------|
| v1.0.0         | 1.21+                 |
| v1.1.0         | 1.21+                 |
| v1.2.0         | 1.21+                 |

## Install on a Kubernetes cluster

### 1. Create secret with Yandex Cloud service account key
You need a service account with `k8s.clusters.agent` role.

Create a new kubernetes secret using the following commands:

```
$ yc iam key create -o sa-key.json --service-account-id=<sa-id> --format json 
{
  "id": "<key-id>",
  "service_account_id": "<sa-id>",
...

$ kubectl create secret generic yc-csi-sa-key --from-file=sa-key.json -n kube-system
secret/yc-csi-sa-key created
```

### 2. Create config map with csi config
Create a new configmap and provide folder where resources will be created:

```
$ kubectl apply -f - <<EOF
  apiVersion: v1
  kind: ConfigMap
  metadata:
    name: yc-csi-config
    namespace: kube-system
  data:
    folderId: <folder-id>
EOF

configmap/yc-csi-config created
```

### 3. Deploy CSI driver
Deploy the driver using the following command:

```
$ kubectl apply -f deploy/vX.Y.Z
```
