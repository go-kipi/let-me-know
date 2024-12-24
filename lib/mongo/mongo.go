package mongo

import (
	"context"
	"github.com/go-kipi/let-me-know/lib/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MongoI interface {
	GetClient() *mongo.Client
	GetOrder(c context.Context, companyId, projectId, userId string) (MongoOrder, error)
	UpsertOrder(c context.Context, companyId, projectId, userId string, item interface{}) (interface{}, error)
}

type MongoS struct {
	mongoClient *mongo.Client
	conf        *config.Config
}

type MongoOrder struct {
	Id        string
	Order     []interface{}
	Timestamp time.Time
}
type MongoItem struct {
	Id     int
	Extras []int
}

func (m *MongoS) GetClient() *mongo.Client {
	return m.mongoClient
}

func (m *MongoS) GetOrder(c context.Context, companyId, projectId, userId string) (MongoOrder, error) {
	var ordering []MongoOrder
	if result, err := m.mongoClient.Database(companyId).Collection(projectId).Find(c, bson.M{"id": userId}); err != nil {
		return MongoOrder{}, err
	} else {
		err := result.All(c, &ordering)
		if err != nil {
			return MongoOrder{}, err
		}
		if len(ordering) > 0 {
			return ordering[0], nil
		}
		return MongoOrder{}, nil

	}

}

func (m *MongoS) UpsertOrder(c context.Context, companyId, projectId, userId string, item interface{}) (interface{}, error) {
	ordering, err := m.GetOrder(c, companyId, projectId, userId)
	if err != nil {
		return nil, err
	}
	if ordering.Id != "" {
		filter := bson.M{"id": userId}
		ordering.Order = append(ordering.Order, item)
		update := bson.D{{"$set", bson.D{{"order", ordering.Order}}}}

		result, err := m.mongoClient.Database(companyId).Collection(projectId).UpdateOne(c, filter, update)
		if err != nil {
			return 0, err
		}
		return result, nil
	} else {
		result, err := m.mongoClient.Database(companyId).Collection(projectId).InsertOne(c, MongoOrder{
			Id:        userId,
			Timestamp: time.Now(),
			Order:     append(ordering.Order, item),
		})
		if err != nil {
			return 0, err
		}
		return result, nil
	}

}

func NewMongo(config *config.Config) MongoI {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(config.Mongo.Url).
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}
	return &MongoS{mongoClient: client, conf: config}
}
