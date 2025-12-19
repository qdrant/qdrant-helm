# Changelog

## [qdrant-1.16.3](https://github.com/qdrant/qdrant-helm/tree/qdrant-1.16.3) (2025-12-19)

- Update Qdrant to v1.16.3

## [qdrant-1.16.2](https://github.com/qdrant/qdrant-helm/tree/qdrant-1.16.2) (2025-12-03)

- Update Qdrant to v1.16.2

## [qdrant-1.16.1](https://github.com/qdrant/qdrant-helm/tree/qdrant-1.16.1) (2025-11-26)

- Update Qdrant to v1.16.1
- Add custom authorization support for ServiceMonitor [#407](https://github.com/qdrant/qdrant-helm/pull/407)

## [qdrant-1.16.0](https://github.com/qdrant/qdrant-helm/tree/qdrant-1.16.0) (2025-11-17)

- Update Qdrant to v1.16.0
- Add support for `volumeAttributesClass` for PVCs [#396](https://github.com/qdrant/qdrant-helm/pull/396)
- Fix helm warning because of wrong data type in env default value [#398](https://github.com/qdrant/qdrant-helm/pull/398)

## [qdrant-1.15.5](https://github.com/qdrant/qdrant-helm/tree/qdrant-1.15.5) (2025-09-30)

- Update Qdrant to v1.15.5

## [qdrant-1.15.4](https://github.com/qdrant/qdrant-helm/tree/qdrant-1.15.4) (2025-08-27)

- Update Qdrant to v1.15.4

## [qdrant-1.15.3](https://github.com/qdrant/qdrant-helm/tree/qdrant-1.15.3) (2025-08-14)

- Update Qdrant to v1.15.3

## [qdrant-1.15.2](https://github.com/qdrant/qdrant-helm/tree/qdrant-1.15.2) (2025-08-12)

- Update Qdrant to v1.15.2

## [qdrant-1.15.1](https://github.com/qdrant/qdrant-helm/tree/qdrant-1.15.1) (2025-07-24)

- Update Qdrant to v1.15.1
- Add support for additional annotations in Kubernetes resource templates [#359](https://github.com/qdrant/qdrant-helm/pull/359)

## [qdrant-1.15.0](https://github.com/qdrant/qdrant-helm/tree/qdrant-1.15.0) (2025-07-18)

- Update Qdrant to v1.15.0
- Use correct value in readOnlyApiKey check [#340](https://github.com/qdrant/qdrant-helm/pull/340)
- Ensure that open file descriptor limit is set correctly on startup [#355](https://github.com/qdrant/qdrant-helm/pull/355)
- Fix fsGroup and runAsUser not being int64 type [#354](https://github.com/qdrant/qdrant-helm/pull/354), [#357](https://github.com/qdrant/qdrant-helm/pull/357)

## [qdrant-1.14.1](https://github.com/qdrant/qdrant-helm/tree/qdrant-1.14.1) (2025-05-23)

- Update Qdrant to v1.14.1
- Fix typo in README [#323](https://github.com/qdrant/qdrant-helm/pull/323)

## [qdrant-1.14.0](https://github.com/qdrant/qdrant-helm/tree/qdrant-1.14.0) (2025-04-22)

- Update Qdrant to v1.14.0
- Make chart test image configurable [#320](https://github.com/qdrant/qdrant-helm/pull/320)
- Do not unnecessarily update file system ownership when unprivilged image is not used [#321](https://github.com/qdrant/qdrant-helm/pull/321) 

## [qdrant-1.13.6](https://github.com/qdrant/qdrant-helm/tree/qdrant-1.13.6) (2025-04-01)

- Update Qdrant to v1.13.6

## [qdrant-1.13.5](https://github.com/qdrant/qdrant-helm/tree/qdrant-1.13.5) (2025-03-21)

- Update Qdrant to v1.13.5
- Add support for P2P TLS [#314](https://github.com/qdrant/qdrant-helm/pull/314)
- Make Service appProtocol configurable [#313](https://github.com/qdrant/qdrant-helm/pull/313)
- Templating for config values [#308](https://github.com/qdrant/qdrant-helm/pull/308)
- Templating for topologySpreadConstraints [#309](https://github.com/qdrant/qdrant-helm/pull/309)
- Use tolerations and nodeSelector in test hook [#307](https://github.com/qdrant/qdrant-helm/pull/307)
- Set Service annotations on headless Service [#305](https://github.com/qdrant/qdrant-helm/pull/305)
- Support optional subPath for volumeMounts [#271](https://github.com/qdrant/qdrant-helm/pull/271)

## [qdrant-1.13.4](https://github.com/qdrant/qdrant-helm/tree/qdrant-1.13.4) (2025-02-17)

- Update Qdrant to v1.13.4

## [qdrant-1.13.3](https://github.com/qdrant/qdrant-helm/tree/qdrant-1.13.3) (2025-02-11)

- Update Qdrant to v1.13.3

## [qdrant-1.13.2](https://github.com/qdrant/qdrant-helm/tree/qdrant-1.13.2) (2025-01-28)

- Update Qdrant to v1.13.2

## [qdrant-1.13.1](https://github.com/qdrant/qdrant-helm/tree/qdrant-1.13.1) (2025-01-23)

- Update Qdrant to v1.13.1

## [qdrant-1.13.0](https://github.com/qdrant/qdrant-helm/tree/qdrant-1.13.0) (2025-01-16)

- Update Qdrant to v1.13.0

## [qdrant-1.12.6](https://github.com/qdrant/qdrant-helm/tree/qdrant-1.12.6) (2025-01-08)

- Update Qdrant to v1.12.6

## [qdrant-1.12.5](https://github.com/qdrant/qdrant-helm/tree/qdrant-1.12.5) (2024-12-09)

- Update Qdrant to v1.12.5

## [qdrant-1.12.4](https://github.com/qdrant/qdrant-helm/tree/qdrant-1.12.4) (2024-11-18)

- Update Qdrant to v1.12.4

## [qdrant-1.12.3](https://github.com/qdrant/qdrant-helm/tree/qdrant-1.12.3) (2024-11-11)

- Update Qdrant to v1.12.2

## [qdrant-1.12.2](https://github.com/qdrant/qdrant-helm/tree/qdrant-1.12.2) (2024-10-28)

- Fix release

## [qdrant-1.12.1](https://github.com/qdrant/qdrant-helm/tree/qdrant-1.12.1) (2024-10-14)

- Update Qdrant to v1.12.1

## [qdrant-1.12.0](https://github.com/qdrant/qdrant-helm/tree/qdrant-1.12.0) (2024-10-09)

- Update Qdrant to v1.12.0

## [qdrant-1.11.5](https://github.com/qdrant/qdrant-helm/tree/qdrant-1.11.5) (2024-09-23)

- Update Qdrant to v1.11.5

## [qdrant-1.11.4](https://github.com/qdrant/qdrant-helm/tree/qdrant-1.11.4) (2024-09-18)

- Update Qdrant to v1.11.4
- Prefer read_only_api_key in ServiceMonitor [#221](https://github.com/qdrant/qdrant-helm/pull/221)
- Added support for reading apiKey and readOnlyApiKey from external secrets [#230](https://github.com/qdrant/qdrant-helm/pull/230)

## [qdrant-1.11.3](https://github.com/qdrant/qdrant-helm/tree/qdrant-1.11.3) (2024-08-30)

- Update Qdrant to v1.11.3

## [qdrant-1.11.2](https://github.com/qdrant/qdrant-helm/tree/qdrant-1.11.2) (2024-08-28)

- Update Qdrant to v1.11.2

## [qdrant-1.11.1](https://github.com/qdrant/qdrant-helm/tree/qdrant-1.11.1) (2024-08-27)

- Update Qdrant to v1.11.1

## [qdrant-1.11.0](https://github.com/qdrant/qdrant-helm/tree/qdrant-1.11.0) (2024-08-12)

- Update Qdrant to v1.11.0
- Apply additional label to headless Service + ServiceMonitor to avoid duplicate scraping [#214](https://github.com/qdrant/qdrant-helm/pull/214)
- Apply tpl() to affinity values to enable reuse of helpers / labels [#213](https://github.com/qdrant/qdrant-helm/pull/213)

## [qdrant-0.10.1](https://github.com/qdrant/qdrant-helm/tree/qdrant-0.10.1) (2024-07-18)

- Update Qdrant to v1.10.1

## [qdrant-0.10.0](https://github.com/qdrant/qdrant-helm/tree/qdrant-0.10.0) (2024-07-02)

- Update Qdrant to v1.10.0

## [qdrant-0.9.4](https://github.com/qdrant/qdrant-helm/tree/qdrant-0.9.4) (2024-06-25)

- Update Qdrant to v1.9.7

## [qdrant-0.9.3](https://github.com/qdrant/qdrant-helm/tree/qdrant-0.9.3) (2024-06-22)

- Update Qdrant to v1.9.6

## [qdrant-0.9.2](https://github.com/qdrant/qdrant-helm/tree/qdrant-0.9.2) (2024-06-17)

- Update Qdrant to v1.9.5

## [qdrant-0.9.1](https://github.com/qdrant/qdrant-helm/tree/qdrant-0.9.1) (2024-06-10)

- Update Qdrant to v1.9.4

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
