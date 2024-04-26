package app

import (
	"context"
	"log"

	"github.com/Nic-vwalt/raha-coding-challenge/internal/service"
)

func AddServices(ctx context.Context) (*ServiceRegistry, error) {
    itineraryService, err := service.NewItineraryService(ctx)
    if err != nil {
        log.Printf("Failed to initialize auth service: %v", err)
        return nil, err
    }

    // Initialize and return the service registry
    registry := &ServiceRegistry{
        ItineraryService:     itineraryService,
    }

    return registry, nil
}
