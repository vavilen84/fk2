package s3

import (
	"app/utils"
	"github.com/astaxie/beego"
	"image"
	"os"
	"path"
	"time"
)

const tempDir = "/tmp"

func SaveImage(image image.Image, ext string) (filePath, uuid string, err error) {
	uuid = utils.GenerateUUID()
	filePath, err = SaveImageLocal(image, uuid, ext)
	if err != nil {
		beego.Error(err)
		return
	}
	err = SaveImageToS3(filePath)
	if err != nil {
		beego.Error(err)
		return
	}
	err = os.Remove(path.Join(tempDir, filePath))
	if err != nil {
		beego.Error(err)
		return
	}
	return
}

func generateSubfolderName(fileName string) string {
	currentTime := time.Now()
	return path.Join(currentTime.Format("2006-01-02"), fileName[:8], fileName[9:13])
}
