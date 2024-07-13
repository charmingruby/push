package notification_usecase

import (
	"github.com/charmingruby/push/internal/domain/notification/notification_dto"
	"github.com/charmingruby/push/internal/domain/notification/notification_entity"
	"github.com/charmingruby/push/internal/domain/notification/notification_repository"
)

type NotificationServiceUseCase interface {
	CreateCommunicationChannelUseCase(dto notification_dto.CreateCommunicationChannelDTO) error
	ScheduleNotificationUseCase(dto notification_dto.ScheduleNotificationDTO) error
	GetNotificationUseCase(dto notification_dto.GetNotificationDTO) (*notification_entity.Notification, error)
	CancelNotiticationUseCase(dto notification_dto.CancelNotificationDTO) error
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
