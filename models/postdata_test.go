package models

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/google/uuid"
)

func TestTimestampFormat(t *testing.T) {
	postData := PostData{
		Title:         "Some Title",
		UUID4:         uuid.New().String(),
		UnixTimestamp: 1597805784,
	}

	postData.AfterFind(GetTestDB())
	assert.Equal(t, postData.Timestamp, "2020-08-19T02:56:24+00:00")
}
