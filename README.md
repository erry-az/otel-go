# Opentelemetry Go

OpenTelemetry is a collection of APIs, SDKs, and tools. Use it to instrument, generate, collect, and export telemetry data (metrics, logs, and traces) to help you analyze your software's performance and behavior.

This package just want to help implement SDKs easily with some default configuration.

- Source: https://opentelemetry.io/docs/specs/otel/protocol/exporter/

## How to implement

### Basic
On the go code, you can import this for setup simple and configured opentelemetry providers.

```go
package main

import "github.com/erry-az/otel-go"

// init open telemetry
otelProviders, err := otel.NewProviders(ctx)
if err != nil {
    return nil, err
}
```
- on the service host need to add the env, for example
```dotenv
# your service name
OTEL_SERVICE_NAME=grpc-service
# your opentelemetry collector endpoint type (stdout/grpc/http)
OTEL_EXPORTER_OTLP_TYPE=grpc
# your opentelemetry collector endpoint
OTEL_EXPORTER_OTLP_ENDPOINT=http://otel.service:4317
```
## Middleware / Instrumentation
- on otel go contrib
    - https://github.com/open-telemetry/opentelemetry-go-contrib/tree/main/instrumentation
- go fiber
    - https://docs.gofiber.io/contrib/otelfiber/

## Environment

Available environment variable for setup the OTLP.
OTLP stands for (opentelemetry provider)

### OTEL Service Name and Resource Attributes

| Environment Variable     | Description                                                              | Default Value | Available Values                  |
|--------------------------|--------------------------------------------------------------------------|---------------|-----------------------------------|
| OTEL_PROVIDERS           | Set provider to enable                                                   | trace,metric  | trace,metric,log                  |
| OTEL_SERVICE_NAME        | Set service name tag for all opentelemetry metric, traces, and log       | -             | -                                 |
| OTEL_RESOURCE_ATTRIBUTES | Set additional tag / label for all opentelemetry metric, traces, and log | -             | Format: `key1=value1,key2=value2` |

### OTLP Exporter Type

| Environment Variable            | Description                            | Default Value | Available Values            |
|---------------------------------|----------------------------------------|---------------|-----------------------------|
| OTEL_EXPORTER_OTLP_TYPE         | Set the global OTLP exporter type      | -             | stdout/grpc/http            |
| OTEL_EXPORTER_OTLP_TRACES_TYPE  | Set the OTLP exporter type for traces  | -             | stdout/grpc/http            |
| OTEL_EXPORTER_OTLP_METRICS_TYPE | Set the OTLP exporter type for metrics | -             | stdout/grpc/http/prometheus |
| OTEL_EXPORTER_OTLP_LOGS_TYPE    | Set the OTLP exporter type for logs    | -             | stdout/grpc/http            |

### OTLP Exporter Endpoint

| Environment Variable                | Description                                | Default Value   | Available Values |
|-------------------------------------|--------------------------------------------|-----------------|------------------|
| OTEL_EXPORTER_OTLP_ENDPOINT         | Set the global OTLP exporter endpoint      | Varies by type  | -                |
| OTEL_EXPORTER_OTLP_TRACES_ENDPOINT  | Set the OTLP exporter endpoint for traces  | -               | -                |
| OTEL_EXPORTER_OTLP_METRICS_ENDPOINT | Set the OTLP exporter endpoint for metrics | -               | -                |
| OTEL_EXPORTER_OTLP_LOGS_ENDPOINT    | Set the OTLP exporter endpoint for logs    | -               | -                |

### OTLP Exporter Endpoint Insecure

| Environment Variable                | Description                         | Default Value | Available Values |
|-------------------------------------|-------------------------------------|---------------|------------------|
| OTEL_EXPORTER_OTLP_INSECURE         | Set global insecure connection      | false         | true/false       |
| OTEL_EXPORTER_OTLP_TRACES_INSECURE  | Set insecure connection for traces  | false         | true/false       |
| OTEL_EXPORTER_OTLP_METRICS_INSECURE | Set insecure connection for metrics | false         | true/false       |
| OTEL_EXPORTER_OTLP_LOGS_INSECURE    | Set insecure connection for logs    | false         | true/false       |

### OTLP Exporter Endpoint Headers

| Environment Variable               | Description                                | Default Value | Available Values                   |
|------------------------------------|--------------------------------------------|---------------|------------------------------------|
| OTEL_EXPORTER_OTLP_HEADERS         | Set global headers for gRPC/HTTP requests  | -             | Format: `key1=value1,key2=value2`  |
| OTEL_EXPORTER_OTLP_TRACES_HEADERS  | Set headers for traces gRPC/HTTP requests  | -             | Format: `key1=value1,key2=value2`  |
| OTEL_EXPORTER_OTLP_METRICS_HEADERS | Set headers for metrics gRPC/HTTP requests | -             | Format: `key1=value1,key2=value2`  |
| OTEL_EXPORTER_OTLP_LOGS_HEADERS    | Set headers for logs gRPC/HTTP requests    | -             | Format: `key1=value1,key2=value2`  |

### OTLP Exporter Timeout

| Environment Variable               | Description                               | Default Value | Available Values |
|------------------------------------|-------------------------------------------|---------------|------------------|
| OTEL_EXPORTER_OTLP_TIMEOUT         | Set global timeout for batch export (ms)  | 10000         | -                |
| OTEL_EXPORTER_OTLP_TRACES_TIMEOUT  | Set timeout for traces batch export (ms)  | 10000         | -                |
| OTEL_EXPORTER_OTLP_METRICS_TIMEOUT | Set timeout for metrics batch export (ms) | 10000         | -                |
| OTEL_EXPORTER_OTLP_LOGS_TIMEOUT    | Set timeout for logs batch export (ms)    | 10000         | -                |

### OTLP Exporter Compression

| Environment Variable                   | Description                          | Default Value | Available Values |
|----------------------------------------|--------------------------------------|---------------|------------------|
| OTEL_EXPORTER_OTLP_COMPRESSION         | Set global gRPC/HTTP compressor      | -             | gzip             |
| OTEL_EXPORTER_OTLP_TRACES_COMPRESSION  | Set gRPC/HTTP compressor for traces  | -             | gzip             |
| OTEL_EXPORTER_OTLP_METRICS_COMPRESSION | Set gRPC/HTTP compressor for metrics | -             | gzip             |
| OTEL_EXPORTER_OTLP_LOGS_COMPRESSION    | Set gRPC/HTTP compressor for logs    | -             | gzip             |

### OTLP Exporter Secure Certificate Path

| Environment Variable                   | Description                                         | Default Value | Available Values |
|----------------------------------------|-----------------------------------------------------|---------------|------------------|
| OTEL_EXPORTER_OTLP_CERTIFICATE         | Set global filepath to the trusted certificate      | -             | -                |
| OTEL_EXPORTER_OTLP_TRACES_CERTIFICATE  | Set filepath to the trusted certificate for traces  | -             | -                |
| OTEL_EXPORTER_OTLP_METRICS_CERTIFICATE | Set filepath to the trusted certificate for metrics | -             | -                |
| OTEL_EXPORTER_OTLP_LOGS_CERTIFICATE    | Set filepath to the trusted certificate for logs    | -             | -                |

### OTLP Exporter Secure Client Certificate Path

| Environment Variable                          | Description                                             | Default Value | Available Values |
|-----------------------------------------------|---------------------------------------------------------|---------------|------------------|
| OTEL_EXPORTER_OTLP_CLIENT_CERTIFICATE         | Set global filepath to the client certificate for mTLS  | -             | -                |
| OTEL_EXPORTER_OTLP_TRACES_CLIENT_CERTIFICATE  | Set filepath to the client certificate for traces mTLS  | -             | -                |
| OTEL_EXPORTER_OTLP_METRICS_CLIENT_CERTIFICATE | Set filepath to the client certificate for metrics mTLS | -             | -                |
| OTEL_EXPORTER_OTLP_LOGS_CLIENT_CERTIFICATE    | Set filepath to the client certificate for logs mTLS    | -             | -                |

### OTLP Exporter Secure Client Key Path

| Environment Variable                  | Description                                               | Default Value | Available Values |
|---------------------------------------|-----------------------------------------------------------|---------------|------------------|
| OTEL_EXPORTER_OTLP_CLIENT_KEY         | Set global filepath to the client's private key for mTLS  | -             | -                |
| OTEL_EXPORTER_OTLP_TRACES_CLIENT_KEY  | Set filepath to the client's private key for traces mTLS  | -             | -                |
| OTEL_EXPORTER_OTLP_METRICS_CLIENT_KEY | Set filepath to the client's private key for metrics mTLS | -             | -                |
| OTEL_EXPORTER_OTLP_LOGS_CLIENT_KEY    | Set filepath to the client's private key for logs mTLS    | -             | -                |

### OTLP Exporter Metrics Temporality Preference (metrics only)

| Environment Variable                               | Description                                         | Default Value | Available Values            |
|----------------------------------------------------|-----------------------------------------------------|---------------|-----------------------------|
| OTEL_EXPORTER_OTLP_METRICS_TEMPORALITY_PREFERENCE  | Set aggregation temporality preference for metrics  | cumulative    | cumulative/delta/lowmemory  |

### OTLP Exporter Metrics Default Histogram Aggregation (metrics only)

| Environment Variable                                      | Description                                        | Default Value              | Available Values                                              |
|-----------------------------------------------------------|----------------------------------------------------|----------------------------|---------------------------------------------------------------|
| OTEL_EXPORTER_OTLP_METRICS_DEFAULT_HISTOGRAM_AGGREGATION  | Set default aggregation for histogram instruments  | explicit_bucket_histogram  | explicit_bucket_histogram/base2_exponential_bucket_histogram  |

Note: The configuration can be overridden by various `With...` options such as `WithEndpoint`, `WithEndpointURL`, `WithInsecure`, `WithHeaders`, `WithTimeout`, `WithCompressor`, `WithTLSCredentials`, and (grpc only: `WithGRPCConn`).