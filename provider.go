package otel

import (
	"context"

	sdklog "go.opentelemetry.io/otel/sdk/log"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

// Providers open telemetry struct that hold trace and metric provider
type Providers struct {
	TraceProvider  *sdktrace.TracerProvider
	MetricProvider *sdkmetric.MeterProvider
	LogProvider    *sdklog.LoggerProvider
}

// NewProviders init Open Telemetry config
func NewProviders(ctx context.Context) (*Providers, error) {
	var providers Providers

	providersEnable, err := getProvidersEnable()
	if err != nil {
		return nil, err
	}

	resource, err := NewResources(ctx)
	if err != nil {
		return nil, err
	}

	if providersEnable.Trace {
		traceProvider, err := InitTraceProvider(ctx, resource)
		if err != nil {
			return nil, err
		}

		if traceProvider != nil {
			SetGlobalTraceProvider(traceProvider)
			SetGlobalContextPropagation()
			providers.TraceProvider = traceProvider
		}
	}

	if providersEnable.Metric {
		metricProvider, err := InitMetricProvider(ctx, resource)
		if err != nil {
			return nil, err
		}

		if metricProvider != nil {
			SetGlobalMetricProvider(metricProvider)
			providers.MetricProvider = metricProvider
		}
	}

	if providersEnable.Log {
		logProvider, err := InitLogProvider(ctx, resource)
		if err != nil {
			return nil, err
		}

		if logProvider != nil {
			providers.LogProvider = logProvider
		}
	}

	return &providers, nil
}

// Shutdown turn of trace and metric
func (o *Providers) Shutdown(ctx context.Context) error {
	if o.TraceProvider != nil {
		err := o.TraceProvider.Shutdown(ctx)
		if err != nil {
			return err
		}
	}

	if o.MetricProvider != nil {
		err := o.MetricProvider.Shutdown(ctx)
		if err != nil {
			return err
		}
	}

	if o.LogProvider != nil {
		err := o.LogProvider.Shutdown(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}
