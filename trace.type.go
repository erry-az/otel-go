package otel

import "errors"

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

// ErrInvalidTraceExporterType invalid trace exporter type error
var ErrInvalidTraceExporterType = errors.New("invalid trace exporter type")
