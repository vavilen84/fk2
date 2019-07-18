package utils

import (
	"github.com/satori/go.uuid"
	"image"
	"io"
	"path"
)

func GenerateUUID() string {
	return uuid.NewV4().String()
}

func GetImageExtension(r io.Reader) (format string, err error) {
	_, format, err = image.DecodeConfig(r)
	return
}

func GetTmpDir(subDir string) string {
	return path.Join("/tmp", subDir)
}
