package http

import (
	"errors"
	"strconv"
	"strings"
	goTime "time"

	"github.com/go-resty/resty/v2"
	"github.com/joaquimcalvogubianas/energy-pricing-parser/domain"
	"github.com/joaquimcalvogubianas/energy-pricing-parser/time"
)

func FetchOmieFilePrices(client HttpClient) ([]domain.Price, error) {
	fileUrlGenerator := OmieFileUrlGenerator{
		timeFetcher: time.LocalTimeFetcher{},
	}
	url := fileUrlGenerator.GetCurrentDateFileUrl()
	request := client.Request()
	response, error := request.Get(url)
	if error != nil {
		return nil, error
	}

	responseBody := response.Body()
	stringResponseBody := string(responseBody[:])

	return parseResponseBody(stringResponseBody)
}

func parseResponseBody(body string) (prices []domain.Price, error error) {
	defer func() {
		if r := recover(); r != nil {
			prices = nil
			error = errors.New("Errors trying to parse incomming file")
		}
	}()

	rows := strings.Split(body, "\n")
	rows = cleanRows(rows)
	parsedPrices := make([]domain.Price, 0)
	for rowIndex, row := range rows {
		if rowIndex != 0 {
			rowContent := strings.Split(row, ";")
			rowDate := rowContent[0] + rowContent[1] + rowContent[2] + rowContent[3]
			priceDate, error := goTime.Parse("2006010215", rowDate)
			if error != nil {
				return nil, errors.New("Error tring to parse file price date on row " + strconv.Itoa(rowIndex))
			}

			price, error := strconv.ParseFloat(rowContent[3], 4)
			if error != nil {
				return nil, errors.New("Error trying to parse price from row " + strconv.Itoa(rowIndex))
			}

			newPrice := domain.Price{
				Date:  priceDate,
				Price: price,
			}

			parsedPrices = append(parsedPrices, newPrice)
		}
	}

	return parsedPrices, nil
}

func cleanRows(rows rows) rows {
	var newRows = rows.removeFromIndex(0)
	newRows = newRows.removeFromIndex(24)
	newRows = newRows.removeFromIndex(23)
	return newRows.removeFromIndex(23)
}

type rows []string

func (r rows) removeFromIndex(index int) rows {
	copy(r[index:], r[index+1:])
	r[len(r)-1] = ""
	return r[:len(r)-1]
}

type RestyHttpClient struct {
	Client *resty.Client
}

type RestyHttpRequest struct {
	request *resty.Request
}

type RestyHttpResponse struct {
	response *resty.Response
	error    error
}

func (c RestyHttpClient) Request() HttpRequest {
	return RestyHttpRequest{
		request: c.Client.R(),
	}
}

func (r RestyHttpRequest) Get(url string) (RawHttpResponse, error) {
	restyResponse, error := r.request.Get(url)
	return RestyHttpResponse{
		response: restyResponse,
		error:    error,
	}, nil
}

func (rr RestyHttpResponse) Body() []byte {
	return rr.response.Body()
}
