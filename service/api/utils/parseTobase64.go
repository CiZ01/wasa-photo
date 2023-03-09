package utils

import (
	"encoding/base64"
	"io"
	"os"
)

func ImageToBase64(filename string) (string, error) {
	imageFile, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer func() { err = imageFile.Close() }()

	imageData, err := io.ReadAll(imageFile)
	if err != nil {
		return "", err
	}

	base64 := base64.StdEncoding.EncodeToString(imageData)
	return base64, err
}
