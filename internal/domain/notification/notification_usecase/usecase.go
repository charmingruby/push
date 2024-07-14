package notification_usecase

import (
	"github.com/charmingruby/push/internal/domain/notification/notification_adapter"
	"github.com/charmingruby/push/internal/domain/notification/notification_dto"
	"github.com/charmingruby/push/internal/domain/notification/notification_entity"
	"github.com/charmingruby/push/internal/domain/notification/notification_repository"
)

type NotificationServiceUseCase interface {
	CreateCommunicationChannelUseCase(dto notification_dto.CreateCommunicationChannelDTO) error
	ScheduleNotificationUseCase(dto notification_dto.ScheduleNotificationDTO) error
	GetNotificationUseCase(dto notification_dto.GetNotificationDTO) (*notification_entity.Notification, error)
	CancelNotiticationUseCase(dto notification_dto.CancelNotificationDTO) error
	CheckAndSendNotificationUseCase() ([]notification_entity.Notification, error)
}

func NewNotificationUseCaseRegistry(
	notificationRepo notification_repository.NotificationRepository,
	communicationChannelRepo notification_repository.CommunicationChannelRepository,
	dispatcher notification_adapter.Dispatcher,
) *NotificationUseCaseRegistry {
	return &NotificationUseCaseRegistry{
		notificationRepo:         notificationRepo,
		communicationChannelRepo: communicationChannelRepo,
		dispatcher:               dispatcher,
	}
}

type NotificationUseCaseRegistry struct {
	notificationRepo         notification_repository.NotificationRepository
	communicationChannelRepo notification_repository.CommunicationChannelRepository
	dispatcher               notification_adapter.Dispatcher
}
