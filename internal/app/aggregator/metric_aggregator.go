package aggregator

import (
	"sort"
	"math"
)

func Aggregate(sensors []Sensor) ([]Metric) {
	temperatures := groupTemperaturesById(sensors)
	metrics := convertTemperaturesToMetrics(temperatures)
	sortMetricsByAverage(metrics)
	return metrics
}

func groupTemperaturesById(sensors []Sensor) map[string][]float64 {
	grouped := make(map[string][]float64)
	for _, sensor := range sensors {
		grouped[sensor.Id] = append(grouped[sensor.Id], sensor.Temperature)
	}
	return grouped
}

func sortMetricsByAverage(metrics []Metric) {
	sort.Slice(metrics, func(i, j int) bool {
		return metrics[i].Average < metrics[j].Average
	})
}

func convertTemperaturesToMetrics(temperatures map[string][]float64) []Metric {
	var metrics []Metric
	for id, values := range temperatures {
		sort.Float64s(values)
		metric := Metric{
			Id:      id,
			Average: calculateAverage(values),
			Median:  calculateMedian(values),
			Mode:    calculateMode(values),
		}
		metrics = append(metrics, metric)
	}
	return metrics
}

func calculateAverage(values []float64) float64 {
	sum := 0.0
	for _, value := range values {
		sum += value
	}
	average := sum / float64(len(values))
	return round(average, .5, 2)
}

func calculateMedian(values []float64) float64 {
	var median float64
	if len(values)%2 == 0 {
		middle := len(values) / 2
		median = calculateAverage([]float64{values[middle-1], values[middle]})
	} else {
		median = values[len(values)/2]
	}
	return round(median, .5, 2)
}

func calculateMode(values []float64) []float64 {
	countFrequencies := make(map[float64]int)
	for _, value := range values {
		countFrequencies[value] = countFrequencies[value] + 1
	}
	frequencies := getValues(countFrequencies)
	if len(distinct(frequencies)) == 1 {
		return []float64{}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(frequencies)))
	maxFrequency := frequencies[0]
	var mode []float64
	for value, frequency := range countFrequencies {
		if frequency == maxFrequency {
			mode = append(mode, round(value, .5, 2))
		}
	}
	sort.Float64s(mode)
	return mode
}

func round(val float64, roundOn float64, places int) float64 {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	return round / pow
}

func getValues(table map[float64]int) []int {
	values := make([]int, 0, len(table))
	for _, value := range table {
		values = append(values, value)
	}
	return values
}

func distinct(elements []int) []int {
	duplicates := map[int]bool{}
	var result []int
	for v := range elements {
		if duplicates[elements[v]] != true {
			duplicates[elements[v]] = true
			result = append(result, elements[v])
		}
	}
	return result
}

type Metric struct {
	Id      string
	Average float64
	Median  float64
	Mode    []float64
}
