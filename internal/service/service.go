package service

import (
	"junior/internal/model"
	"junior/internal/repository"
	"time"
)

type SubscriptionService struct {
	repo *repository.SubscriptionRepository
}

func NewSubscriptionService(r *repository.SubscriptionRepository) *SubscriptionService {
	return &SubscriptionService{repo: r}
}

func (s *SubscriptionService) Create(sub *model.Subscription) error {
	return s.repo.Create(sub)
}

func (s *SubscriptionService) GetAll() ([]model.Subscription, error) {
	return s.repo.GetAll()
}

func (s *SubscriptionService) GetByID(id uint) (*model.Subscription, error) {
	return s.repo.GetByID(id)
}

func (s *SubscriptionService) Update(sub *model.Subscription) error {
	return s.repo.Update(sub)
}

func (s *SubscriptionService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *SubscriptionService) CalculateTotal(from, to time.Time, userID, serviceName string) (int, error) {
	return s.repo.CalculateTotal(from, to, userID, serviceName)
}
