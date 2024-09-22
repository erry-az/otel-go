package otel

import "errors"

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
