package main

import (
	"flag"
	"fmt"
	resizeImages "github.com/agiac/image-resizer/resize-images"
	"log"
	"os"
	"time"
)

var inFolder string
var outFolder string
var width int

func init() {
	// Define command line arguments
	flag.StringVar(&inFolder, "in", "", "The input folder containing the images")
	flag.StringVar(&outFolder, "out", "", "The output folder where the results will be stored")
	flag.IntVar(&width, "width", 0, "The resize width")

	// Parse arguments
	flag.Parse()

	if inFolder == "" {
		log.Fatalf("Input folder is undefined")
	}

	if outFolder == "" {
		log.Fatal("Output folder is undefined")
	}

	if width == 0 {
		log.Fatal("Width is undefined")
	}

	if width < 0 {
		log.Fatal("Width must be a positive number")
	}

	// Check whether the input folder exists
	_, err := os.Stat(inFolder)

	if os.IsNotExist(err) {
		log.Fatalf("Input folder '%s' does not exist", inFolder)
	}

	// Delete if it already exists and create a new output folder
	err = os.RemoveAll(outFolder)

	if err != nil {
		log.Fatal(err.Error())
	}

	err = os.Mkdir(outFolder, 0755)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	start := time.Now()

	resizeImages.ResizeImages(inFolder, outFolder, width)

	duration := int(time.Since(start).Seconds())

	fmt.Printf("Done in %d seconds\n", duration)
}
