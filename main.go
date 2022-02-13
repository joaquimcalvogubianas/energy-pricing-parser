package main

import (
	"fmt"
	"github.com/joaquimcalvogubianas/energy-pricing-parser/di"
	"github.com/joaquimcalvogubianas/energy-pricing-parser/service"
)

func main() {
	err := di.DiContext.Invoke(func(pf service.PriceFetcherService) error {
		println("Started price importation process")
		pf.CleanPriceOnMongoDatabase()
		pf.LoadPricesToMongoDatabase()
		println("Finished price importation process")
		return nil
	})

	if err != nil {
		fmt.Printf("error trying to run application %v", err)
	}
}
