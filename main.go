package main

import (
	"github.com/joaquimcalvogubianas/energy-pricing-parser/service"
)

func main() {
	println("Started price importation process")
	service.CleanPriceOnMongoDatabase()
	service.LoadPricesToMongoDatabase()
	println("Finished price importation process")
}
