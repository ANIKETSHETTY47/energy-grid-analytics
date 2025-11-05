# energy-grid-analytics-go

A small Go module for energy analytics: unit conversion, anomaly detection, data aggregation, and maintenance prediction.

## Install

```bash
go get github.com/yourusername/energy-grid-analytics-go@v1.0.0
```

## Packages

- `converter`: Unit conversions, cost and efficiency calculations.
- `anomaly`: Sliding-window spike detection and IQR-based outlier detection.
- `aggregator`: Rolling and grouped aggregations for energy data.
- `maintenance`: Simple failure-risk and next-service predictors.

## Quick Start

See [`examples/usage.go`](examples/usage.go) for a runnable example:

```bash
cd examples
go run .
```

## Versioning

This repo follows semantic versioning. Tag a release to publish as a module:

```bash
git tag v1.0.0
git push origin v1.0.0
```
