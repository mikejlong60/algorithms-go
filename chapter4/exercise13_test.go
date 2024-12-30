package chapter4

import (
	"github.com/go-test/deep"
	"testing"
)

func TestUsingDijkstraStartingAtA(t *testing.T) {
	na := Node4{
		Id:          "A",
		Connections: nil,
	}
	nb := Node4{
		Id:          "B",
		Connections: nil,
	}
	nc := Node4{
		Id:          "C",
		Connections: nil,
	}
	nd := Node4{
		Id:          "D",
		Connections: nil,
	}
	na.Connections = map[string]Node4{"B": {Id: "B", Distance: 10}, "C": {Id: "C", Distance: 30}}
	nb.Connections = map[string]Node4{"C": {Id: "C", Distance: 3}, "D": {Id: "D", Distance: 20}, "A": {Id: "A", Distance: 40}}
	nc.Connections = map[string]Node4{"D": {Id: "D", Distance: 1}}
	nd.Connections = map[string]Node4{}

	graph := map[string]Node4{"A": na, "B": nb, "C": nc, "D": nd}
	actual := DijkstraSearch(graph, "A")
	expected := map[string]*Pq{"A": {Id: "A", Distance: 0}, "B": {Id: "B", Distance: 10}, "C": {Id: "C", Distance: 13}, "D": {Id: "D", Distance: 14}}
	if diff := deep.Equal(actual, expected); diff != nil {
		t.Error(diff)
	}
}

func TestUsingDijkstraStartingAtB(t *testing.T) {
	na := Node4{
		Id:          "A",
		Connections: nil,
	}
	nb := Node4{
		Id:          "B",
		Connections: nil,
	}
	nc := Node4{
		Id:          "C",
		Connections: nil,
	}
	nd := Node4{
		Id:          "D",
		Connections: nil,
	}
	na.Connections = map[string]Node4{"B": {Id: "B", Distance: 10}, "C": {Id: "C", Distance: 30}}
	nb.Connections = map[string]Node4{"C": {Id: "C", Distance: 3}, "D": {Id: "D", Distance: 20}, "A": {Id: "A", Distance: 40}}
	nc.Connections = map[string]Node4{"D": {Id: "D", Distance: 1}}
	nd.Connections = map[string]Node4{}

	graph := map[string]Node4{"A": na, "B": nb, "C": nc, "D": nd}
	actual := DijkstraSearch(graph, "B")
	expected := map[string]*Pq{"A": {Id: "A", Distance: 40}, "B": {Id: "B", Distance: 0}, "C": {Id: "C", Distance: 3}, "D": {Id: "D", Distance: 4}}
	if diff := deep.Equal(actual, expected); diff != nil {
		t.Error(diff)
	}
}
