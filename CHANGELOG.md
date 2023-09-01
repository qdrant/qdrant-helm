# Changelog

## [qdrant-0.4.0](https://github.com/qdrant/qdrant-helm/tree/qdrant-0.4.0) (2023-09-01)

- Make it possible to set topologySpreadConstraints [\#65](https://github.com/qdrant/qdrant-helm/issues/65)
- Add support for additional labels on StatefulSets  [\#55](https://github.com/qdrant/qdrant-helm/issues/55)
- Enable `entrypoint.sh` to add recovery mode functionality [\#48](https://github.com/qdrant/qdrant-helm/issues/48)
- Configurable Qdrant API key [\#46](https://github.com/qdrant/qdrant-helm/issues/46)
- Add PodDisruptionBudget [\#36](https://github.com/qdrant/qdrant-helm/issues/36)
- Allow mounting of additional volumes [\#46](https://github.com/qdrant/qdrant-helm/issues/46)
- Add default ServiceAccount [\#56](https://github.com/qdrant/qdrant-helm/pull/56)
- Support for Qdrant recovery mode [\#54](https://github.com/qdrant/qdrant-helm/pull/54)
- Make relabeling and metricRelabeling configurable in ServiceMonitor [\#61](https://github.com/qdrant/qdrant-helm/pull/61)
- Allow annotations on volumeClaimTemplate of Qdrant StatefulSet [\#45](https://github.com/qdrant/qdrant-helm/issues/45)
- Add default container and pod securityContexts and make it configurable [\#60](https://github.com/qdrant/qdrant-helm/pull/60)

## [qdrant-0.3.1](https://github.com/qdrant/qdrant-helm/tree/qdrant-0.3.1) (2023-08-23)

- Change target port for serviceMonitor to rely on new naming \(http\) [\#58](https://github.com/qdrant/qdrant-helm/pull/58)
