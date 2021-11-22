package service

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/joaquimcalvogubianas/energy-pricing-parser/di"
	"github.com/joaquimcalvogubianas/energy-pricing-parser/http"
	"github.com/joaquimcalvogubianas/energy-pricing-parser/mongo"
	"os"
)

type PriceFetcherService struct{}

func CleanPriceOnMongoDatabase() {
	mongoDbClient, cancelFunction, error := di.GetMongoDbClient()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error trying to persist prices in mongo database", r)
		}
		cancelFunction()
	}()

	if error != nil {
		panic("Failure trying to connect to mongodatabase")
	}

	err := mongoDbClient.Database("").Drop(nil)
	print("Erased all pricing content on MongoDatabase")
	if err != nil {
		fmt.Println()
	}
}

func LoadPricesToMongoDatabase() {
	client := http.RestyHttpClient{
		Client: resty.New(),
	}

	prices, error := http.FetchOmieFilePrices(client)
	if error != nil {
		os.Exit(5)
	}

	fmt.Printf("There are going to be inserted %d prices\n", len(prices))

	mongoDbClient, cancelFunction, error := di.GetMongoDbClient()
	if error != nil {
		panic("Failure trying to connect to mongodatabase")
	}

	pricesRepository := mongo.MongoDbPricesRepository{
		MongoClient:    mongoDbClient,
		CancelFunction: cancelFunction,
	}

	pricesRepository.PersistPrices(prices)
}
