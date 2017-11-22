package resize

import (
	"image/jpeg"
	"log"
	"os"
)

func ResizeImage(fileName string, fileRename string) {
	// open "test.jpg"
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	// decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	m := Resize(1000, 0, img, Lanczos3)

	out, err := os.Create(fileRename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, m, nil)
	os.Remove(fileName)
}
