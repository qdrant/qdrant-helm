# Changelog

## [qdrant-0.9.0](https://github.com/qdrant/qdrant-helm/tree/qdrant-0.9.0) (2024-05-06)

- Update Qdrant to v1.9.1
- Add labels to ConfigMap and Secret [#174](https://github.com/qdrant/qdrant-helm/pull/174)
- Make lifecycle hooks configurable [#175](https://github.com/qdrant/qdrant-helm/pull/175)
- Support storing snapshots in a separate PVC [#177](https://github.com/qdrant/qdrant-helm/pull/177)
- Add both storage and snapshot volume names to values.yaml [#181](https://github.com/qdrant/qdrant-helm/pull/181)

## [qdrant-0.8.5](https://github.com/qdrant/qdrant-helm/tree/qdrant-0.8.5) (2024-04-25)

- Update Qdrant to v1.9.0
- Allow ports to be configured as NodePort [#165](https://github.com/qdrant/qdrant-helm/pull/165)

## [qdrant-0.8.4](https://github.com/qdrant/qdrant-helm/tree/qdrant-0.8.4) (2024-04-03)

- Update Qdrant to v1.8.4
- Allow configurable podManagementPolicy [#158](https://github.com/qdrant/qdrant-helm/pull/158)


## [qdrant-0.8.3](https://github.com/qdrant/qdrant-helm/tree/qdrant-0.8.3) (2024-03-19)

- Update Qdrant to v1.8.3

## [qdrant-0.8.2](https://github.com/qdrant/qdrant-helm/tree/qdrant-0.8.2) (2024-03-18)

- Update Qdrant to v1.8.2

## [qdrant-0.8.1](https://github.com/qdrant/qdrant-helm/tree/qdrant-0.8.1) (2024-03-09)

- Update Qdrant to v1.8.1

## [qdrant-0.8.0](https://github.com/qdrant/qdrant-helm/tree/qdrant-0.8.0) (2024-03-06)

- Update Qdrant to v1.8.0
- Fix typos in service template [#147](https://github.com/qdrant/qdrant-helm/pull/147)
- Support read_only_api_key in Qdrant config [#146](https://github.com/qdrant/qdrant-helm/pull/146)

## [qdrant-0.7.6](https://github.com/qdrant/qdrant-helm/tree/qdrant-0.7.6) (2024-01-30)

- Update Qdrant to v1.7.4

## [qdrant-0.7.5](https://github.com/qdrant/qdrant-helm/tree/qdrant-0.7.5) (2024-01-17)

- Configurable static IP for LoadBalancer services [#122](https://github.com/qdrant/qdrant-helm/pull/122)
- Fix metrics scraping if api key is enabled [#127](https://github.com/qdrant/qdrant-helm/pull/127)
- Use /readyz endpoint for readiness probe for Qdrant >= 1.7.3 [#124](https://github.com/qdrant/qdrant-helm/pull/124)
- Use FQDN for container image [#125](https://github.com/qdrant/qdrant-helm/pull/125)

## [qdrant-0.7.4](https://github.com/qdrant/qdrant-helm/tree/qdrant-0.7.4) (2023-12-29)

- Update Qdrant to v1.7.3
- Add preStop hook to allow graceful network shutdown [#121](https://github.com/qdrant/qdrant-helm/pull/121)

## [qdrant-0.7.3](https://github.com/qdrant/qdrant-helm/tree/qdrant-0.7.3) (2023-12-14)

- Update Qdrant to v1.7.2
- Publish DNS for not-ready pods via the headless service [#115](https://github.com/qdrant/qdrant-helm/pull/115)
-
## [qdrant-0.7.2](https://github.com/qdrant/qdrant-helm/tree/qdrant-0.7.2) (2023-12-12)

- Use / for probes instead of /readyz endpoint

## [qdrant-0.7.1](https://github.com/qdrant/qdrant-helm/tree/qdrant-0.7.1) (2023-12-12)

- Update Qdrant to v1.7.1

## [qdrant-0.7.0](https://github.com/qdrant/qdrant-helm/tree/qdrant-0.7.0) (2023-12-08)

- Update Qdrant to v1.7.0
- Fix snapshot restoration [#96](https://github.com/qdrant/qdrant-helm/pull/96)

## [qdrant-0.6.1](https://github.com/qdrant/qdrant-helm/tree/qdrant-0.6.1) (2023-10-12)

- Propagate signals in init script correctly to entrypoint [#91](https://github.com/qdrant/qdrant-helm/pull/91)
- Make annotations for the ServiceAccount configurable [#92](https://github.com/qdrant/qdrant-helm/pull/92)
- Update Qdrant to v1.6.1

## [qdrant-0.6.0](https://github.com/qdrant/qdrant-helm/tree/qdrant-0.6.0) (2023-10-09)

- Make ingressClassName configurable [#86](https://github.com/qdrant/qdrant-helm/pull/86)
- Fix probes to work correctly if TLS is enabled [#79](https://github.com/qdrant/qdrant-helm/pull/79)
- Update Qdrant to v1.6.0

## [qdrant-0.5.1](https://github.com/qdrant/qdrant-helm/tree/qdrant-0.5.1) (2023-09-12)

- Update Qdrant to v1.5.1
- Ensure that the qdrant-init-file-path is on a writable, ephemeral volume [\#74](https://github.com/qdrant/qdrant-helm/issues/74)

## [qdrant-0.5.0](https://github.com/qdrant/qdrant-helm/tree/qdrant-0.5.0) (2023-09-07)

- Update to Qdrant 1.5.0 [\#72](https://github.com/qdrant/qdrant-helm/issues/72)
- Use new Qdrant readiness and liveness endpoints [\#71](https://github.com/qdrant/qdrant-helm/issues/71)

## [qdrant-0.4.1](https://github.com/qdrant/qdrant-helm/tree/qdrant-0.4.1) (2023-09-04)

- Add PriorityClass support to StatefulSet Pod template [\#68](https://github.com/qdrant/qdrant-helm/pull/68)
- Don't use alpine image for file permission updates [\#69](https://github.com/qdrant/qdrant-helm/pull/69)

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
