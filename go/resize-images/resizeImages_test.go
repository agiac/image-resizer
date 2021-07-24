package resize_images

import (
	"os"
	"testing"
)

func TestResizeImages(t *testing.T) {
	inFolder := "../../test-images"
	outFolder := "/tmp/image-resizer-output"
	width := 500

	_ = os.RemoveAll(outFolder)
	_ = os.Mkdir(outFolder, 0755)

	ResizeImages(inFolder, outFolder, width)

	inFiles, _ := os.ReadDir(inFolder)
	outFiles, _ := os.ReadDir(outFolder)

	if len(inFiles) != len(outFiles) {
		t.Errorf("Expecting %d files in output folder, found: %d", len(inFiles), len(outFiles))
	}
}
