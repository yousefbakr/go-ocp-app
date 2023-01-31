package main

import (
        "log"

        "github.com/minio/minio-go/v7"
        "github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {
        endpoint := "http://localhost:9000"
        accessKeyID := "hz9zhbwDXQpxYAjX"
        secretAccessKey := "3B8lOjyaP0l1JqLXx4bPa65rueJbhpiZG"
        useSSL := true

        // Initialize minio client object.
        minioClient, err := minio.New(endpoint, &minio.Options{
                Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
                Secure: useSSL,
        })
        if err != nil {
                log.Fatalln(err)
        }

        log.Printf("%#v\n", minioClient) // minioClient is now set up
}
