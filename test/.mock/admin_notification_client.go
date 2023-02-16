package mock

import (
	"github.com/shonjord/tracker/internal/pkg/domain/entity"
)

type (
	AdminNotificationClient struct {
		NotifyFunc func(*entity.Computer) error
	}
)

// Notify refer to the consumer of the interface for documentation.
func (m *AdminNotificationClient) Notify(computer *entity.Computer) error {
	return m.NotifyFunc(computer)
}
