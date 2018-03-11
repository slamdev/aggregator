package main

import (
	"log"
	"os"
	"github.com/slamdev/aggregator/internal/app/aggregator"
)

func main() {
	var err error
	args, err := aggregator.ParseArgs()
	verifyErr(err)
	sensors, err := aggregator.Read(args.Filename)
	verifyErr(err)
	metrics := aggregator.Aggregate(sensors)
	err = aggregator.Write(os.Stdout, metrics)
	verifyErr(err)
}

func verifyErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
