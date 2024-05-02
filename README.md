# Qdrant helm repository

[Qdrant documentation](https://qdrant.tech/documentation/)

This repository hosts the following helm charts:

* [Qdrant](charts/qdrant/README.md)

## Usage

```bash
helm repo add qdrant https://qdrant.github.io/qdrant-helm
helm repo update
helm upgrade -i your-qdrant-installation-name qdrant/qdrant
```

## Upgrading

This helm chart installs the latest version of Qdrant by default. When a new version of Qdrant is available, upgrade the helm chart with the following commands:

```bash
helm repo update
helm upgrade your-qdrant-installation-name qdrant/qdrant
```

This command performs a rolling upgrade of your Qdrant cluster, updating one node at a time.

If you have overridden the Qdrant image tag in `values.yaml`, you will also need to update that tag before running `helm upgrade`.

```yaml
image:
  tag: v1.9.0
```

## Restoring from Snapshots

This helm chart allows you to restore a snapshot into your Qdrant cluster either from an internal or external PersistentVolumeClaim.

### Restoring from the built-in PVC

If you have set `snapshotPersistence.enabled: true` (recommended for production), this helm chart will create a separate PersistentVolume for snapshots, and any snapshots you create will be stored in that PersistentVolume.

To restore from one of these snapshots, set the following values:

```yaml
snapshotRestoration:
  enabled: true
  # Set blank to indicate we are not using an external PVC
  pvcName: ""
  snapshots:
  - /qdrant/snapshots/<collection_name>/<filename>/:<collection_name>
```

And run "helm upgrade". This will restart your cluster and restore the specified collection from the snapshot. Qdrant will refuse to overwrite an existing collection, so ensure the collection is deleted before restoring.

After the snapshot is restored, remove the above values and run "helm upgrade" again to trigger another rolling restart. Otherwise, the snapshot restore will be attempted again if your cluster ever restarts.

### Restoring from an external PVC

If you wish to restore from an externally-created snapshot, using the API is recommended: https://qdrant.github.io/qdrant/redoc/index.html#tag/collections/operation/recover_from_uploaded_snapshot

If the file is too large, you can separatly create a PersistentVolumeClaim, store your data in there, and refer to this separate PersistentVolumeClaim in this helm chart.

Once you have created this PersistentVolumeClaim (must be in the same namespace as your Qdrant cluster), set the following values:

```
snapshotRestoration:
  enabled: true
  pvcName: "<the name of your PVC>"
  snapshots:
  - /qdrant/snapshots/<collection_name>/<filename>/:<collection_name>
```

And run "helm upgrade". This will restart your cluster and restore the specified collection from the snapshot. Qdrant will refuse to overwrite an existing collection, so ensure the collection is deleted before restoring.

After the snapshot is restored, remove the above values and run "helm upgrade" again to trigger another rolling restart. Otherwise, the snapshot restore will be attempted again if your cluster ever restarts.

## Contributing

See [CONTRIBUTING.md](./CONTRIBUTING.md).
