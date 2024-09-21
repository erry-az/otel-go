package otel

import (
	"context"
	"errors"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

// TraceExporterType endpoint type for OTLP exporter
type TraceExporterType string

const (
	// GrpcTraceExporter exporter grpc type
	GrpcTraceExporter TraceExporterType = "grpc"
	// HttpTraceExporter exporter http type
	HttpTraceExporter TraceExporterType = "http"
	// StdOutTraceExporter exporter stdout type
	StdOutTraceExporter TraceExporterType = "stdout"
)

// TraceExporterOption option for trace exporter
type TraceExporterOption struct {
	GrpcOpts []otlptracegrpc.Option
	HttpOpts []otlptracehttp.Option
}

// ErrInvalidTraceExporterType invalid metric exporter type error
var ErrInvalidTraceExporterType = errors.New("invalid metric exporter type")

// NewTraceExporter new trace exporter with defined type
// for grpc and http there is env will be provided
// OTEL_EXPORTER_OTLP_ENDPOINT, OTEL_EXPORTER_OTLP_TRACES_ENDPOINT = (default: grpc : "https://localhost:4317", http :"https://localhost:4318" ("/v1/traces" is appended))
// The configuration can be overridden by  WithEndpoint, WithEndpointURL, WithInsecure, and (grpc only : WithGRPCConn)
// OTEL_EXPORTER_OTLP_INSECURE, OTEL_EXPORTER_OTLP_TRACES_INSECURE = (default: "false")
// The configuration can be overridden by WithInsecure, WithGRPCConn options.
// OTEL_EXPORTER_OTLP_HEADERS, OTEL_EXPORTER_OTLP_TRACES_HEADERS = (default: none)
// key-value pairs used as gRPC metadata associated with gRPC requests. you can fill with format "key1=value1,key2=value2"
// The configuration can be overridden by WithHeaders option.
// OTEL_EXPORTER_OTLP_TIMEOUT, OTEL_EXPORTER_OTLP_TRACES_TIMEOUT = (default: "10000")
// maximum time in milliseconds the OTLP exporter waits for each batch export.
// The configuration can be overridden by WithTimeout option.
// OTEL_EXPORTER_OTLP_COMPRESSION, OTEL_EXPORTER_OTLP_TRACES_COMPRESSION = (default: none) supported value "gzip"
// The configuration can be overridden by WithCompressor, WithGRPCConn options
// OTEL_EXPORTER_OTLP_CERTIFICATE, OTEL_EXPORTER_OTLP_TRACES_CERTIFICATE = (default: none)
// the filepath to the trusted certificate to use when verifying a server's TLS credentials.
// The configuration can be overridden by WithTLSCredentials, WithGRPCConn options.
// OTEL_EXPORTER_OTLP_CLIENT_CERTIFICATE, OTEL_EXPORTER_OTLP_TRACES_CLIENT_CERTIFICATE = (default: none)
// the filepath to the client certificate/chain trust for client's private key to use in mTLS communication in PEM format.
// The configuration can be overridden by WithTLSCredentials, WithGRPCConn options.
// OTEL_EXPORTER_OTLP_CLIENT_KEY, OTEL_EXPORTER_OTLP_TRACES_CLIENT_KEY = (default: none)
// the filepath to the client's private key to use in mTLS communication in PEM format.
// The configuration can be overridden by WithTLSCredentials, WithGRPCConn option.
//
// stdout just will print out the trace
func NewTraceExporter(ctx context.Context, endpointType TraceExporterType, opt TraceExporterOption) (sdktrace.SpanExporter, error) {
	switch endpointType {
	case HttpTraceExporter:
		return otlptracehttp.New(ctx, opt.HttpOpts...)
	case GrpcTraceExporter:
		return otlptracegrpc.New(ctx, opt.GrpcOpts...)
	case StdOutTraceExporter:
		return stdouttrace.New(stdouttrace.WithPrettyPrint())
	}

	return nil, ErrInvalidTraceExporterType
}

// NewTraceProvider initiate provider for trace
func NewTraceProvider(res *resource.Resource, exporter sdktrace.SpanExporter, opts ...sdktrace.TracerProviderOption) (*sdktrace.TracerProvider, error) {
	return sdktrace.NewTracerProvider(
		append([]sdktrace.TracerProviderOption{
			sdktrace.WithSampler(sdktrace.AlwaysSample()),
			sdktrace.WithResource(res),
			sdktrace.WithSpanProcessor(sdktrace.NewBatchSpanProcessor(exporter)),
		}, opts...)...,
	), nil
}

// SetGlobalTraceProvider set trace provider as global trace provider
func SetGlobalTraceProvider(traceProvider *sdktrace.TracerProvider) {
	otel.SetTracerProvider(traceProvider)
}

// SetGlobalContextPropagation set global context propagation setting with trace context and baggage as default
// and can be added more propagators
func SetGlobalContextPropagation(propagators ...propagation.TextMapPropagator) {
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(append(
			[]propagation.TextMapPropagator{propagation.TraceContext{}, propagation.Baggage{}},
			propagators...,
		)...),
	)
}

// InitTraceProvider using basic init trace without option
// this will do init trace exporter by exporterType argument
// pass the exporter to trace provider
// set new trace provider to global
// and set global context propagation using trace context and baggage as propagator
func InitTraceProvider(ctx context.Context, res *resource.Resource) (*sdktrace.TracerProvider, error) {
	exporterType := getTraceExporterTypeFromEnv()

	if exporterType == "" {
		return nil, nil
	}

	exporter, err := NewTraceExporter(ctx, exporterType, TraceExporterOption{})
	if err != nil {
		return nil, err
	}

	traceProvider, err := NewTraceProvider(res, exporter)
	if err != nil {
		return nil, err
	}

	return traceProvider, nil
}
