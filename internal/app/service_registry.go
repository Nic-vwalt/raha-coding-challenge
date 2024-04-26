package app

import (
	"context"

	"github.com/Nic-vwalt/raha-coding-challenge/internal/service"
)

// ServiceRegistry holds all the services required by the application.
type ServiceRegistry struct {
    ItineraryService *service.ItineraryService
}

// NewServiceRegistry creates a new ServiceRegistry with the required services initialized.
func NewServiceRegistry(ctx context.Context) (*ServiceRegistry, error) {
    itinService, err := service.NewItineraryService(ctx)
    if err != nil {
        return nil, err
    }

    return &ServiceRegistry{
        ItineraryService: itinService,
    }, nil
}
