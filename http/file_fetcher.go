package http

import (
	"errors"
	"github.com/joaquimcalvogubianas/energy-pricing-parser/di"
	"github.com/joaquimcalvogubianas/energy-pricing-parser/domain"
	"strconv"
	"strings"
	"time"
)

func FetchOmieFile() (string, error) {
	client := di.GetHttpClient()
	url := getCurrentDateFileUrl()
	response, error := client.R().Get(url)
	if (error != nil) {
		return "", error
	}

	responseBody := response.Body()
	stringResponseBody := string(responseBody[:])

	parseResponseBody(stringResponseBody)
}

func getCurrentDateFileUrl() string {
	var currentTime = time.Now()
	var currentTimeFormatted = currentTime.Format("YYYYMMDD")
	return "https://www.omie.es/es/file-download?parents%5B0%5D=marginalpdbc&filename=marginalpdbc_" + currentTimeFormatted + ".1"
}

func parseResponseBody(body string) ([]domain.Price, error) {
	rows := strings.Split(body, "\n")
	parsedPrices := make([]domain.Price, 0)
	for rowIndex, row := range rows {
		rowContent := strings.Split(row, ";")
		rowDate := rowContent[0] + rowContent[1] + rowContent[2]
		priceDate, error := time.Parse("YYYYMMDD", rowDate)
		if error != nil {
			return nil, errors.New("Error tring to parse file price date on row" + strconv.Itoa(rowIndex))
		}

		price, error := strconv.ParseFloat(rowContent[3], 4)
		if error != nil {
			return nil, errors.New("Error trying to parse price from row " + strconv.Itoa(rowIndex))
		}

		newPrice :=  domain.Price{
			priceDate,
			price,
		}

		parsedPrices = append(parsedPrices, newPrice)
	}

	return parsedPrices, nil
}