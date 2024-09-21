package otel

import (
	"context"

	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

// Providers open telemetry struct that hold trace and metric provider
type Providers struct {
	TraceProvider  *sdktrace.TracerProvider
	MetricProvider *sdkmetric.MeterProvider
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

	return nil
}

// NewProviders init Open Telemetry config
func NewProviders(ctx context.Context) (*Providers, error) {
	resource, err := NewResources(ctx)
	if err != nil {
		return nil, err
	}

	traceProvider, err := InitTraceProvider(ctx, resource)
	if err != nil {
		return nil, err
	}

	metricProvider, err := InitMetricProvider(ctx, resource)
	if err != nil {
		return nil, err
	}

	if traceProvider != nil {
		SetGlobalTraceProvider(traceProvider)
		SetGlobalContextPropagation()
	}

	if metricProvider != nil {
		SetGlobalMetricProvider(metricProvider)
	}

	return &Providers{
		TraceProvider:  traceProvider,
		MetricProvider: metricProvider,
	}, nil
}
