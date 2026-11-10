# Changelog

## 0.1.2

### Added

- Namespace-scoped installs. The manager reads `WATCH_NAMESPACE`, a
  comma-separated list, and watches only those namespaces, so it needs RBAC
  only there. Unset keeps the cluster-wide watch.
- Chart values `manager.watchNamespaces` (sets `WATCH_NAMESPACE`) and
  `rbac.create` (set `false` when the platform grants the manager's RBAC to an
  existing ServiceAccount, as ACE strict RBAC does). With the defaults the
  chart renders exactly what 0.1.1 did.

## 0.1.1

### Fixed

- Pool `maxQueryMemLimit` and `minQueryMemLimit` are now written to
  llama-site.xml as byte counts. Impala only accepts plain numbers there, so
  values such as `4gb` used to break the pool. The CRD validates the syntax.

### Added

- CI publishes a multi-arch operator image to
  `ghcr.io/rozsapeter96/impala-operator` and, on release tags, the Helm chart
  to `oci://ghcr.io/rozsapeter96/charts/impala-operator`. The chart's default
  image repository points at the published image.
- Unit coverage for unauthenticated REST catalogs (no `security` key) and
  static `header.*` properties.

### Changed

- The e2e suite uses RustFS instead of the discontinued MinIO images.

### Documentation

- REST catalog authentication, static headers, and why catalog base locations
  must use `s3a://` (Impala cannot scan or write `s3://` locations).

## 0.1.0

- Initial release: ImpalaCluster CRD, coordinator and executor-group
  topology, admission-control pools, executor-group autoscaling, TLS,
  Kerberos and LDAP, Hive Metastore and S3 storage, Iceberg REST catalogs
  (Apache Polaris) and catalogd-less mode.
