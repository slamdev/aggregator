package aggregator

import (
	"testing"
	"reflect"
)

func assertEmptyAppArgs(value AppArgs, t *testing.T) {
	if !reflect.DeepEqual(value, AppArgs{}) {
		t.Fatalf("Expected value to be empty but got %v", value)
	}
}

func assertNil(value interface{}, t *testing.T) {
	if value != nil {
		t.Fatalf("Expected value to be nil but got %v", value)
	}
}

func assertNotNil(value interface{}, t *testing.T) {
	if value == nil {
		t.Fatalf("Expected value to be nil but got %v", value)
	}
}

func assertErrorMessage(err error, message string, t *testing.T) {
	assertNotNil(err, t)
	if err.Error() != message {
		t.Fatalf("Expected error to have message [%s] but got [%s]", message, err.Error())
	}
}

func assertFloatsLen(expected int, slice []float64, t *testing.T) {
	assertLen(expected, len(slice), t)
}

func assertMetricsLen(expected int, slice []Metric, t *testing.T) {
	assertLen(expected, len(slice), t)
}

func assertSensorsLen(expected int, slice []Sensor, t *testing.T) {
	assertLen(expected, len(slice), t)
}

func assertLen(expected int, actual int, t *testing.T) {
	if actual != expected {
		t.Fatalf("Expected slice to have %d size but got %d", expected, actual)
	}
}

func assertEqual(expected interface{}, actual interface{}, t *testing.T) {
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected %v be equal to %v", expected, actual)
	}
}
