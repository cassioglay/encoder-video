package services_test

import (
	"log"
	"os"
	"testing"

	"github.com/cassioglay/encoder/aplication/services"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func init() {

	err := godotenv.Load("../../../.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

}

func TestVideoUploadService(t *testing.T) {

	video, repo := prepare()

	//Create new video
	videoService := services.NewVideoService()
	videoService.Video = video
	videoService.VideoRepository = repo

	//Download video from bucket
	err := videoService.DownLoad("video-encoder-storage")
	require.Nil(t, err)

	//Fragement video
	err = videoService.Fragment()
	require.Nil(t, err)

	//Encoder
	err = videoService.Encode()
	require.Nil(t, err)

	//Upload video to bucket
	videoUpload := services.NewVideoUpload()
	videoUpload.OutputBucket = "video-encoder-storage"
	videoUpload.VideoPath = os.Getenv("localStoragePath") + "/" + video.ID

	doneUpload := make(chan string)
	go videoUpload.ProcessUpload(50, doneUpload)

	result := <-doneUpload
	require.Equal(t, result, "upload completed")
}
