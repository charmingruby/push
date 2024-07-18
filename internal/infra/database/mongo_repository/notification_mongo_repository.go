package mongo_repository

import (
	"time"

	"github.com/charmingruby/push/internal/domain/notification/notification_entity"
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
	return nil
}

func (r *NotificationsMongoRepository) GetNotificationByID(id string) (*notification_entity.Notification, error) {
	return nil, nil
}

func (r *NotificationsMongoRepository) ListAvailableNotificationsBeforeDate(date time.Time) ([]notification_entity.Notification, error) {
	return nil, nil

}

func (r *NotificationsMongoRepository) SaveNotificationStatus(n *notification_entity.Notification) error {
	return nil
}
