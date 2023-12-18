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

## Running tests

This repository has unit and integration tests for the charts. All charts are also linted.

### Linting

Linting is done with `helm lint`.

Prerequisites:

* Helm

```bash
brew install helm
```

To lint all charts:

```bash
make lint
```

### Unit tests

Unit tests are in the `./test` directory and written in Go with [terratest](https://github.com/gruntwork-io/terratest).

Prerequisites:

* Go

```bash
brew install go
```

To run the tests:

```bash
make test-unit
```

### Integration tests

Integration tests are in the `./test/integration` directory and written with [bats](https://bats-core.readthedocs.io/).

There is an additional simple Helm test in `./charts/qdrant/templates/tests`.

Prerequisites:

* Docker
* Kind
* Kubectl
* Helm
* Bats

```bash
brew install helm kubectl kind bats-core homebrew/cask/docker
```

To run the tests:

```bash
make test-integration
```

## Releasing

Generally, we choose to release a new chart when there are important new releases to Qdrant or important chart-related changes.

1. Check the release notes of [qdrant/qdrant](https://github.com/qdrant/qdrant/releases) to see if any environment variables or other settings need to be adjusted.
2. Sync your local clone/fork of the [qdrant-helm](https://github.com/qdrant/qdrant-helm) repo: `git checkout main && git pull`
3. Checkout a new feature branch: `git checkout -b feat/<name>/qdrant-<version>`
4. Edit `Chart.yaml` to bump the appVersion and chartVersion.
5. Edit `Chart.yaml` to update `artifacthub.io/changes` to mention the new changes.
6. Edit `charts/CHANGELOG.md` to mention the same changes
7. Edit the root `CHANGELOG.md` to mention the same changes
    1. Why so many changelog changes? Each changelog file is for a different audience (artifacthub, chart browsing, github browsing). This could be automated in the future.
8. Push your changes to GitHub and create a pull request. This allows the integration tests to run.

As soon as these changes are merged to main, there is a [github action that detects changes to Chart.yaml](https://github.com/qdrant/qdrant-helm/blob/cea92d092ac330493147536e27f3edeb465ffe75/.github/workflows/release-workflow.yaml#L7) which will perform the remainder of the release operations (creating a github release, publishing the helm chart, updating index.html for the [github pages site](https://qdrant.github.io/qdrant-helm/), etc.)
