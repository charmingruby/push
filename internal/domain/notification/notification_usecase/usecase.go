package notification_usecase

import "github.com/charmingruby/push/internal/domain/notification/notification_repository"

type NotificationServiceUseCase interface {
	CreateCommunicationChannelUseCase()
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
