package http

import (
	customTime "github.com/joaquimcalvogubianas/energy-pricing-parser/time"
)

type OmieFileUrlGenerator struct {
	TimeFetcher customTime.TimeFetcher
}

func (ofug OmieFileUrlGenerator) GetCurrentDateFileUrl() string {
	var currentTime = ofug.TimeFetcher.GetNow()
	var currentTimeFormatted = currentTime.Format("20060102")
	return "https://www.omie.es/es/file-download?parents%5B0%5D=marginalpdbc&filename=marginalpdbc_" + currentTimeFormatted + ".1"
}
