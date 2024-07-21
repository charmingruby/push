package mongo_repository

import (
	"context"
	"log/slog"
	"time"

	"github.com/charmingruby/push/internal/core"
	"github.com/charmingruby/push/internal/domain/notification/notification_entity"
	"github.com/charmingruby/push/internal/infra/database/mongo_repository/mongo_repository_mapper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewNotificationsMongoRepository(db *mongo.Database) *NotificationsMongoRepository {
	return &NotificationsMongoRepository{
		db: db,
	}
}

type NotificationsMongoRepository struct {
	db *mongo.Database
}

func (r *NotificationsMongoRepository) Store(n *notification_entity.Notification) error {
	doc := mongo_repository_mapper.NotificationToMongo(n)

	collection := r.db.Collection(NOTIFICATION_COLLECTION)

	if _, err := collection.InsertOne(context.Background(), doc); err != nil {
		slog.Error(err.Error())
		return err
	}

	return nil
}

func (r *NotificationsMongoRepository) FindByID(id string) (*notification_entity.Notification, error) {
	filter := bson.D{{Key: "_id", Value: id}}

	collection := r.db.Collection(NOTIFICATION_COLLECTION)

	res := collection.FindOne(context.Background(), filter)
	if res == nil {
		return nil, core.NewNotFoundErr("notification")
	}

	var mongoNotification *mongo_repository_mapper.MongoNotification
	if err := res.Decode(&mongoNotification); err != nil {
		return nil, core.NewInternalErr("notification unmarshal error")
	}

	notification := mongo_repository_mapper.MongoNotificationToDomain(
		mongoNotification,
	)

	return notification, nil
}

func (r *NotificationsMongoRepository) ListAvailableNotificationsBeforeADate(date time.Time) ([]notification_entity.Notification, error) {
	var notifications []notification_entity.Notification

	ctx := context.Background()

	collection := r.db.Collection(NOTIFICATION_COLLECTION)

	cur, err := collection.Find(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var n mongo_repository_mapper.MongoNotification
		err := cur.Decode(&n)
		if err != nil {
			return nil, err
		}

		if n.Date.Before(date) && (n.Status == "PENDING" || n.Status == "RETRYING") {
			domainNotification := mongo_repository_mapper.MongoNotificationToDomain(&n)
			notifications = append(notifications, *domainNotification)
		}
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	cur.Close(ctx)

	return notifications, nil
}

func (r *NotificationsMongoRepository) SaveStatus(n *notification_entity.Notification) error {
	filter := bson.D{{Key: "_id", Value: n.ID}}

	update := bson.M{"$set": bson.M{"retries": n.Retries, "status": n.Status}}

	collection := r.db.Collection(NOTIFICATION_COLLECTION)

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}
