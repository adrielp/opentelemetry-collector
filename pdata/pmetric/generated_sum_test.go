// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by "pdata/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package pmetric

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.opentelemetry.io/collector/pdata/internal"
	otlpmetrics "go.opentelemetry.io/collector/pdata/internal/data/protogen/metrics/v1"
)

func TestSum_MoveTo(t *testing.T) {
	ms := generateTestSum()
	dest := NewSum()
	ms.MoveTo(dest)
	assert.Equal(t, NewSum(), ms)
	assert.Equal(t, generateTestSum(), dest)
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { ms.MoveTo(newSum(&otlpmetrics.Sum{}, &sharedState)) })
	assert.Panics(t, func() { newSum(&otlpmetrics.Sum{}, &sharedState).MoveTo(dest) })
}

func TestSum_CopyTo(t *testing.T) {
	ms := NewSum()
	orig := NewSum()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
	orig = generateTestSum()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { ms.CopyTo(newSum(&otlpmetrics.Sum{}, &sharedState)) })
}

func TestSum_AggregationTemporality(t *testing.T) {
	ms := NewSum()
	assert.Equal(t, AggregationTemporality(otlpmetrics.AggregationTemporality(0)), ms.AggregationTemporality())
	testValAggregationTemporality := AggregationTemporality(otlpmetrics.AggregationTemporality(1))
	ms.SetAggregationTemporality(testValAggregationTemporality)
	assert.Equal(t, testValAggregationTemporality, ms.AggregationTemporality())
}

func TestSum_IsMonotonic(t *testing.T) {
	ms := NewSum()
	assert.False(t, ms.IsMonotonic())
	ms.SetIsMonotonic(true)
	assert.True(t, ms.IsMonotonic())
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { newSum(&otlpmetrics.Sum{}, &sharedState).SetIsMonotonic(true) })
}

func TestSum_DataPoints(t *testing.T) {
	ms := NewSum()
	assert.Equal(t, NewNumberDataPointSlice(), ms.DataPoints())
	fillTestNumberDataPointSlice(ms.DataPoints())
	assert.Equal(t, generateTestNumberDataPointSlice(), ms.DataPoints())
}

func generateTestSum() Sum {
	tv := NewSum()
	fillTestSum(tv)
	return tv
}

func fillTestSum(tv Sum) {
	tv.orig.AggregationTemporality = otlpmetrics.AggregationTemporality(1)
	tv.orig.IsMonotonic = true
	fillTestNumberDataPointSlice(newNumberDataPointSlice(&tv.orig.DataPoints, tv.state))
}
