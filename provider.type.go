package otel

import "errors"

type ProviderType string

const (
	providerTypeTrace  ProviderType = "trace"
	providerTypeMetric ProviderType = "metric"
	providerTypeLog    ProviderType = "log"
)

// ErrInvalidProviderType invalid provider type error
var ErrInvalidProviderType = errors.New("invalid provider type")

func validateProviderType(t string) (ProviderType, error) {
	if !(t == string(providerTypeTrace) || t == string(providerTypeMetric) || t == string(providerTypeLog)) {
		return "", ErrInvalidProviderType
	}

	return ProviderType(t), nil
}

type ProvidersEnable struct {
	Trace  bool
	Metric bool
	Log    bool
}
