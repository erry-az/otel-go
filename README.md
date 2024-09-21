# Opentelemetry Go

OpenTelemetry is a collection of APIs, SDKs, and tools. Use it to instrument, generate, collect, and export telemetry data (metrics, logs, and traces) to help you analyze your software’s performance and behavior.

This package just want to help implement SDKs easily with some default configuration.

- Source: https://opentelemetry.io/docs/specs/otel/protocol/exporter/

## How to implement 

### Basic
On the go code, you can import this for setup simple and configured opentelemetry providers.

```go
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

### OTEL Service Name
Set service name tag for all opentelemetry metric and traces.

```dotenv
OTEL_SERVICE_NAME=
```

### OTEL Resource Attributes
Set additional tag / lable for all opentelemetry metric and traces. With value `key1=value1,key2=value2`
```dotenv
OTEL_RESOURCE_ATTRIBUTES=
```

### OTLP Exporter Type

Set the OTLP expoter type for define what type of open telemetry collector target endpoint.
There are 3 ways of doing it. Set global, traces, and metrics exporter type.
- default value : `stdout`
- available values : `stdout`/`grpc`/`http` (metrics only: `prometheus`) 

```dotenv
# global
OTEL_EXPORTER_OTLP_TYPE=

# traces only
OTEL_EXPORTER_OTLP_TRACES_TYPE=

# metrics only 
OTEL_EXPORTER_OTLP_METRICS_TYPE=
```

### OTLP Exporter Endpoint

Set the OTLP exporter endpoint based on exporter type. 
The configuration can be overridden by `WithEndpoint`, `WithEndpointURL`, `WithInsecure`, and (grpc only : `WithGRPCConn`).
- default value : (stdout: "", grpc: `http://localhost:4317`, grpc: `http://localhost:4318`)

```dotenv
# global
OTEL_EXPORTER_OTLP_ENDPOINT=

# traces only
OTEL_EXPORTER_OTLP_TRACES_ENDPOINT=

# metrics only
OTEL_EXPORTER_OTLP_METRICS_ENDPOINT=
```

### OTLP Exporter Endpoint Insecure

Set the OTLP exporter endpoint is using insecure connection or not.
The configuration can be overridden by `WithInsecure`, (grpc only : `WithGRPCConn`) options.
- default value : "false"
- available values : "true" and "false"

```dotenv
# global
OTEL_EXPORTER_OTLP_INSECURE=
# traces only
OTEL_EXPORTER_OTLP_TRACES_INSECURE=
# metrics only
OTEL_EXPORTER_OTLP_METRICS_INSECURE=
```
### OTLP Exporter Endpoint Headers

Set key-value pairs used as gRPC metadata/HTTP Headers associated with gRPC/HTTP requests. 
you can fill with format "key1=value1,key2=value2". The configuration can be overridden by `WithHeaders` option.
- default value : None

```dotenv
# global
OTEL_EXPORTER_OTLP_HEADERS=
# traces only
OTEL_EXPORTER_OTLP_TRACES_HEADERS=
# metrics only
OTEL_EXPORTER_OTLP_METRICS_HEADERS=
```
 
### OTLP Exporter Timeout

Maximum time in milliseconds the OTLP exporter waits for each batch export.
The configuration can be overridden by `WithTimeout` option.
- default value : 10000

```dotenv
# global
OTEL_EXPORTER_OTLP_TIMEOUT=
# traces only
OTEL_EXPORTER_OTLP_TRACES_TIMEOUT=
# metrics only
OTEL_EXPORTER_OTLP_METRICS_TIMEOUT=
```

### OTLP Exporter Compression

gRPC/HTTP compressor the exporter uses.
The configuration can be overridden by `WithCompressor`,  (grpc only : `WithGRPCConn`) options.

- default value : None
- available values : "gzip"

```dotenv
# global
OTEL_EXPORTER_OTLP_COMPRESSION=
# traces only
OTEL_EXPORTER_OTLP_TRACES_COMPRESSION=
# metrics only
OTEL_EXPORTER_OTLP_METRICS_COMPRESSION=
```

### OTLP Exporter Secure Certificate Path
The filepath to the trusted certificate to use when verifying a server's TLS credentials.
The configuration can be overridden by `WithTLSCredentials`,  (grpc only : `WithGRPCConn`) options.

- default value : None

```dotenv
# global
OTEL_EXPORTER_OTLP_CERTIFICATE=
# traces only
OTEL_EXPORTER_OTLP_TRACES_CERTIFICATE=
# metrics only
OTEL_EXPORTER_OTLP_METRICS_CERTIFICATE=
```

### OTLP Exporter Secure Client Certificate Path
The filepath to the client certificate/chain trust for client's private key to use in mTLS communication in PEM format.
The configuration can be overridden by `WithTLSCredentials`, `WithGRPCConn` options.

- default value : None

```dotenv
# global
OTEL_EXPORTER_OTLP_CLIENT_CERTIFICATE=
# traces only
OTEL_EXPORTER_OTLP_TRACES_CLIENT_CERTIFICATE=
# metrics only
OTEL_EXPORTER_OTLP_METRICS_CLIENT_CERTIFICATE=
```

### OTLP Exporter Secure Client Key Path
The filepath to the client's private key to use in mTLS communication in PEM format.
The configuration can be overridden by `WithTLSCredentials`, `WithGRPCConn` option.

- default value : None

```dotenv
# global
OTEL_EXPORTER_OTLP_CLIENT_KEY=
# traces only
OTEL_EXPORTER_OTLP_TRACES_CLIENT_KEY=
# metrics only
OTEL_EXPORTER_OTLP_METRICS_CLIENT_KEY=
```

### OTLP Exporter Metrics Temporality Preference (metrics only)

Aggregation temporality to use on the basis of instrument kind. 
The configuration can be overridden by `WithTemporalitySelector` option.

- default value : cumulative
- Supported values:
  - "cumulative" - Cumulative aggregation temporality for all instrument kinds,
  - "delta" - Delta aggregation temporality for Counter, Asynchronous Counter and Histogram instrument kinds; Cumulative aggregation for UpDownCounter and Asynchronous UpDownCounter instrument kinds,
  - "lowmemory" - Delta aggregation temporality for Synchronous Counter and Histogram instrument kinds; Cumulative aggregation temporality for Synchronous UpDownCounter, Asynchronous Counter, and Asynchronous UpDownCounter instrument kinds.

```dotenv
OTEL_EXPORTER_OTLP_METRICS_TEMPORALITY_PREFERENCE=
```

### OTLP Exporter Metrics Default Histogram Aggregation (metrics only)

Default aggregation to use for histogram instruments. 
- default value : explicit_bucket_histogram
- Supported values:
  - "explicit_bucket_histogram" - Explicit Bucket Histogram Aggregation,
  - "base2_exponential_bucket_histogram" - Base2 Exponential Bucket Histogram Aggregation.

```dotenv
OTEL_EXPORTER_OTLP_METRICS_DEFAULT_HISTOGRAM_AGGREGATION=
```