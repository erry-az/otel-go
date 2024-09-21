package otel

import "os"

// environment for exporter type
const (
	exporterTypeEnv       = "OTEL_EXPORTER_OTLP_TYPE"
	traceExporterTypeEnv  = "OTEL_EXPORTER_OTLP_TRACES_TYPE"
	metricExporterTypeEnv = "OTEL_EXPORTER_OTLP_METRICS_TYPE"
)

func getTraceExporterTypeFromEnv() TraceExporterType {
	var (
		envExporterType      = os.Getenv(exporterTypeEnv)
		envTraceExporterType = os.Getenv(traceExporterTypeEnv)
	)

	if envExporterType != "" {
		return TraceExporterType(envExporterType)
	}

	if envTraceExporterType != "" {
		return TraceExporterType(envTraceExporterType)
	}

	return ""
}

func getMetricExporterTypeFromEnv() MetricExporterType {
	var (
		envExporterType       = os.Getenv(exporterTypeEnv)
		envMetricExporterType = os.Getenv(metricExporterTypeEnv)
	)

	if envExporterType != "" {
		return MetricExporterType(envExporterType)
	}

	if envMetricExporterType != "" {
		return MetricExporterType(envMetricExporterType)
	}

	return ""
}
