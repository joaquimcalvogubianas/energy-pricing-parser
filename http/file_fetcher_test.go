package http

import (
	"fmt"
	functional_extensions "github.com/joaquimcalvogubianas/energy-pricing-parser/functional-extensions"
	resources_manager "github.com/joaquimcalvogubianas/energy-pricing-parser/resources-manager"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFetchOmieFilePrices_should_returnTodayPrices(t *testing.T) {

	rawResponse := MockRawHttpResponse{
		RawResponse: resources_manager.GetTestResource("mocks/pricesFileMock.1"),
	}

	request := MockHttpRequest{
		Response: rawResponse,
	}

	client := MockHttpCient{
		Req: request,
	}

	fileFetcher := FileFetcher{
		Client: &client,
	}
	prices, err := fileFetcher.FetchOmieFilePrices()

	if err != nil {
		t.Error(fmt.Sprintf("Test file can not be parsed, cause: %q", err.Error()))
	}

	firstHourPrice := functional_extensions.PricesList(prices).First()
	firstHourPriceDate := firstHourPrice.Date
	assert.Equal(t, 2022, firstHourPriceDate.Year())
	assert.Equal(t, time.January, firstHourPriceDate.Month())
	assert.Equal(t, 30, firstHourPriceDate.Day())
	assert.Equal(t, 2, firstHourPriceDate.Hour())
	assert.Equal(t, 0, firstHourPriceDate.Minute())
	assert.Equal(t, 0, firstHourPriceDate.Second())
}
