package service

import (
	"context"
	"fmt"
	"github.com/joaquimcalvogubianas/energy-pricing-parser/http"
	"github.com/joaquimcalvogubianas/energy-pricing-parser/mongo"
	mongo2 "go.mongodb.org/mongo-driver/mongo"
	"os"
)

type PriceFetcherService struct {
	MongoClient       *mongo2.Client
	CancelFunction    context.CancelFunc
	FilePricesFetcher http.FileFetcher
	PricesRepository  mongo.MongoDbPricesRepository
}

func (pfs PriceFetcherService) CleanPriceOnMongoDatabase() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error trying to drop all current prices from mongo", r)
		}
		pfs.CancelFunction()
	}()

	err := pfs.MongoClient.Database("pricing-app").Drop(nil)
	print("Erased all pricing content on MongoDatabase")
	if err != nil {
		fmt.Println()
	}
}

func (pfs PriceFetcherService) LoadPricesToMongoDatabase() {
	prices, error := pfs.FilePricesFetcher.FetchOmieFilePrices()
	if error != nil {
		os.Exit(5)
	}

	fmt.Printf("There are going to be inserted %d prices\n", len(prices))

	pfs.PricesRepository.PersistPrices(prices)
}
