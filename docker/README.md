## Content

- [collector](./collector) directory contains OpenTelemetry Collector configuration file.
- [grafana](./grafana) directory contains Grafana configuration file.
- [loki](./loki) directory contains Loki configuration file.
- [prometheus](./prometheus) directory contains Prometheus configuration file.

## Try

1. Run `docker-compose up`
1. Open Jaeger at http://localhost:16686
1. Open Prometheus at http://localhost:9090
1. Open Grafana at http://localhost:3000 (login with admin/admin)
    - Application dashboard: http://localhost:3000/d/application-monitoring/application-monitoring
    - System dashboard: http://localhost:3000/d/system-monitoring/system-monitoring
