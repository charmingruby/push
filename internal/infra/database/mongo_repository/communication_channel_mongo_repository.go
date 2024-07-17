package mongo_repository

import (
	"context"
	"log/slog"

	"github.com/charmingruby/push/internal/core"
	"github.com/charmingruby/push/internal/domain/notification/notification_entity"
	"github.com/charmingruby/push/internal/infra/database/mongo_repository/mongo_repository_mapper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewCommunicationChannelMongoRepository(db *mongo.Database) *CommunicationChannelMongoRepository {
	return &CommunicationChannelMongoRepository{
		db: db,
	}
}

type CommunicationChannelMongoRepository struct {
	db *mongo.Database
}

func (r *CommunicationChannelMongoRepository) Store(cc *notification_entity.CommunicationChannel) error {
	doc := mongo_repository_mapper.CommunicationChannelToMongo(cc)

	collection := r.db.Collection(COMMUNICATION_CHANNEL_COLLECTION)

	if _, err := collection.InsertOne(context.Background(), doc); err != nil {
		slog.Error(err.Error())
		return err
	}

	return nil
}

func (r *CommunicationChannelMongoRepository) FindByName(name string) (*notification_entity.CommunicationChannel, error) {
	filter := bson.D{{Key: "name", Value: name}}

	collection := r.db.Collection(COMMUNICATION_CHANNEL_COLLECTION)

	res := collection.FindOne(context.Background(), filter)
	if res == nil {
		return nil, core.NewNotFoundErr("communication channel")
	}

	var mongoCommunicationChannel *mongo_repository_mapper.MongoCommunicationChannel
	if err := res.Decode(&mongoCommunicationChannel); err != nil {
		return nil, core.NewInternalErr("communication channel unmarshal error")
	}

	notification := mongo_repository_mapper.MongoCommunicationChannelToDomain(
		mongoCommunicationChannel,
	)

	return notification, nil
}

func (r *CommunicationChannelMongoRepository) FindByID(id string) (*notification_entity.CommunicationChannel, error) {
	filter := bson.D{{Key: "id", Value: id}}

	collection := r.db.Collection(COMMUNICATION_CHANNEL_COLLECTION)

	res := collection.FindOne(context.Background(), filter)
	if res == nil {
		return nil, core.NewNotFoundErr("communication channel")
	}

	var mongoCommunicationChannel *mongo_repository_mapper.MongoCommunicationChannel
	if err := res.Decode(&mongoCommunicationChannel); err != nil {
		return nil, core.NewInternalErr("communication channel unmarshal error")
	}

	notification := mongo_repository_mapper.MongoCommunicationChannelToDomain(
		mongoCommunicationChannel,
	)

	return notification, nil

}
