package utils

import (
	"image/jpeg"
	"os"

	"github.com/nfnt/resize"
)

func SaveAndCrop(filename string, w uint, h uint) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer func() { err = file.Close() }()

	// Decodifica l'immagine in un oggetto image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		return err
	}

	resizedImg := resize.Resize(w, h, img, resize.NearestNeighbor)
	// Salva l'immagine croppata su disco
	out, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func() { err = out.Close() }()
	if err := jpeg.Encode(out, resizedImg, nil); err != nil {
		return err
	}

	return err
}
