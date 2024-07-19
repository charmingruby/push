package notification_usecase

import (
	"github.com/charmingruby/push/internal/core"
	"github.com/charmingruby/push/internal/domain/notification/notification_dto"
)

func (r *NotificationUseCaseRegistry) CancelNotiticationUseCase(
	dto notification_dto.CancelNotificationDTO,
) error {
	notification, err := r.notificationRepo.FindByID(dto.NotificationID)
	if err != nil {
		return core.NewNotFoundErr("notification")
	}

	if notification.Status == "SENT" {
		return core.NewValidationErr("notification is already sent")
	}

	notification.StatusCanceled()

	if err := r.notificationRepo.SaveStatus(notification); err != nil {
		return core.NewInternalErr("cancel notification use case: save notification status")
	}

	return nil
}
