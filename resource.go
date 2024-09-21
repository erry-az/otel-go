package otel

import (
	"context"

	"go.opentelemetry.io/otel/sdk/resource"
)

// NewResources init new OTLP resource with get from env by default
// env var is OTEL_RESOURCE_ATTRIBUTES and OTEL_SERVICE_NAME
// example env :
// OTEL_SERVICE_NAME=example-service
// OTEL_RESOURCE_ATTRIBUTES=container=docker,host=local
func NewResources(ctx context.Context, opts ...resource.Option) (*resource.Resource, error) {
	return resource.New(ctx, append([]resource.Option{resource.WithFromEnv()}, opts...)...)
}
