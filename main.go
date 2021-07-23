package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/disintegration/imaging"
)

var inFolder string
var outFolder string

func init() {
	// Define command line arguments
	flag.StringVar(&inFolder, "in", "", "The input folder containing the images")

	flag.StringVar(&outFolder, "out", "", "The output folder where the results will be stored")

	// Print help message
	if len(os.Args) > 1 && (os.Args[1] == "--help" || os.Args[1] == "-h") {
		fmt.Print("IMAGE RESIZER\n\n")
		fmt.Print("Command line options:\n\n")
		flag.VisitAll(func(f *flag.Flag) {
			fmt.Printf("  --%s: %s\n", f.Name, f.Usage)
		})
		os.Exit(0)
	}

	// Parse arguments
	flag.Parse()

	if inFolder == "" {
		log.Fatalf("Input folder is undefined")
	}

	if outFolder == "" {
		log.Fatal("Output folder is undefined")
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

func resizeImage(file os.DirEntry) {
	src, err := imaging.Open(path.Join(inFolder, file.Name()))
	if err != nil {
		log.Fatal(err.Error())
	}

	src = imaging.Resize(src, 200, 0, imaging.Lanczos)

	err = imaging.Save(src, path.Join(outFolder, file.Name()))
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
	}
}

func main() {

	files, err := os.ReadDir(inFolder)
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, file := range files {
		resizeImage(file)
	}
}
