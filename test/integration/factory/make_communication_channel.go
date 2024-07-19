package factory

import (
	"github.com/charmingruby/push/internal/domain/notification/notification_entity"
	"github.com/charmingruby/push/internal/domain/notification/notification_repository"
)

func MakeCommunicationChannel(
	repo notification_repository.CommunicationChannelRepository,
	name,
	description string,
) (*notification_entity.CommunicationChannel, error) {
	cc, err := notification_entity.NewCommunicationChannel(
		name,
		description,
	)
	if err != nil {
		return nil, err
	}

	if err := repo.Store(cc); err != nil {
		return nil, err
	}

	return cc, nil
}
