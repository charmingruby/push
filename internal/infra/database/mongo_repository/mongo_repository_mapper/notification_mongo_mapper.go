package mongo_repository_mapper

import (
	"time"

	"github.com/charmingruby/push/internal/domain/notification/notification_entity"
)

func NotificationToMongo(domainObj *notification_entity.Notification) *MongoNotification {
	return &MongoNotification{
		ID:                     domainObj.ID,
		Destination:            domainObj.Destination,
		Date:                   domainObj.Date,
		Status:                 domainObj.Status,
		CommunicationChannelID: domainObj.CommunicationChannelID,
		Retries:                domainObj.Retries,
		CreatedAt:              domainObj.CreatedAt,
	}
}

func MongoNotificationToDomain(mongoObj *MongoNotification) *notification_entity.Notification {
	return &notification_entity.Notification{
		ID:                     mongoObj.ID,
		Destination:            mongoObj.Destination,
		Date:                   mongoObj.Date,
		Status:                 mongoObj.Status,
		CommunicationChannelID: mongoObj.CommunicationChannelID,
		Retries:                mongoObj.Retries,
		CreatedAt:              mongoObj.CreatedAt,
	}
}

type MongoNotification struct {
	ID                     string    `json:"_id" bson:"_id"`
	Destination            string    `json:"destination" bson:"destination"`
	Date                   time.Time `json:"date" bson:"date"`
	Status                 string    `json:"status" bson:"status"`
	CommunicationChannelID string    `json:"communication_channel_id" bson:"communication_channel_id"`
	Retries                int       `json:"retries" bson:"retries"`
	CreatedAt              time.Time `json:"created_at" bson:"created_at"`
}
