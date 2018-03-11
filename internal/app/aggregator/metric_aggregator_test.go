package aggregator

import (
	"testing"
)

func TestAggregateToReturnEmptyListForEmptySensors(t *testing.T) {
	var sensors []Sensor
	metrics := Aggregate(sensors)
	assertMetricsLen(0, metrics, t)
}

func TestAggregateToReturnSingleMetricWithSameIdAndTemperatureForSingleSensor(t *testing.T) {
	id := "id"
	temperature := 1.
	sensor := Sensor{id, 1, temperature}
	metrics := Aggregate([]Sensor{sensor})
	metric := Metric{id, temperature, temperature, []float64{}}
	assertMetricsLen(1, metrics, t)
	assertEqual(metric, metrics[0], t)
}

func TestAggregateToGroupSensorsById(t *testing.T) {
	id1 := "id1"
	id2 := "id2"
	metrics := Aggregate([]Sensor{
		{id1, 0, 0},
		{id2, 1, 1},
		{id1, 2, 0},
	})
	assertMetricsLen(2, metrics, t)
	assertEqual(id1, metrics[0].Id, t)
	assertEqual(id2, metrics[1].Id, t)
}

func TestAggregateToOrderSensorsByAverage(t *testing.T) {
	id1 := "id1"
	id2 := "id2"
	id3 := "id3"
	metrics := Aggregate([]Sensor{
		{id1, 3, 3},
		{id2, 1, 1},
		{id3, 2, 2},
	})
	assertMetricsLen(3, metrics, t)
	assertEqual(id2, metrics[0].Id, t)
	assertEqual(id3, metrics[1].Id, t)
	assertEqual(id1, metrics[2].Id, t)
}

func TestAggregateToCountAverageMetricForNonZeroTemperature(t *testing.T) {
	metrics := Aggregate([]Sensor{
		{"id", 1, 1},
		{"id", 1, 2},
		{"id", 1, 3},
	})
	assertMetricsLen(1, metrics, t)
	assertEqual(2., metrics[0].Average, t)
}

func TestAggregateToCountAverageMetricForZeroTemperature(t *testing.T) {
	metrics := Aggregate([]Sensor{
		{"id", 1, 0},
		{"id", 1, 0},
		{"id", 1, 0},
	})
	assertMetricsLen(1, metrics, t)
	assertEqual(0., metrics[0].Average, t)
}

func TestAggregateToCountMedianMetricForNonZeroOddTemperature(t *testing.T) {
	metrics := Aggregate([]Sensor{
		{"id", 1, 1},
		{"id", 1, 2},
		{"id", 1, 3},
	})
	assertMetricsLen(1, metrics, t)
	assertEqual(2., metrics[0].Median, t)
}

func TestAggregateToCountMedianMetricForNonZeroEvenTemperature(t *testing.T) {
	metrics := Aggregate([]Sensor{
		{"id", 1, 1},
		{"id", 1, 2},
		{"id", 1, 3},
		{"id", 1, 4},
	})
	assertMetricsLen(1, metrics, t)
	assertEqual(2.5, metrics[0].Median, t)
}

func TestAggregateToCountMedianMetricForZeroTemperature(t *testing.T) {
	metrics := Aggregate([]Sensor{
		{"id", 1, 0},
		{"id", 1, 0},
		{"id", 1, 0},
	})
	assertMetricsLen(1, metrics, t)
	assertEqual(0., metrics[0].Median, t)
}

func TestAggregateToCountModeMetricForNonZeroNonRepeatTemperature(t *testing.T) {
	metrics := Aggregate([]Sensor{
		{"id", 1, 1},
		{"id", 1, 2},
		{"id", 1, 3},
	})
	assertMetricsLen(1, metrics, t)
	assertFloatsLen(0, metrics[0].Mode, t)
}

func TestAggregateToCountModeMetricForNonZeroSingleRepeatTemperature(t *testing.T) {
	metrics := Aggregate([]Sensor{
		{"id", 1, 1},
		{"id", 1, 2},
		{"id", 1, 2},
	})
	assertMetricsLen(1, metrics, t)
	assertFloatsLen(1, metrics[0].Mode, t)
	assertEqual(2., metrics[0].Mode[0], t)
}

func TestAggregateToCountModeMetricForNonZeroMultiRepeatTemperature(t *testing.T) {
	metrics := Aggregate([]Sensor{
		{"id", 1, 1},
		{"id", 1, 1},
		{"id", 1, 2},
		{"id", 1, 2},
		{"id", 1, 3},
	})
	assertMetricsLen(1, metrics, t)
	assertFloatsLen(2, metrics[0].Mode, t)
	assertEqual(1., metrics[0].Mode[0], t)
	assertEqual(2., metrics[0].Mode[1], t)
}

func TestAggregateToOrderModeMetric(t *testing.T) {
	metrics := Aggregate([]Sensor{
		{"id", 1, 2},
		{"id", 1, 1},
		{"id", 1, 2},
		{"id", 1, 1},
		{"id", 1, 3},
	})
	assertMetricsLen(1, metrics, t)
	assertFloatsLen(2, metrics[0].Mode, t)
	assertEqual(1., metrics[0].Mode[0], t)
	assertEqual(2., metrics[0].Mode[1], t)
}

func TestAggregateToCountModeMetricForZeroTemperature(t *testing.T) {
	metrics := Aggregate([]Sensor{
		{"id", 1, 0},
		{"id", 1, 0},
		{"id", 1, 0},
	})
	assertMetricsLen(1, metrics, t)
	assertFloatsLen(0, metrics[0].Mode, t)
}

func TestAggregateToRoundMetricsTo2DecimalPoints(t *testing.T) {
	metrics := Aggregate([]Sensor{
		{"id", 1, 1.155},
		{"id", 1, 1.155},
	})
	expected := 1.16
	assertMetricsLen(1, metrics, t)
	assertEqual(expected, metrics[0].Average, t)
	assertEqual(expected, metrics[0].Median, t)
}
