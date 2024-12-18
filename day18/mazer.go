package main

import (
	"container/heap"

	"github.com/AdelSehic/advent-of-code-2024/helpers"
)

func PathTraceDijkstra(start, end *helpers.Coord, maze *helpers.Field) *PathTracer {
	solutions := make([]*PathTracer, 0)
	visited := make(map[helpers.Coord]bool)
	cost := make(map[helpers.Coord]int)
	pq := &PriorityQueue{}
	heap.Init(pq)

	heap.Push(pq, &Item{
		tracer:   NewTracer(start, start),
		priority: 0,
	})
	cost[*start] = 0

	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)
		current := item.tracer
		currentCost := item.priority

		if visited[*current.Pos()] || maze.GetLetter(current.Pos()) == '#' {
			continue
		}
		visited[*current.Pos()] = true

		if maze.GetLetter(current.Iter.Position) == 'E' {
			solutions = append(solutions, current)
			return current
		}

		neigbors := []*PathTracer{}
		neigbors = append(neigbors, current.Copy())
		neigbors = append(neigbors, current.Rotate())
		neigbors = append(neigbors, current.Rotate().Rotate())
		neigbors = append(neigbors, current.RotateOther())

		for _, neigh := range neigbors {
			neigh.Move()
			if visited[*neigh.Pos()] || maze.GetLetter(neigh.Pos()) == '#' {
				continue
			}

			newCost := currentCost + 1
			if existingCost, exists := cost[*neigh.Pos()]; !exists || newCost < existingCost {
				cost[*neigh.Pos()] = newCost
				heap.Push(pq, &Item{
					tracer:   neigh,
					priority: newCost,
				})
			}
		}
	}
	return nil
}
