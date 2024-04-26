package service

import (
	"context"
	"fmt"
	"sort"
)

// ItineraryService provides methods for handling itineraries.
type ItineraryService struct {
}

// NewItineraryService creates a new instance of ItineraryService.
func NewItineraryService(ctx context.Context) (*ItineraryService, error) {
    return &ItineraryService{}, nil
}

func (is *ItineraryService) ReconstructItinerary(ctx context.Context, tickets [][]string) ([]string, error) {
	if len(tickets) == 0 {
		return nil, fmt.Errorf("no itineraries provided")
	}

	graph := make(map[string][]string)
	var start string = "JFK" // Default starting point if present in tickets

	for _, ticket := range tickets {
		source, dest := ticket[0], ticket[1]
		graph[source] = append(graph[source], dest)
		if source == "JFK" {
			start = source // Set start if JFK is indeed a starting point in tickets
		}
	}

	for key := range graph {
		sort.Strings(graph[key]) // Sort destinations to ensure lexical order
	}

	if _, exists := graph[start]; !exists {
		return nil, fmt.Errorf("starting point '%s' not found in tickets", start)
	}

	var itinerary []string
	var dfs func(current string)
	dfs = func(current string) {
		for len(graph[current]) > 0 {
			next := graph[current][0]
			graph[current] = graph[current][1:]
			dfs(next)
		}
		itinerary = append([]string{current}, itinerary...)
	}

	dfs(start) // Start DFS from determined or default start

	return itinerary, nil
}
