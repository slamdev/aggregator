package aggregator

import (
	"encoding/json"
	"io/ioutil"
)

func Read(filename string) ([]Sensor, error) {
	var sensors []Sensor
	var err error
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return sensors, err
	}
	err = json.Unmarshal(content, &sensors)
	if err != nil {
		return sensors, err
	}
	return sensors, nil
}

type Sensor struct {
	Id          string
	Timestamp   int64
	Temperature float64
}
