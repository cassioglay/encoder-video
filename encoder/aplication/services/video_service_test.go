package services_test

import (
	"log"
	"testing"
	"time"

	"github.com/cassioglay/encoder/aplication/repositories"
	"github.com/cassioglay/encoder/aplication/services"
	"github.com/cassioglay/encoder/domain"
	"github.com/cassioglay/encoder/framework/database"
	"github.com/joho/godotenv"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func init() {

	err := godotenv.Load("../../.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

}

func prepare() (*domain.Video, repositories.VideoRepositoryDb) {
	//create a new video to test the Job
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "videos/bird.mp4"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}

	return video, repo
}

func TestDownload(t *testing.T) {

	video, repo := prepare()

	videoService := services.NewVideoService()
	videoService.Video = video
	videoService.VideoRepository = repo

	err := videoService.DownLoad("video-encoder-storage")
	require.Nil(t, err)

	err = videoService.Fragment()
	require.Nil(t, err)

	err = videoService.Encode()
	require.Nil(t, err)

	err = videoService.Finish()
	require.Nil(t, err)
}
