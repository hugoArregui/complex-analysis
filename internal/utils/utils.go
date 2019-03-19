package utils

import (
	"image"
	"image/png"
	"os"
)

func Save(im image.Image, filename string) error {
	file, err := os.Create(filename)

	if err != nil {
		return err
	}

	defer file.Close()
	if err = png.Encode(file, im); err != nil {
		return err
	}

	return nil
}
