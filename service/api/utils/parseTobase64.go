package utils

import (
	"encoding/base64"
	"io/ioutil"
	"os"
)

func ImageToBase64(filename string) (string, error) {
	imageFile, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer imageFile.Close()

	imageData, err := ioutil.ReadAll(imageFile)
	if err != nil {
		return "", err
	}

	base64 := base64.StdEncoding.EncodeToString(imageData)
	return base64, nil
}
