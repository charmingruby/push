package mongo_repository_mapper

import (
	"time"

	"github.com/charmingruby/push/internal/domain/notification/notification_entity"
)

func CommunicationChannelToMongo(domainObj *notification_entity.CommunicationChannel) *MongoCommunicationChannel {
	return &MongoCommunicationChannel{
		ID:          domainObj.ID,
		Name:        domainObj.Name,
		Description: domainObj.Description,
		CreatedAt:   domainObj.CreatedAt,
	}
}

func MongoCommunicationChannelToDomain(mongoObj *MongoCommunicationChannel) *notification_entity.CommunicationChannel {
	return &notification_entity.CommunicationChannel{
		ID:          mongoObj.ID,
		Name:        mongoObj.Name,
		Description: mongoObj.Description,
		CreatedAt:   mongoObj.CreatedAt,
	}
}

type MongoCommunicationChannel struct {
	ID          string    `json:"_id" bson:"_id"`
	Name        string    `json:"name" bson:"name"`
	Description string    `json:"description" bson:"description"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
}
