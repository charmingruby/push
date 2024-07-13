package notification_usecase

import (
	"github.com/charmingruby/push/internal/domain/notification/notification_dto"
	"github.com/charmingruby/push/internal/domain/notification/notification_repository"
)

type NotificationServiceUseCase interface {
	CreateCommunicationChannelUseCase(dto notification_dto.CreateCommunicationChannelDTO) error
	ScheduleNotificationUseCase()
	GetNotificationUseCase()
	CancelNotiticationUseCase()
	CheckAndSendNotificationUseCase()
}

func NewNotificationUseCaseRegistry(
	notificationRepo notification_repository.NotificationRepository,
	communicationChannelRepo notification_repository.CommunicationChannelRepository,
) *NotificationUseCaseRegistry {
	return &NotificationUseCaseRegistry{
		notificationRepo:         notificationRepo,
		communicationChannelRepo: communicationChannelRepo,
	}
}

type NotificationUseCaseRegistry struct {
	notificationRepo         notification_repository.NotificationRepository
	communicationChannelRepo notification_repository.CommunicationChannelRepository
}
