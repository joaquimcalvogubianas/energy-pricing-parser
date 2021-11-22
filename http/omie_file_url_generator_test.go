package http

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type CustomTimeFetcher struct {
	mockedCurrentTime time.Time
}

func (ctf CustomTimeFetcher) GetNow() time.Time {
	return ctf.mockedCurrentTime
}

func TestGetCurrentDateFileUrl_ShouldReturnTodayFileUrl(t *testing.T) {
	time, _ := time.Parse("2006-01-02", "2020-12-13")
	customTimeFetcher := CustomTimeFetcher{
		mockedCurrentTime: time,
	}
	urlGenerator := OmieFileUrlGenerator{
		timeFetcher: customTimeFetcher,
	}

	omieFileUrl := urlGenerator.GetCurrentDateFileUrl()

	assert.Equal(t, "https://www.omie.es/es/file-download?parents%5B0%5D=marginalpdbc&filename=marginalpdbc_20201213.1", omieFileUrl)
}
