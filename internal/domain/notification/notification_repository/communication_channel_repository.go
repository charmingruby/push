package notification_repository

import "github.com/charmingruby/push/internal/domain/notification/notification_entity"

type CommunicationChannelRepository interface {
	Store(cc *notification_entity.CommunicationChannel) error
	FindByName(name string) (*notification_entity.CommunicationChannel, error)
	FindByID(id string) (*notification_entity.CommunicationChannel, error)
}
