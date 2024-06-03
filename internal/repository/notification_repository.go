package repository

import (
	"Ticketing/entity"
	"context"

	"gorm.io/gorm"
)

type NotificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) *NotificationRepository {
	return &NotificationRepository{
		db: db,
	}
}

// get all notification
func (r *NotificationRepository) GetAllNotification(ctx context.Context) ([]*entity.Notification, error) {
	Notifications := make([]*entity.Notification, 0)
	result := r.db.WithContext(ctx).Find(&Notifications)
	if result.Error != nil {
		return nil, result.Error
	}
	return Notifications, nil
}

// create notification
func (r *NotificationRepository) CreateNotification(ctx context.Context, Notification *entity.Notification) error {
	result := r.db.WithContext(ctx).Create(&Notification)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// get notification after get change value isRead to true and only get notification if is_read false UserGetNotification
func (r *NotificationRepository) UserGetNotification(ctx context.Context) ([]*entity.Notification, error) {
	Notifications := make([]*entity.Notification, 0)

	// Retrieve notifications with is_read = false
	result := r.db.WithContext(ctx).Where("is_read = ?", false).Find(&Notifications)
	if result.Error != nil {
		return nil, result.Error
	}

	// Mark retrieved notifications as read
	for _, notification := range Notifications {
		// Assuming you have a method to update the is_read field
		err := r.MarkNotificationAsRead(ctx, notification.ID)
		if err != nil {
			return nil, err
		}
	}

	return Notifications, nil
}

func (r *NotificationRepository) MarkNotificationAsRead(ctx context.Context, notificationID int) error {
	result := r.db.WithContext(ctx).Model(&entity.Notification{}).Where("id = ?", notificationID).Update("is_read", true)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
