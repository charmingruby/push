package notification_usecase

import (
	"github.com/charmingruby/push/internal/core"
	"github.com/charmingruby/push/internal/domain/notification/notification_dto"
	"github.com/charmingruby/push/internal/domain/notification/notification_entity"
)

func (r *NotificationUseCaseRegistry) ScheduleNotificationUseCase(dto notification_dto.ScheduleNotificationDTO) error {
	if _, err := r.communicationChannelRepo.FindByID(dto.CommunicationChannelID); err != nil {
		return core.NewNotFoundErr("communication channel")
	}

	n, err := notification_entity.NewNotification(
		dto.Destination,
		dto.RawDate,
		dto.CommunicationChannelID,
	)
	if err != nil {
		return err
	}

	if err := r.notificationRepo.Store(n); err != nil {
		return core.NewInternalErr("schedule notification use case: store notification")
	}

	return nil
}
