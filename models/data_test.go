package models

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/google/uuid"
)

func TestTimestampFormat(t *testing.T) {
	data := Data{
		Name:          "Some Title",
		Stage:         1,
		Score:         100,
		UUID4:         uuid.New().String(),
		UnixTimestamp: 1597805784,
	}

	data.AfterFind(GetTestDB())
	assert.Equal(t, data.Timestamp, "2020-08-19T02:56:24+00:00")
}
