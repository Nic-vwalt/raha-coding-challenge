# Raha Coding Challenge

This repository contains the solution to a coding challenge - Reconstruct Itinerary.

## Run Command

To run the application, execute the following command in the terminal:

```bash
go run cmd/debug/main.go
```

## POST Endpoint

The POST endpoint setup to handle your JSON input.

```bash
localhost:4000/itinerary/transform
```

## Solution

The solution to this itinerary problem is:

- A graph is created as a map where each key is a source and the value is a slice of destinations. The original leet code question always starts at JFK, but I've altered my solution to be dynamic as a starting airport was not specified.
- For each ticket, the source and destination are added to the graph.
- The destinations for each source in the graph is sorted in lexicographical order.
- A Depth-First Search (DFS) is performed starting from the determined or default start point.
- Once we reach an airport with no outgoing edges/destinations we start backtracking. Starting with the current airport, building the rest of the itinerary in reverse.
- Finally reverse the current list because we have been building from end to start.
