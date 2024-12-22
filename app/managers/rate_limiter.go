package managers

import (
	"fmt"
	"gorm.io/gorm"
	"noisy_neighbour/app/models"
	"noisy_neighbour/common/database/base_model"
	"time"
)

type RateLimiterManager struct {
	base_model.BaseManager
	rateLimiter models.RateLimiter
}

func (m *RateLimiterManager) Take() (bool, error) {
	m.refill()
	if m.rateLimiter.Tokens <= 0 {
		return false, nil
	}

	m.rateLimiter.Tokens--

	err := m.Save(m.rateLimiter)

	if err != nil {
		return false, fmt.Errorf("failed to save rate limiter: %w", err)
	}

	return true, nil
}

func (m *RateLimiterManager) refill() {
	elapsed := time.Since(m.rateLimiter.LastCheck)
	m.rateLimiter.LastCheck = time.Now()
	m.rateLimiter.Tokens += int(elapsed.Seconds()) * m.rateLimiter.Rate
	if m.rateLimiter.Tokens > m.rateLimiter.Capacity {
		m.rateLimiter.Tokens = m.rateLimiter.Capacity
	}
}

func NewRateLimiterManager(db *gorm.DB, rateLimiter models.RateLimiter) *RateLimiterManager {
	return &RateLimiterManager{
		BaseManager: *base_model.NewBaseManager(db),
		rateLimiter: rateLimiter,
	}
}
