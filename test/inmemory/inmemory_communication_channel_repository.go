package inmemory

import (
	"github.com/charmingruby/push/internal/core"
	"github.com/charmingruby/push/internal/domain/notification/notification_entity"
)

func NewInMemoryCommunicationChannelRepository() *InMemoryCommunicationChannelRepository {
	return &InMemoryCommunicationChannelRepository{
		Items: []notification_entity.CommunicationChannel{},
	}
}

type InMemoryCommunicationChannelRepository struct {
	Items []notification_entity.CommunicationChannel
}

func (r *InMemoryCommunicationChannelRepository) Store(e *notification_entity.CommunicationChannel) error {
	r.Items = append(r.Items, *e)
	return nil
}

func (r *InMemoryCommunicationChannelRepository) FindByName(name string) (*notification_entity.CommunicationChannel, error) {
	for _, e := range r.Items {
		if e.Name == name {
			return &e, nil
		}
	}

	return nil, core.NewNotFoundErr("communication channel")
}
