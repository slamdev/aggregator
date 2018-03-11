package aggregator

import (
	"testing"
	"bytes"
)

func TestWriteToWriteEmptyArrayWhenNoMetricsProvided(t *testing.T) {
	var writer bytes.Buffer
	Write(&writer, []Metric{})
	assertEqual("[]", writer.String(), t)
}

func TestWriteToWriteMetricsAsJsonArray(t *testing.T) {
	var writer bytes.Buffer
	Write(&writer, []Metric{
		{"id1", 1., 2., []float64{3.}},
	})
	assertEqual("[{\"Id\":\"id1\",\"Average\":1,\"Median\":2,\"Mode\":[3]}]", writer.String(), t)
}
