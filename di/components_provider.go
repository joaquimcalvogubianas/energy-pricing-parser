package di

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/joaquimcalvogubianas/energy-pricing-parser/http"
	mongo2 "github.com/joaquimcalvogubianas/energy-pricing-parser/mongo"
	"github.com/joaquimcalvogubianas/energy-pricing-parser/service"
	time2 "github.com/joaquimcalvogubianas/energy-pricing-parser/time"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/dig"
)

var DiContext *dig.Container

func init() {
	DiContext = dig.New()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:7017")
	mongoClient, error := mongo.Connect(ctx, clientOptions)

	if error != nil {
		panic("Error trying to connect to mongo database")
	}

	DiContext.Provide(func() *mongo.Client {
		return mongoClient
	})
	DiContext.Provide(func() context.CancelFunc {
		return cancel
	})
	DiContext.Provide(func() http.HttpClient {
		return http.RestyHttpClient{
			Client: resty.New(),
		}
	})
	DiContext.Provide(func(timeFetcher time2.TimeFetcher) http.OmieFileUrlGenerator {
		return http.OmieFileUrlGenerator{}
	})
	DiContext.Provide(func(http2 http.HttpClient) http.FileFetcher {
		return http.FileFetcher{
			Client: http2,
		}
	})
	DiContext.Provide(func(mc *mongo.Client, cf context.CancelFunc) mongo2.MongoDbPricesRepository {
		return mongo2.MongoDbPricesRepository{
			MongoClient:    mc,
			CancelFunction: cf,
		}
	})
	DiContext.Provide(func(
		mc *mongo.Client,
		pr mongo2.MongoDbPricesRepository,
		cf context.CancelFunc,
		pf http.FileFetcher,
	) service.PriceFetcherService {
		return service.PriceFetcherService{
			MongoClient:       mc,
			PricesRepository:  pr,
			CancelFunction:    cf,
			FilePricesFetcher: pf,
		}
	})
}
