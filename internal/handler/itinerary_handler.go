package handler

import (
	"net/http"

	"github.com/Nic-vwalt/raha-coding-challenge/internal/service"
	"github.com/labstack/echo/v4"
)

type ItineraryHandler struct {
    itineraryService *service.ItineraryService
}

func NewItinHandler(itineraryService *service.ItineraryService) *ItineraryHandler {
    return &ItineraryHandler{
        itineraryService: itineraryService,
    }
}

// TransformItinerary is an Echo handler that reads a request, processes it, and returns a JSON response.
func (h *ItineraryHandler) ReconstructItinerary(c echo.Context) error {
    var routePairs [][]string
    if err := c.Bind(&routePairs); err != nil {
        return c.String(http.StatusBadRequest, "Invalid request body")
    }

    result, err := h.itineraryService.ReconstructItinerary(c.Request().Context(), routePairs)
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Failed to process itineraries")
    }

    return c.JSON(http.StatusOK, result)
}
