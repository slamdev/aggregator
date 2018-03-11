# Sensor Aggregator [![Build Status](https://travis-ci.org/slamdev/aggregator.svg?branch=master)](https://travis-ci.org/slamdev/aggregator)

[Alternative Java implementation](https://github.com/slamdev/sensor-aggregator)

[Alternative JavaScript implementation](https://github.com/slamdev/aggregator-js)

## Run

Prerequisites:
* docker should be installed

Run from root project directory:

```sh
docker build . -t aggregator
docker run -v $(pwd)/test/sampleInput.json:/file.json aggregator /file.json
```

## Test

Prerequisites:
* docker should be installed

Run from root project directory:

```sh
docker run --rm -v $(pwd):/go/src/github.com/slamdev/aggregator -w /go/src/github.com/slamdev/aggregator golang go test ./...
```

## Used tools

* Go 1.10
* Docker (split into build and run on order to reduce resulting image size)
* Travis CI for continuous integration

## Purpose

We are collecting temperature data about fridges in a supermarket. Imagine we have data from different fridge sensors aggregated into a single JSON array (where an individual sensor is identified by an id):

```json
[{"id": "a","timestamp": 1509493641,"temperature": 3.53},
{"id": "b","timestamp": 1509493642,"temperature": 4.13},
{"id": "c","timestamp": 1509493643,"temperature": 3.96},
{"id": "a","timestamp": 1509493644,"temperature": 3.63},
{"id": "c","timestamp": 1509493645,"temperature": 3.96},
{"id": "a","timestamp": 1509493645,"temperature": 4.63},
{"id": "a","timestamp": 1509493646,"temperature": 3.53},
{"id": "b","timestamp": 1509493647,"temperature": 4.15},
{"id": "c","timestamp": 1509493655,"temperature": 3.95},
{"id": "a","timestamp": 1509493677,"temperature": 3.66},
{"id": "b","timestamp": 1510113646,"temperature": 4.15},
{"id": "c","timestamp": 1510127886,"temperature": 3.36},
{"id": "c","timestamp": 1510127892,"temperature": 3.36},
{"id": "a","timestamp": 1510128112,"temperature": 3.67},
{"id": "b","timestamp": 1510128115,"temperature": 3.88}]
```

Create an application that outputs the average, median and mode temperature for each fridge sensor to 2 decimal places in the following JSON format:

```json
[{"id":"c","average":3.72,"median":3.95,"mode":[3.36,3.96]},
{"id":"a","average":3.78,"median":3.65,"mode":[3.53]},
{"id":"b","average":4.08,"median":4.14,"mode":[4.15]}]
```
