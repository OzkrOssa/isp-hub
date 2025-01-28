package instrumenting

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/sdk/trace"
)

type ExporterType string

const (
	HTTP ExporterType = "http"
	GRPC ExporterType = "grpc"
)

type InstrumentingConfig struct {
	ServiceName    string
	ServiceVersion string
	ExporterType   ExporterType
}

func NewTracing(ctx context.Context, config InstrumentingConfig) (*trace.TracerProvider, error) {
	switch config.ExporterType {
	case HTTP:
		return NewHttpTraceProvider(ctx, config.ServiceName, config.ServiceVersion)
	case GRPC:
		return NewGrpcTraceProvider(ctx, config.ServiceName, config.ServiceVersion)
	default:
		return nil, fmt.Errorf("unsupported exporter type: %s", config.ExporterType)
	}
}
