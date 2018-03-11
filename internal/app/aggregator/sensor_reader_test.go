package aggregator

import (
	"testing"
	"io/ioutil"
	"os"
	"fmt"
)

func TestReadToReturnErrorWhenFileContentIsNotValidJson(t *testing.T) {
	file := createFile("invalid data")
	defer os.Remove(file)
	sensors, err := Read(file)
	assertSensorsLen(0, sensors, t)
	assertErrorMessage(err, "invalid character 'i' looking for beginning of value", t)
}

func TestReadToReturnErrorWhenFileContentIsNotValidJsonArray(t *testing.T) {
	file := createFile("{}")
	defer os.Remove(file)
	sensors, err := Read(file)
	assertSensorsLen(0, sensors, t)
	assertErrorMessage(err, "json: cannot unmarshal object into Go value of type []aggregator.Sensor", t)
}

func TestReadToConvertJsonArrayToListOfSensors(t *testing.T) {
	id := "a"
	timestamp := int64(1509493641)
	temperature := 3.53
	file := createFile(fmt.Sprintf("[{\"id\": \"%s\",\"timestamp\": %d,\"temperature\": %f}]", id, timestamp, temperature))
	sensors, err := Read(file)
	assertNil(err, t)
	assertSensorsLen(1, sensors, t)
	assertEqual(Sensor{id, timestamp, temperature}, sensors[0], t)
}

func createFile(content string) string {
	file, _ := ioutil.TempFile("", "")
	ioutil.WriteFile(file.Name(), []byte(content), 0644)
	return file.Name()
}
