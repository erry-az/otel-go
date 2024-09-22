package otel

import "errors"

// LogExporterType endpoint type for OTLP log exporter
type LogExporterType string

const (
	// GrpcLogExporter exporter grpc type
	GrpcLogExporter LogExporterType = "grpc"
	// HttpLogExporter exporter http type
	HttpLogExporter LogExporterType = "http"
	// StdOutLogExporter exporter stdout type
	StdOutLogExporter LogExporterType = "stdout"
)

// ErrInvalidLogExporterType invalid log exporter type error
var ErrInvalidLogExporterType = errors.New("invalid log exporter type")
