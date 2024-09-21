package otel

import (
	"context"
	"errors"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
)

// MetricExporterType endpoint type for OTLP exporter
type MetricExporterType string

const (
	// GrpcMetricExporter exporter grpc type
	GrpcMetricExporter MetricExporterType = "grpc"
	// HttpMetricExporter exporter http type
	HttpMetricExporter MetricExporterType = "http"
	// PrometheusMetricExporter exporter prometheus type
	PrometheusMetricExporter MetricExporterType = "prometheus"
	// StdOutMetricExporter exporter prometheus type
	StdOutMetricExporter MetricExporterType = "stdout"
)

// ErrInvalidMetricExporterType invalid metric exporter type error
var ErrInvalidMetricExporterType = errors.New("invalid metric exporter type")

// MetricExporterOption option for metric exporter
type MetricExporterOption struct {
	GrpcOpts   []otlpmetricgrpc.Option
	HttpOpts   []otlpmetrichttp.Option
	ReaderOpts []sdkmetric.PeriodicReaderOption
}

// NewMetricsExporter new metrics exporter with defined type
// for grpc and http there is env will be provided
// OTEL_EXPORTER_OTLP_ENDPOINT, OTEL_EXPORTER_OTLP_METRICS_ENDPOINT = (default: grpc : "http://localhost:4317", http :"http://localhost:4318" ("/v1/traces" is appended))
// The configuration can be overridden by  WithEndpoint, WithEndpointURL, WithInsecure, and (grpc only : WithGRPCConn)
// OTEL_EXPORTER_OTLP_INSECURE, OTEL_EXPORTER_OTLP_METRICS_INSECURE = (default: "false")
// The configuration can be overridden by WithInsecure, WithGRPCConn options.
// OTEL_EXPORTER_OTLP_HEADERS, OTEL_EXPORTER_OTLP_METRICS_HEADERS = (default: none)
// key-value pairs used as gRPC metadata associated with gRPC requests. you can fill with format "key1=value1,key2=value2"
// The configuration can be overridden by WithHeaders option.
// OTEL_EXPORTER_OTLP_TIMEOUT, OTEL_EXPORTER_OTLP_METRICS_TIMEOUT = (default: "10000")
// maximum time in milliseconds the OTLP exporter waits for each batch export.
// The configuration can be overridden by WithTimeout option.
// OTEL_EXPORTER_OTLP_COMPRESSION, OTEL_EXPORTER_OTLP_METRICS_COMPRESSION = (default: none) supported value "gzip"
// The configuration can be overridden by WithCompressor, WithGRPCConn options
// OTEL_EXPORTER_OTLP_CERTIFICATE, OTEL_EXPORTER_OTLP_METRICS_CERTIFICATE = (default: none)
// the filepath to the trusted certificate to use when verifying a server's TLS credentials.
// The configuration can be overridden by WithTLSCredentials, WithGRPCConn options.
// OTEL_EXPORTER_OTLP_CLIENT_CERTIFICATE, OTEL_EXPORTER_OTLP_METRICS_CLIENT_CERTIFICATE = (default: none)
// the filepath to the client certificate/chain trust for client's private key to use in mTLS communication in PEM format.
// The configuration can be overridden by WithTLSCredentials, WithGRPCConn options.
// OTEL_EXPORTER_OTLP_CLIENT_KEY, OTEL_EXPORTER_OTLP_METRICS_CLIENT_KEY = (default: none)
// the filepath to the client's private key to use in mTLS communication in PEM format.
// The configuration can be overridden by WithTLSCredentials, WithGRPCConn option.
// OTEL_EXPORTER_OTLP_METRICS_TEMPORALITY_PREFERENCE = (default: "cumulative")
// Supported values:
// - "cumulative" - Cumulative aggregation temporality for all instrument kinds,
// - "delta" - Delta aggregation temporality for Counter, Asynchronous Counter and Histogram instrument kinds; Cumulative aggregation for UpDownCounter and Asynchronous UpDownCounter instrument kinds,
// - "lowmemory" - Delta aggregation temporality for Synchronous Counter and Histogram instrument kinds; Cumulative aggregation temporality for Synchronous UpDownCounter, Asynchronous Counter, and Asynchronous UpDownCounter instrument kinds.
// OTEL_EXPORTER_OTLP_METRICS_DEFAULT_HISTOGRAM_AGGREGATION = (default: "explicit_bucket_histogram")
// Supported values:
// - "explicit_bucket_histogram" - Explicit Bucket Histogram Aggregation https://github.com/open-telemetry/opentelemetry-specification/blob/v1.26.0/specification/metrics/sdk.md#explicit-bucket-histogram-aggregation,
// - "base2_exponential_bucket_histogram" - Base2 Exponential Bucket Histogram Aggregation https://github.com/open-telemetry/opentelemetry-specification/blob/v1.26.0/specification/metrics/sdk.md#base2-exponential-bucket-histogram-aggregation.
//
// stdout just will print out the trace
//
// prometheus using prometheus
func NewMetricsExporter(ctx context.Context, endpointType MetricExporterType, opts MetricExporterOption) (sdkmetric.Reader, error) {
	var (
		exporter sdkmetric.Exporter
		err      error
	)

	switch endpointType {
	case HttpMetricExporter:
		exporter, err = otlpmetrichttp.New(ctx, opts.HttpOpts...)
	case GrpcMetricExporter:
		exporter, err = otlpmetricgrpc.New(ctx, opts.GrpcOpts...)
	case StdOutMetricExporter:
		exporter, err = stdoutmetric.New(stdoutmetric.WithPrettyPrint())
	case PrometheusMetricExporter:
		return prometheus.New()
	default:
		return nil, ErrInvalidMetricExporterType
	}

	if err != nil {
		return nil, err
	}

	return sdkmetric.NewPeriodicReader(exporter, opts.ReaderOpts...), nil
}

// NewMetricProvider initiate provider for metric
func NewMetricProvider(res *resource.Resource, reader sdkmetric.Reader, opts ...sdkmetric.Option) (*sdkmetric.MeterProvider, error) {
	return sdkmetric.NewMeterProvider(append([]sdkmetric.Option{
		sdkmetric.WithResource(res),
		sdkmetric.WithReader(reader),
	}, opts...)...), nil
}

// SetGlobalMetricProvider set metric provider as global meter provider
func SetGlobalMetricProvider(metricProvider *sdkmetric.MeterProvider) {
	otel.SetMeterProvider(metricProvider)
}

// InitMetricProvider using basic init metric provider without option
// this will do init metric exporter by exporterType argument
// pass the exporter to metric provider
// set new metric provider to global
func InitMetricProvider(ctx context.Context, res *resource.Resource) (*sdkmetric.MeterProvider, error) {
	exporterType := getMetricExporterTypeFromEnv()

	if exporterType == "" {
		return nil, nil
	}

	exporter, err := NewMetricsExporter(ctx, exporterType, MetricExporterOption{})
	if err != nil {
		return nil, err
	}

	provider, err := NewMetricProvider(res, exporter)
	if err != nil {
		return nil, err
	}

	return provider, nil
}
