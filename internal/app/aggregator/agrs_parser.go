package aggregator

import (
	"errors"
	"flag"
	"os"
)

func ParseArgs() (AppArgs, error) {
	var args AppArgs
	flag.Parse()
	if flag.NArg() == 0 {
		return args, errors.New("path to a file should be provided")
	}
	if flag.NArg() > 1 {
		return args, errors.New("only one argument is supported")
	}
	file := flag.Arg(0)
	fileInfo, err := os.Stat(file)
	if err != nil {
		return args, err
	}
	if !fileInfo.Mode().IsRegular() {
		return args, errors.New("provided path is not a file")
	}
	return AppArgs{Filename: file}, nil
}

type AppArgs struct {
	Filename string
}
