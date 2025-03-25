#!/bin/bash

export OTEL_SERVICE_NAME="learning-otel-typescript"
export OTEL_METRICS_EXPORTER="otlp"
export OTEL_TRACES_EXPORTER="otlp"
export OTEL_LOGS_EXPORTER="otlp"
export OTEL_EXPORTER_OTLP_PROTOCOL="grpc"
export OTEL_EXPORTER_OTLP_ENDPOINT="http://localhost:4317"
export OTEL_NODE_RESOURCE_DETECTORS="all"
export NODE_OPTIONS="--require @opentelemetry/auto-instrumentations-node/register"

nest start --watch
