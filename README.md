# Qdrant helm chart

[Qdrant documentation](https://qdrant.tech/documentation/) 

## TLDR;

Checkout the repository and

```bash
helm install your-qdrant-installation-name .
```

## Description

This chart installs and bootstraps a Qdrant instance.


## Prerequisites

- Kubernetes
- Helm
- PV provisioner (by the infrastructure)

## Installation & Setup

You can install the chart via:

```bash
helm install your-qdrant-installation-name .
```

Unistall via:

```bash
helm delete your-qdrant-installation-name .
```

Delete the volume with

```bash
kubectl delete pvc -l kubectl delete pvc -l app.kubernetes.io/instance=your-qdrant-installation-name
```

## Configuration

For documentation of the settings please refer to [Qdrant Configuration File](https://github.com/qdrant/qdrant/blob/master/config/config.yaml)
All of these configuration options could be overwritten under config in `values.yaml`. 
A modifcation example is provided there.

### Distributed setup

Running a distributed cluster just needs a few changes in your `values.yaml` file.
Increase the number of replicas to the desired number of nodes and set `config.cluster.enabled` to true.

Depending on your environment or cloud provider you might need to change the service in the `values.yaml` as well.
For example on AWS EKS you would need to change the `cluster.type` to `NodePort`.
