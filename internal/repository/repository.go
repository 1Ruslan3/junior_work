package repository

import (
	"junior/internal/model"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SubscriptionRepository interface {
	Create(subscription *model.Subscription) error
	GetAll() ([]model.Subscription, error)
	GetByID(id uint) (*model.Subscription, error)
	Update(subscription *model.Subscription) error
	Delete(id uint) error
	CalculateTotal(from, to time.Time, userID *uuid.UUID, serviceName *string) (int64, error)
}

type subscriptionRepo struct {
	db *gorm.DB
}

func NewSubscriptionRepository(db *gorm.DB) SubscriptionRepository {
	return &subscriptionRepo{db: db}
}

func (r *subscriptionRepo) Create(s *model.Subscription) error {
	return r.db.Create(s).Error
}

func (r *subscriptionRepo) GetAll() ([]model.Subscription, error) {
	var subs []model.Subscription
	err := r.db.Find(&subs).Error
	return subs, err
}

func (r *subscriptionRepo) GetByID(id uint) (*model.Subscription, error) {
	var sub model.Subscription
	err := r.db.First(&sub, id).Error
	return &sub, err
}

func (r *subscriptionRepo) Update(s *model.Subscription) error {
	return r.db.Save(s).Error
}

func (r *subscriptionRepo) Delete(id uint) error {
	return r.db.Delete(&model.Subscription{}, id).Error
}

func (r *subscriptionRepo) CalculateTotal(from, to time.Time, userID *uuid.UUID, serviceName *string) (int64, error) {
	var total int64
	query := r.db.Model(&model.Subscription{}).
		Where("start_date <= ? AND (end_date IS NULL OR end_date >= ?)", to, from)

	if userID != nil {
		query = query.Where("user_id = ?", *userID)
	}
	if serviceName != nil {
		query = query.Where("service_name = ?", *serviceName)
	}

	err := query.Select("SUM(price)").Scan(&total).Error
	return total, err
}
