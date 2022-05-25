# Qdrant helm chart

[Qdrant documentation](https://qdrant.tech/documentation/) 

## TLDR;

Checkout the repository and

```bash
helm install your_qdrant_installation_name .
```

## Description

This chart installs and bootstraps a Qdrant instance.

### Disclaimer
The current version (0.7.0) does not yet support a distributed setup or replication. 
Increasing replicas does not have the intended effect right now

## Prerequisites

- Kubernetes
- Helm
- PV provisioner (by the infrastructure)

## Installation & Setup

You can install the chart via:

```bash
helm install your_qdrant_installation_name .
```

Unistall via:

```bash
helm delete your_qdrant_installation_name .
```

Delete the volume with

```bash
kubectl delete pvc -l kubectl delete pvc -l app.kubernetes.io/instance=your_qdrant_installation_name
```

## Configuration

For documentation of the settings please refer to [Qdrant Configuration File](https://github.com/qdrant/qdrant/blob/master/config/config.yaml)
All of these configuration options are available in `values.yaml`.