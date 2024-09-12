package services

import (
	"context"
	"io"
	"os"
	"strings"

	"cloud.google.com/go/storage"
)

type VideoUpload struct {
	Paths        []string
	VideoPath    string
	OutputBucket string
	Errors       []string
}

func NewVideoUpload() *VideoUpload {
	return &VideoUpload{}
}

func (vu *VideoUpload) uploadOject(objectPath string, client *storage.Client, ctx context.Context) error {

	//Get file to upload
	path := strings.Split(objectPath, os.Getenv("localStoragePath")+"/")

	f, err := os.Open(objectPath)

	if err != nil {
		return err
	}

	f.Close()

	//Connect to google cloud storage
	wc := client.Bucket(vu.OutputBucket).Object(path[1]).NewWriter(ctx)
	//Add permission to google cloud storage
	wc.ACL = []storage.ACLRule{{Entity: storage.AllUsers, Role: storage.RoleReader}}

	if _, err = io.Copy(wc, f); err != nil {
		return err
	}

	if err := wc.Close(); err != nil {
		return err
	}

	return nil

}
