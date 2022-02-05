package ezs3lib

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3Connection struct {
	session    *session.Session
	uploader   *s3manager.Uploader
	downloader *s3manager.Downloader
	deleter    *s3manager.BatchDelete
	bucket     string
}

func ConnectS3(bucket string, endpoint string, region string) *S3Connection {

	s := session.Must(session.NewSession(&aws.Config{
		Region:   &region,
		Endpoint: &endpoint,
	}))

	uploader := s3manager.NewUploader(s)
	downloader := s3manager.NewDownloader(s)
	deleter := s3manager.NewBatchDelete(s)

	return &S3Connection{
		session:    s,
		uploader:   uploader,
		downloader: downloader,
		deleter:    deleter,
		bucket:     bucket,
	}
}

func (conn *S3Connection) UploadFile(origin string, target string) error {
	file, err := os.Open(origin)

	if err != nil {
		return err
	}
	defer file.Close()

	_, err = conn.uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(conn.bucket),
		Key:    aws.String(target),
		Body:   file,
	})

	return err
}

func (conn *S3Connection) DownloadFile(origin string, target string) error {
	newFile, err := os.Create(target)
	if err != nil {
		return err
	}
	_, err = conn.downloader.Download(
		newFile,
		&s3.GetObjectInput{
			Bucket: aws.String(conn.bucket),
			Key:    aws.String(origin),
		})
	return err
}

func (conn *S3Connection) DeleteFile(origin string) error {
	deleteObject := []s3manager.BatchDeleteObject{
		{
			Object: &s3.DeleteObjectInput{
				Bucket: aws.String(conn.bucket),
				Key:    aws.String(origin),
			},
		},
	}
	err := conn.deleter.Delete(aws.BackgroundContext(), &s3manager.DeleteObjectsIterator{
		Objects: deleteObject,
	})

	return err
}
