// Code generated by mdatagen. DO NOT EDIT.

package metadatatest

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
	"go.opentelemetry.io/otel/sdk/metric/metricdata/metricdatatest"

	"go.opentelemetry.io/collector/cmd/mdatagen/internal/samplereceiver/internal/metadata"
	"go.opentelemetry.io/collector/component/componenttest"
)

func TestSetupTelemetry(t *testing.T) {
	testTel := componenttest.NewTelemetry()
	tb, err := metadata.NewTelemetryBuilder(testTel.NewTelemetrySettings())
	require.NoError(t, err)
	defer tb.Shutdown()
	require.NoError(t, tb.RegisterProcessRuntimeTotalAllocBytesCallback(func(_ context.Context, observer metric.Int64Observer) error {
		observer.Observe(1)
		return nil
	}))
	require.NoError(t, tb.RegisterQueueLengthCallback(func(_ context.Context, observer metric.Int64Observer) error {
		observer.Observe(1)
		return nil
	}))
	tb.BatchSizeTriggerSend.Add(context.Background(), 1)
	tb.QueueCapacity.Record(context.Background(), 1)
	tb.RequestDuration.Record(context.Background(), 1)
	AssertEqualBatchSizeTriggerSend(t, testTel,
		[]metricdata.DataPoint[int64]{{Value: 1}},
		metricdatatest.IgnoreTimestamp())
	AssertEqualProcessRuntimeTotalAllocBytes(t, testTel,
		[]metricdata.DataPoint[int64]{{Value: 1}},
		metricdatatest.IgnoreTimestamp())
	AssertEqualQueueCapacity(t, testTel,
		[]metricdata.DataPoint[int64]{{Value: 1}},
		metricdatatest.IgnoreTimestamp())
	AssertEqualQueueLength(t, testTel,
		[]metricdata.DataPoint[int64]{{Value: 1}},
		metricdatatest.IgnoreTimestamp())
	AssertEqualRequestDuration(t, testTel,
		[]metricdata.HistogramDataPoint[float64]{{}}, metricdatatest.IgnoreValue(),
		metricdatatest.IgnoreTimestamp())

	require.NoError(t, testTel.Shutdown(context.Background()))
}
