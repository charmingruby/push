package notification_usecase

import (
	"github.com/charmingruby/push/internal/core"
	"github.com/charmingruby/push/internal/domain/notification/notification_dto"
	"github.com/charmingruby/push/internal/domain/notification/notification_entity"
)

func (r *NotificationUseCaseRegistry) CreateCommunicationChannelUseCase(dto notification_dto.CreateCommunicationChannelDTO) error {
	_, err := r.communicationChannelRepo.FindByName(dto.Name)
	if err == nil {
		return core.NewConflictErr("communication channel", "name")
	}

	cc, err := notification_entity.NewCommunicationChannel(dto.Name, dto.Description)
	if err != nil {
		return err
	}

	if err := r.communicationChannelRepo.Store(cc); err != nil {
		return core.NewInternalErr("create communication use case: store communication channel")
	}

	return nil
}
