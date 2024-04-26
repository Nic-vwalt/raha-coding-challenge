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

// Graph based approach to reconstructing an itinerary from a list of tickets.
func (is *ItineraryService) ReconstructItinerary(ctx context.Context, tickets [][]string) ([]string, error) {
	if len(tickets) == 0 {
		return nil, fmt.Errorf("no itineraries provided")
	}

	graph := make(map[string][]string)
	departureCount := make(map[string]int)
	arrivalCount := make(map[string]int)

	// Build the graph and count departures and arrivals
	for _, ticket := range tickets {
		source, dest := ticket[0], ticket[1]
		graph[source] = append(graph[source], dest)
		departureCount[source]++
		arrivalCount[dest]++
	}

	// Sort destinations to ensure lexical order
	for key := range graph {
		sort.Strings(graph[key])
	}

	// Determine the starting point
	start := "JFK" // Default to JFK if no unique starting point is found
	for airport := range departureCount {
		if departureCount[airport] > arrivalCount[airport] {
			start = airport
			break
		}
	}

	// Check if the starting point is valid
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

	// Start DFS from the determined start point
	dfs(start)

	return itinerary, nil
}

// Improved version of ReconstructItinerary that uses a stack and slice approach for iterative DFS. But might not scale as well as the graph based approach.
func (is *ItineraryService) StackSliceReconstructItinerary(ctx context.Context, tickets [][]string) ([]string, error) {
    if len(tickets) == 0 {
        return nil, fmt.Errorf("no tickets provided")
    }

    adjList := make(map[string][]string)
    departureCount := make(map[string]int)
    arrivalCount := make(map[string]int)

    // Build the adjacency list and count departures and arrivals
    for _, ticket := range tickets {
        source, dest := ticket[0], ticket[1]
        adjList[source] = append(adjList[source], dest)
        departureCount[source]++
        arrivalCount[dest]++
    }

    // Sort each list in reverse order to use the slice as a stack
    for key := range adjList {
        sort.Sort(sort.Reverse(sort.StringSlice(adjList[key])))
    }

    // Determine the starting point dynamically
    start := "JFK" // Default to JFK if no unique starting point is found
    for airport, dep := range departureCount {
        if dep > arrivalCount[airport] {
            start = airport
            break
        }
    }

    // Edge case: Ensure starting point has outgoing tickets
    if _, exists := adjList[start]; !exists {
        return nil, fmt.Errorf("starting point '%s' not found in tickets", start)
    }

    // Use iterative DFS with a stack
    stack := []string{start}
    route := make([]string, 0)

    for len(stack) > 0 {
        top := stack[len(stack)-1]
        if destinations, ok := adjList[top]; ok && len(destinations) > 0 {
            nextDest := destinations[len(destinations)-1]
            adjList[top] = destinations[:len(destinations)-1] // Pop destination
            stack = append(stack, nextDest)
        } else {
            stack = stack[:len(stack)-1]
            route = append([]string{top}, route...) // Prepend to build the itinerary in correct order
        }
    }

    return route, nil
}
