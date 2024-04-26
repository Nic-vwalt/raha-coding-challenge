package app

import (
	"github.com/Nic-vwalt/raha-coding-challenge/internal/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func AddRoutes(e *echo.Echo, registry *ServiceRegistry) {
    // Base group for all routes
    baseGroup := e.Group("")

    itinHandler := handler.NewItinHandler(registry.ItineraryService)

    // Middleware for logging
    baseGroup.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
        Format: "method=${method}, uri=${uri}, status=${status}\n",
    }))

    // Itinerary routes nested under the base path
    itinGroup := baseGroup.Group("/itinerary")
    {   
        itinGroup.POST("/reconstruct", itinHandler.ReconstructItinerary)
    }
}
