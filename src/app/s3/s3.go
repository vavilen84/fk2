package s3

import (
	"app/utils"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"
)

func SaveImageToS3(subDir, filename string) (err error) {
	tmpDir := utils.GetTmpDir(subDir)
	filepath := path.Join(tmpDir, filename)
	bucket := os.Getenv("AWS_S3_BUCKET")
	file, err := os.Open(filepath)
	if err != nil {
		err = errors.New("Unable to open file: " + filepath)
		return err
	}
	region := os.Getenv("AWS_DEFAULT_REGION")
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	uploader := s3manager.NewUploader(sess)
	contentType, err := GetFileContentType(file)
	if err != nil {
		err = errors.New("Error get content type: " + filepath)
		return err
	}
	defer file.Close()
	_, err = uploader.Upload(&s3manager.UploadInput{
		ContentType: aws.String(contentType),
		Bucket:      aws.String(bucket),
		Key:         aws.String(path.Join(subDir, filename)),
		Body:        file,
	})
	if err != nil {
		err = errors.New("Unable to upload file: " + filepath)
		return err
	}
	fmt.Printf("Successfully uploaded: %s", filepath)

	return nil
}

func GetFileContentType(out *os.File) (string, error) {

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}

	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}

func GenerateSubfolderName(fileName string) string {
	year, month, day := time.Now().Date()
	return path.Join(
		strconv.Itoa(int(year)),
		strconv.Itoa(int(month)),
		strconv.Itoa(int(day)),
		fileName[:8],
		fileName[9:13],
	)
}
