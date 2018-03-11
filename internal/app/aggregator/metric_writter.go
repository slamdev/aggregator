package aggregator

import (
	"io"
	"encoding/json"
)

func Write(writer io.Writer, metrics []Metric) error {
	var err error
	data, err := json.Marshal(metrics)
	if err != nil {
		return err
	}
	_, err = writer.Write(data)
	return err
}
