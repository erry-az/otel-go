package otel

import (
	"context"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutlog"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/resource"
)

// LogExporterOption option for log exporter
type LogExporterOption struct {
	GrpcOpts []otlploggrpc.Option
	HttpOpts []otlploghttp.Option
}

// NewLogExporter new log exporter with defined type
// for grpc and http there is env will be provided
// OTEL_EXPORTER_OTLP_ENDPOINT, OTEL_EXPORTER_OTLP_LOGS_ENDPOINT = (default: grpc : "https://localhost:4317", http :"https://localhost:4318" ("/v1/logs" is appended))
// The configuration can be overridden by  WithEndpoint, WithEndpointURL, WithInsecure, and (grpc only : WithGRPCConn)
// OTEL_EXPORTER_OTLP_INSECURE, OTEL_EXPORTER_OTLP_LOGS_INSECURE = (default: "false")
// The configuration can be overridden by WithInsecure, WithGRPCConn options.
// OTEL_EXPORTER_OTLP_HEADERS, OTEL_EXPORTER_OTLP_LOGS_HEADERS = (default: none)
// key-value pairs used as gRPC metadata associated with gRPC requests. you can fill with format "key1=value1,key2=value2"
// The configuration can be overridden by WithHeaders option.
// OTEL_EXPORTER_OTLP_TIMEOUT, OTEL_EXPORTER_OTLP_LOGS_TIMEOUT = (default: "10000")
// maximum time in milliseconds the OTLP exporter waits for each batch export.
// The configuration can be overridden by WithTimeout option.
// OTEL_EXPORTER_OTLP_COMPRESSION, OTEL_EXPORTER_OTLP_LOGS_COMPRESSION = (default: none) supported value "gzip"
// The configuration can be overridden by WithCompressor, WithGRPCConn options
// OTEL_EXPORTER_OTLP_CERTIFICATE, OTEL_EXPORTER_OTLP_LOGS_CERTIFICATE = (default: none)
// the filepath to the trusted certificate to use when verifying a server's TLS credentials.
// The configuration can be overridden by WithTLSCredentials, WithGRPCConn options.
// OTEL_EXPORTER_OTLP_CLIENT_CERTIFICATE, OTEL_EXPORTER_OTLP_LOGS_CLIENT_CERTIFICATE = (default: none)
// the filepath to the client certificate/chain trust for client's private key to use in mTLS communication in PEM format.
// The configuration can be overridden by WithTLSCredentials, WithGRPCConn options.
// OTEL_EXPORTER_OTLP_CLIENT_KEY, OTEL_EXPORTER_OTLP_LOGS_CLIENT_KEY = (default: none)
// the filepath to the client's private key to use in mTLS communication in PEM format.
// The configuration can be overridden by WithTLSCredentials, WithGRPCConn option.
//
// stdout just will print out the log
func NewLogExporter(ctx context.Context, endpointType LogExporterType, opt LogExporterOption) (sdklog.Exporter, error) {
	switch endpointType {
	case HttpLogExporter:
		return otlploghttp.New(ctx, opt.HttpOpts...)
	case GrpcLogExporter:
		return otlploggrpc.New(ctx, opt.GrpcOpts...)
	case StdOutLogExporter:
		return stdoutlog.New(stdoutlog.WithPrettyPrint())
	}

	return nil, ErrInvalidLogExporterType
}

// NewLogProvider initiate provider for log
func NewLogProvider(res *resource.Resource, exporter sdklog.Exporter, opts ...sdklog.LoggerProviderOption) (*sdklog.LoggerProvider, error) {
	return sdklog.NewLoggerProvider(
		append([]sdklog.LoggerProviderOption{
			sdklog.WithResource(res),
			sdklog.WithProcessor(sdklog.NewBatchProcessor(exporter)),
		}, opts...)...,
	), nil
}

// InitLogProvider using basic init log without option
// this will do init log exporter by exporterType argument
// pass the exporter to log provider
// set new log provider to global
// and set global context propagation using log context and baggage as propagator
func InitLogProvider(ctx context.Context, res *resource.Resource) (*sdklog.LoggerProvider, error) {
	exporterType := getLogExporterTypeFromEnv()

	if exporterType == "" {
		return nil, nil
	}

	exporter, err := NewLogExporter(ctx, exporterType, LogExporterOption{})
	if err != nil {
		return nil, err
	}

	logProvider, err := NewLogProvider(res, exporter)
	if err != nil {
		return nil, err
	}

	return logProvider, nil
}
