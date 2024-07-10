package notification_usecase

type NotificationServiceContract interface {
	CreateCommunicationChannelUseCase()
	ScheduleNotificationUseCase()
	GetNotificationUseCase()
	CancelNotiticationUseCase()
	CheckAndSendNotificationUseCase()
}
