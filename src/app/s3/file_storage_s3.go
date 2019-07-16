package s3

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"net/http"
	"os"
	"path"
)

func SaveImageToS3(filePath string) (err error) {
	localFilePath := path.Join(tempDir, filePath)
	bucket := os.Getenv("AWS_S3_BUCKET")
	file, err := os.Open(localFilePath)
	if err != nil {
		err = errors.New("Unable to open file: " + localFilePath)
		return err
	}
	region := os.Getenv("AWS_DEFAULT_REGION")
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	uploader := s3manager.NewUploader(sess)
	contentType, err := GetFileContentType(file)
	if err != nil {
		err = errors.New("Error get content type: " + filePath)
		return err
	}
	defer file.Close()
	_, err = uploader.Upload(&s3manager.UploadInput{
		ContentType: aws.String(contentType),
		Bucket:      aws.String(bucket),
		Key:         aws.String(filePath),
		Body:        file,
	})
	if err != nil {
		err = errors.New("Unable to upload file: " + filePath)
		return err
	}
	fmt.Printf("Successfully uploaded: %s", filePath)

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
