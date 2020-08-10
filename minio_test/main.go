package main

import (
	"fmt"
	"github.com/minio/minio-go"
	uuid "github.com/satori/go.uuid"
	"log"
	"time"
)

func client() *minio.Client {
	endpoint := "192.168.88.14:9000"
	accessKeyID := "AKIAIOSFODNN7EXAMPLE"
	secretAccessKey := "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"
	useSSL := false

	client, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}

	return client
}

func main() {
	client := client()
	bucketName := "test"
	location := "us-east-1"
	exists, err := client.BucketExists(bucketName)
	if err != nil {
		log.Fatalln(err)
	}
	if !exists {
		// make bucket
		err := client.MakeBucket(bucketName, location)
		if err != nil {
			log.Fatalln(err)
		}
	}

	now := time.Now()
	// UPLOAD FILE
	objectName := "rust_book.pdf"
	filePath := "rust_book.pdf"
	contentType := "application/pdf"
	n, err := client.FPutObject(bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}
	since := time.Since(now)
	fmt.Println(since.Seconds())
	log.Printf("Successfully uploaded %s of size %d \n", objectName, n)
}


func NewUUID() string {
	v4 := uuid.NewV4()
	return v4.String()
}