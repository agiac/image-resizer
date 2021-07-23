package resize_images

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestResizeImages(t *testing.T) {
	inFolder := "../test/in"
	outFolder := "../test/out"
	width := 500

	_ = os.RemoveAll(outFolder)
	_ = os.Mkdir(outFolder, 0755)

	start := time.Now()
	ResizeImages(inFolder, outFolder, width)
	duration := time.Since(start).Seconds()

	inFiles, _ := os.ReadDir(inFolder)
	outFiles, _ := os.ReadDir(outFolder)

	if len(inFiles) != len(outFiles) {
		t.Errorf("Expecting %d files in output folder, found: %d", len(inFiles), len(outFiles))
	}

	fmt.Printf("Duration: %d seconds\n", int(duration))
}
