package utils_test

import (
	"testing"

	"github.com/cassioglay/encoder/framework/utils"
	"github.com/stretchr/testify/require"
)

func TestIsJson(t *testing.T) {
	json := `
		{
			"id": "123",
			"file_path" : "convite.mp4",
			"status" : "pending"
		}
	`

	err := utils.IsJson(json)
	require.Nil(t, err)

	json = `ABC`
	err = utils.IsJson(json)
	require.Error(t, err)

}
