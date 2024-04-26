package main

import (
	"context"
	"log"
	"net/http"

	"github.com/Nic-vwalt/raha-coding-challenge/internal/app"
	"github.com/labstack/echo/v4"
)

// Declare global variables to hold reusable resources
var (

    registry  *app.ServiceRegistry
    e         *echo.Echo
)

func init() {
    e = echo.New()
}

func setup() {
    var err error
    // Lazy initialization of the service registry
    if registry == nil {
        registry, err = app.AddServices(context.Background())
        if err != nil {
            log.Fatalf("Failed to initialize services: %v", err)
        }
        log.Println("Services Initialized!")
    }

    // Setup routes
    setupRoutes(e, registry)
}

func setupRoutes(e *echo.Echo, registry *app.ServiceRegistry) {
    // Update AddRoutes to work with Echo context
    app.AddRoutes(e, registry)
    log.Println("Routes Setup!")
}

func startServer(e *echo.Echo) {
    log.Println("Starting server on :4000")
    if err := e.Start("localhost:4000"); err != http.ErrServerClosed {
        log.Fatal("Server start failed: ", err)
    }
}

func main() {
    // Setup the application on each invocation to check if the global vars are initialized
    setup()

    // Start the server
    startServer(e)
}
