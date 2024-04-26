# Raha Coding Challenge

This repository contains the solution to a coding challenge involving the construction of an itinerary from a list of flights. The solution utilizes a graph representation of the flights and performs a Depth-First Search (DFS) to establish the itinerary sequence. Unlike the original problem which always starts from JFK, this implementation allows for a dynamic starting point.

## Run Command

To run the application, execute the following command in the terminal:

```bash
go run cmd/debug/main.go
```

## Solution

The solution to this itinerary problem is to:

- initialize a graph that represents the flights
- DFS (Depth-First Search) function which in the original leet code question always starts at JFK, but I've altered my solution to be dynamic as a starting airport was not specified. Inside this function we loop until we find an airport that has no more destinations.
- Once we reach an airport with no outgoing edges/destinations we start backtracking. Starting with the current airport, building the rest of the itinerary in reverse.
- Finally reverse the current list because we have been building from end to start.
