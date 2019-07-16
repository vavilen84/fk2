package s3

import (
	"image"
	"os"
	"path"
	"time"
)

const tempDir = "/tmp"

func SaveImage(image image.Image, ext string) (filePath string, err error) {
	result, err := SaveImageLocal(image, ext)
	if err != nil {
		return result, err
	}
	err = SaveImageToS3(result)
	if err != nil {
		return result, err
	}
	os.Remove(path.Join(tempDir, result))

	return result, err
}

func generateSubfolderName(fileName string) string {
	currentTime := time.Now()
	return path.Join(currentTime.Format("2006-01-02"), fileName[:8], fileName[9:13])
}
