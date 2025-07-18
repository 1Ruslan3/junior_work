package repository

import (
	"fmt"
	"junior/internal/config"
	"junior/internal/model"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type SubscriptionRepository struct {
	db *gorm.DB
}

func InitDB(cfg config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.DB.Host, cfg.DB.User, cfg.DB.Password, cfg.DB.Name, cfg.DB.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func NewSubscriptionRepository(db *gorm.DB) *SubscriptionRepository {
	return &SubscriptionRepository{db: db}
}

func (r *SubscriptionRepository) Create(sub *model.Subscription) error {
	return r.db.Create(sub).Error
}

func (r *SubscriptionRepository) GetAll() ([]model.Subscription, error) {
	var subs []model.Subscription
	err := r.db.Find(&subs).Error
	return subs, err
}

func (r *SubscriptionRepository) GetByID(id uint) (*model.Subscription, error) {
	var sub model.Subscription
	err := r.db.First(&sub, id).Error
	return &sub, err
}

func (r *SubscriptionRepository) Update(sub *model.Subscription) error {
	return r.db.Save(sub).Error
}

func (r *SubscriptionRepository) Delete(id uint) error {
	return r.db.Delete(&model.Subscription{}, id).Error
}

func (r *SubscriptionRepository) CalculateTotal(from, to time.Time, userID, serviceName string) (int, error) {
	var total int64
	query := r.db.Model(&model.Subscription{}).
		Where("start_date >= ? AND (end_date IS NULL OR end_date <= ?)", from, to)
	if userID != "" {
		query = query.Where("user_id = ?", userID)
	}
	if serviceName != "" {
		query = query.Where("service_name = ?", serviceName)
	}
	if err := query.Select("SUM(price)").Scan(&total).Error; err != nil {
		return 0, err
	}
	return int(total), nil
}
