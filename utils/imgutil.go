package utils

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

// ReadImage : Reads an image and returns a byte slice
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
	absPath, err := filepath.Abs("C:/Users/somoa/Projects/go-work/goblog/images/default.png")
	if err != nil {
		fmt.Println("Path Error:", err)
		return ""
	}
	img := ReadImage(absPath)
	return img
}
