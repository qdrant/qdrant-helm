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

## Limitations

You can use this to run Qdrant in your Kubernetes cluster. While it is possible to deploy Qdrant in a distributed setup with the Helm chart, it does not come with the same level of features for zero-downtime upgrades, up and down-scaling, monitoring, logging, and backup and disaster recovery as the Qdrant Cloud offering or the Qdrant Private Cloud Enterprise Operator. Support for the Helm chart is limited to community support.

The following table gives you an overview about the feature differences between the Qdrant Cloud and the Helm chart:

| Feature                                                | Qdrant Helm Chart | Qdrant Cloud  |
|--------------------------------------------------------|:-----------------:|:-------------:|
| Open-source                                            | ✅                |               |
| Community support only                                 | ✅                |               |
| Quick to get started                                   | ✅                | ✅            |
| Vertical and horizontal scaling                        | ✅                | ✅            |
| API keys with granular access control                  | ✅                | ✅            |
| Qdrant version upgrades                                | ✅                | ✅            |
| Support for transit and storage encryption             | ✅                | ✅            |
| Zero-downtime upgrades with optimized restart strategy |                   | ✅            |
| Production ready out-of the box                        |                   | ✅            |
| Dataloss prevention on downscaling                     |                   | ✅            |
| Automatic shard rebalancing                            |                   | ✅            |
| Full cluster backup and disaster recovery              |                   | ✅            |
| Re-sharding support                                    |                   | ✅            |
| Automatic persistent volume scaling                    |                   | ✅            |
| Advanced telemetry                                     |                   | ✅            |
| One-click API key revoking                             |                   | ✅            |
| Recreating nodes with new volumes in existing cluster  |                   | ✅            |
| Enterprise support                                     |                   | ✅            |

For more information on Qdrant Cloud, including the Hybrid Cloud and Private Cloud deployment options, see [Qdrant Cloud](https://qdrant.tech/cloud/).

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
  tag: vX.Y.Z
```

## Contributing

See [CONTRIBUTING.md](./CONTRIBUTING.md).
