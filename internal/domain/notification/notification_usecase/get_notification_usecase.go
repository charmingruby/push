package notification_usecase

import (
	"github.com/charmingruby/push/internal/core"
	"github.com/charmingruby/push/internal/domain/notification/notification_dto"
	"github.com/charmingruby/push/internal/domain/notification/notification_entity"
)

func (r *NotificationUseCaseRegistry) GetNotificationUseCase(
	dto notification_dto.GetNotificationDTO,
) (*notification_entity.Notification, error) {
	notification, err := r.notificationRepo.FindByID(dto.NotificationID)
	if err != nil {
		return nil, core.NewNotFoundErr("notification")
	}

	return notification, nil
}
