package notification_repository

import "github.com/charmingruby/push/internal/domain/notification/notification_entity"

type CommunicationChannelRepository interface {
	Store(cc *notification_entity.CommunicationChannel) error
}
