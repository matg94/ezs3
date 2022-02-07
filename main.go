package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/matg94/ezs3/ezs3lib"
)

func main() {
	var bucket string
	var endpoint string
	var origin string
	var target string
	var upload bool
	var download bool
	var delete bool

	flag.StringVar(&endpoint, "e", "-", "Specify the S3 target Endpoint.")
	flag.StringVar(&bucket, "b", "-", "Specify the S3 target Bucket.")
	flag.StringVar(&origin, "f", "-", "Specify the origin path for the file you'd like to upload / download.")
	flag.StringVar(&target, "t", "-", "Specify the destination path where you would like the file saved. Default is origin.")
	flag.BoolVar(&upload, "upload", false, "Uploads target file to S3.")
	flag.BoolVar(&download, "download", false, "Downloads the target file from S3.")
	flag.BoolVar(&delete, "delete", false, "Deletes the origin file from S3.")

	flag.Parse()

	if !upload && !download && !delete {
		fmt.Println("No action has been set. Check -h for help.")
		return
	}

	if (upload && download) || (upload && delete) {
		fmt.Println("Cannot upload & download, or upload & delete, in the same command.")
		return
	}

	if target == "-" {
		target = origin
	}

	if endpoint == "-" {
		if os.Getenv("AWS_ENDPOINT") != "" {
			endpoint = os.Getenv("AWS_ENDPOINT")
		} else {
			fmt.Println("Endpoint not found in either command or environment. Set env var AWS_ENDPOINT or use -e.")
			return
		}
	}

	if bucket == "-" {
		if os.Getenv("AWS_BUCKET") != "" {
			bucket = os.Getenv("AWS_BUCKET")
		} else {
			fmt.Println("Bucket not found in either command or environment. Set env var AWS_BUCKET or use -b.")
			return
		}
	}

	if origin == "-" {
		fmt.Println("File origin not set. Use -f to set filepath, or check -h for help.")
		return
	}

	s3Connection := ezs3lib.ConnectS3(bucket, endpoint, strings.Split(endpoint, "-")[0])

	if download {
		err := s3Connection.DownloadFile(origin, target)
		if err != nil {
			fmt.Printf("Error when downloading: %v\n", err)
		}
	}

	if upload {
		err := s3Connection.UploadFile(origin, target)
		if err != nil {
			fmt.Printf("Error when uploading: %v\n", err)
		}
	}

	if delete {
		err := s3Connection.DeleteFile(origin)
		if err != nil {
			fmt.Printf("Error when deleting: %v\n", err)
		}
	}

}
