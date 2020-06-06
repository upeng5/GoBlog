package utils

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

// ReadImage : Reads an image and returns a base64 image encoding as a string
func ReadImage(imgPath string) string {
	imgBytes, err := ioutil.ReadFile(imgPath)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	base64IMG := base64.StdEncoding.EncodeToString(imgBytes)
	return base64IMG
}

// DefaultAvatar : Returns the default avatar bytes
func DefaultAvatar() string {
	absPath, err := filepath.Abs("{Image Location}")
	if err != nil {
		fmt.Println("Path Error:", err)
		return ""
	}
	img := ReadImage(absPath)
	return img
}
