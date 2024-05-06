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

For more in-depth usage documentation, see [the helm chart's README](charts/qdrant/README.md).

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
  tag: v1.X.Y
```

## Contributing

See [CONTRIBUTING.md](./CONTRIBUTING.md).
