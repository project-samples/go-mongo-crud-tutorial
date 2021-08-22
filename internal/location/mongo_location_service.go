package location

import (
	"reflect"

	"go.mongodb.org/mongo-driver/mongo"

	m "github.com/core-go/mongo"
	"github.com/core-go/mongo/query"
	"github.com/core-go/search"
	"github.com/core-go/service"
)

type MongoLocationService struct {
	search.SearchService
	service.GenericService
	Mapper m.Mapper
}

func NewLocationService(db *mongo.Database) *MongoLocationService {
	var model Location
	modelType := reflect.TypeOf(model)
	mapper := m.NewMapper(modelType)
	queryBuilder := query.NewBuilder(modelType)
	searchService, genericService := m.NewSearchWriter(db, "location", modelType, queryBuilder.BuildQuery, search.GetSort, mapper)
	return &MongoLocationService{SearchService: searchService, GenericService: genericService, Mapper: mapper}
}
