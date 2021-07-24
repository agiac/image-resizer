# Image resizer scripts

Two simple scripts for resizing images

### Command line arguments
- _```--in```_ Input folder containing the images to resize;
- _```--out```_ Output folder where the resized images will be saved. ATTENTION: the folder will always be deleted at the beginning of the script, if it is already existing;
- _```--width```_ The width to which to resize the images

## JS

```node main.js --in /Pictures/Venice --out /Pictures/Venice-resized --width 500```

## GO

```go run ./main.go --in /Pictures/Venice --out /Pictures/Venice-resized --width 500```