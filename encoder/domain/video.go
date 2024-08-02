package domain

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type Video struct {
	ID         string
	ResourceID string
	FilePath   string
	CreatedAt  time.Time
}

func NewVideo() *Video {
	return &Video{}
}

func (video *Video) Validate() error {
	_, err := govalidator.ValidateStruct(video)

	if err != nil {
		return err
	}
	return nil
}
