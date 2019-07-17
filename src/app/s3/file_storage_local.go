package s3

import (
	"github.com/astaxie/beego"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path"
)

func SaveImageLocal(image image.Image, uuid, ext string) (filePath string, err error) {
	fileName := uuid + "." + ext
	subDir := generateSubfolderName(fileName)
	result := path.Join(subDir, fileName)
	dirToSave := path.Join(tempDir, subDir)
	fullFilePath := path.Join(dirToSave, fileName)
	err = os.MkdirAll(dirToSave, 0755)
	if err != nil {
		return result, err
	}
	outputFile, err := os.Create(fullFilePath)
	defer outputFile.Close()
	if err != nil {
		return result, err
	}
	err = WriteByType(ext, outputFile, image)
	if err != nil {
		return result, err
	}
	return result, nil
}

func encodeByType(ext string, w io.Writer, img image.Image) (err error) {
	switch ext {
	case "gif":
		err = gif.Encode(w, img, nil)
	case "jpeg":
		err = jpeg.Encode(w, img, nil)
	case "jpg":
		err = jpeg.Encode(w, img, nil)
	case "png":
		err = png.Encode(w, img)
	}
	if err != nil {
		beego.Error(err)
	}
	return
}

func WriteByType(imageType string, file *os.File, imgData image.Image) (err error) {
	err = encodeByType(imageType, file, imgData)
	return
}
