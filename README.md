# Learning OpenTelemetry

This repository contains a collection of applications instrumented with OpenTelemetry. 
The goal is to provide examples of how to use OpenTelemetry in different programming languages and frameworks. 

The applications are instrumented with OpenTelemetry SDKs and are configured to send telemetry data to a OpenTelemetry Collector.

The OpenTelemetry Collector is configured to send the telemetry data to:
- Jaeger and Tempo backend for tracing
- Prometheus for metrics
- Loki for logs

Grafana is used to visualize the telemetry data.

## Content

- [docker](./docker/) directory contains docker-compose and service configuration files.
- [leaning-otel-java](./leaning-otel-java/) directory contains a Java application instrumented with OpenTelemetry Spring Starter.
- [LearningOtelDotnet](./LearningOtelDotnet/) directory contains a .NET application instrumented with OpenTelemetry .NET SDK.
- [learning-otel-php](./learning-otel-php/) directory contains a PHP + Laravel application instrumented with OpenTelemetry PHP SDK.
- [learning-otel-typescript](./learning-otel-typescript/) directory contains a Node + NestJS + TypeScript application instrumented with OpenTelemetry JavaScript SDK.

## Tools

- [OpenTelemetry](https://opentelemetry.io/) - OpenTelemetry is a set of APIs, libraries, agents, and instrumentation to provide observability for applications.
- [Jaeger](https://www.jaegertracing.io/) - Jaeger is an open-source, end-to-end distributed tracing system.
- [Tempo](https://grafana.com/oss/tempo/) - Tempo is a high-scale, cost-effective, and easy-to-use distributed tracing backend.
- [Prometheus](https://prometheus.io/) - Prometheus is an open-source systems monitoring and alerting toolkit.
- [Loki](https://grafana.com/oss/loki/) - Loki is a log aggregation system designed to work with Grafana.
- [Grafana](https://grafana.com/) - Grafana is an open-source platform for monitoring and observability.
