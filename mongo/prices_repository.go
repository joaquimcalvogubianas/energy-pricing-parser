package mongo

import (
	"context"
	"fmt"
	"github.com/joaquimcalvogubianas/energy-pricing-parser/domain"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type MongoDbPricesRepository struct {
	MongoClient    *mongo.Client
	CancelFunction context.CancelFunc
}

func (databaseRepository MongoDbPricesRepository) PersistPrices(prices []domain.Price) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Erro trying to persist prices in mongo database", r)
		}
		databaseRepository.CloseDatabaseConnection()
	}()

	perPricesCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := databaseRepository.MongoClient.Database("pricing-app").Collection("prices")
	convertedPrices := []interface{}{}
	//TODO: what the fuck is this shit
	for _, price := range prices {
		convertedPrices = append(convertedPrices, price)
	}

	insRes, insErr := collection.InsertMany(perPricesCtx, convertedPrices)
	println("Inserted content in mongo database")
	if insErr != nil {
		println(insRes)
	}
}

func (databaseRepository MongoDbPricesRepository) CloseDatabaseConnection() {
	databaseRepository.CancelFunction()
}
