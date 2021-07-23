package main

import (
	"flag"
)

var inFolder *string
var outFolder *string

func init() {
	inFolder = flag.String("in", "", "The input folder containing the images")
	outFolder = flag.String("out", "", "The output folder where the results will be stored")
}

func main() {
	flag.Parse()

	print(inFolder)
	print(outFolder)
}
