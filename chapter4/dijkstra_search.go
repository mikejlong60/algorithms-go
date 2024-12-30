package chapter4

import (
	"math"
)

type Node4 struct {
	Id          string
	Distance    int
	Connections map[string]Node4
}

type Pq struct {
	Distance int
	Id       string
}

//O( )

func DijkstraSearch(graph map[string]Node4, start string) map[string]*Pq {
	var distances = make(map[string]*Pq)
	for i, _ := range graph {
		distances[i] = &Pq{Distance: math.MaxInt64, Id: i}
	}
	distances[start].Distance = 0
	var pq = make([]Pq, 1)
	pq[0] = Pq{0, distances[start].Id}
	for len(pq) > 0 {
		current := pq[0]
		pq = pq[1:]                                               //slice array at second element
		if !(current.Distance > distances[current.Id].Distance) { //If current distance is LE previous distance to this node then replace node path and distance sum with the shorter path distance
			for _, neighbor := range graph[current.Id].Connections {
				distance := current.Distance + neighbor.Distance
				if distance < distances[neighbor.Id].Distance {
					distances[neighbor.Id].Distance = distance
					head := Pq{distance, neighbor.Id}
					headArray := []Pq{head}
					pq = append(headArray, pq...)
				}
			}
		}
	}
	return distances
}
