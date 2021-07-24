package resize_images

import (
	"fmt"
	"github.com/disintegration/imaging"
	"log"
	"os"
	"path"
	"sync"
)

func resizeImage(fileName string, inFolder string, outFolder string, width int, wg *sync.WaitGroup) {
	defer wg.Done()

	src, err := imaging.Open(path.Join(inFolder, fileName))
	if err != nil {
		log.Fatal(err.Error())
	}

	dest := imaging.Resize(src, width, 0, imaging.Lanczos)

	err = imaging.Save(dest, path.Join(outFolder, fileName))
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
	}
}

func ResizeImages(inFolder string, outFolder string, width int) {
	files, err := os.ReadDir(inFolder)
	if err != nil {
		log.Fatal(err.Error())
	}

	var wg = sync.WaitGroup{}

	wg.Add(len(files))

	for _, file := range files {
		go resizeImage(file.Name(), inFolder, outFolder, width, &wg)
	}

	wg.Wait()
}
