package otel

import (
	"os"
	"strings"
)

// environment for exporter type
const (
	exporterTypeEnv       = "OTEL_EXPORTER_OTLP_TYPE"
	traceExporterTypeEnv  = "OTEL_EXPORTER_OTLP_TRACES_TYPE"
	metricExporterTypeEnv = "OTEL_EXPORTER_OTLP_METRICS_TYPE"
	logExporterTypeEnv    = "OTEL_EXPORTER_OTLP_LOGS_TYPE"
	providersEnv          = "OTEL_PROVIDERS"
)

// default env
var (
	providersEnvDefault = ProvidersEnable{Trace: true, Metric: true}
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

func getLogExporterTypeFromEnv() LogExporterType {
	var (
		envExporterType    = os.Getenv(exporterTypeEnv)
		envLogExporterType = os.Getenv(logExporterTypeEnv)
	)

	if envExporterType != "" {
		return LogExporterType(envExporterType)
	}

	if envLogExporterType != "" {
		return LogExporterType(envLogExporterType)
	}

	return ""
}

func getProvidersEnable() (ProvidersEnable, error) {
	var providers ProvidersEnable

	envProviders := os.Getenv(providersEnv)
	if envProviders == "" {
		return providersEnvDefault, nil
	}

	parsedProviders := strings.Split(envProviders, ",")
	for _, providerRaw := range parsedProviders {
		provider, err := validateProviderType(providerRaw)
		if err != nil {
			return ProvidersEnable{}, err
		}

		if provider == providerTypeTrace {
			providers.Trace = true
		}

		if provider == providerTypeMetric {
			providers.Metric = true
		}

		if provider == providerTypeLog {
			providers.Log = true
		}
	}

	return providers, nil
}
