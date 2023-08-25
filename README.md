# Qdrant helm repository

[Qdrant documentation](https://qdrant.tech/documentation/)

This repository hosts the following helm charts:

* [Qdrant](charts/qdrant/README.md)

## Usage

```bash
helm repo add qdrant https://qdrant.github.io/qdrant-helm
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
