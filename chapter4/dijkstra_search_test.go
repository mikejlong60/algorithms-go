package chapter4

import (
	"fmt"
	"github.com/go-test/deep"
	"math"
	"testing"
)

func TestDikjstra1(t *testing.T) {
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
	na.Connections = map[string]Node4{"B": {Id: "B", Distance: 1}, "C": {Id: "C", Distance: 4}}
	nb.Connections = map[string]Node4{"C": {Id: "C", Distance: 2}, "D": {Id: "D", Distance: 5}}
	nc.Connections = map[string]Node4{"D": {Id: "D", Distance: 1}}
	nd.Connections = map[string]Node4{}

	graph := map[string]Node4{"A": na, "B": nb, "C": nc, "D": nd}
	actual := DijkstraSearch(graph, "A")
	expected := map[string]*Pq{"A": &Pq{Id: "A", Distance: 0}, "B": &Pq{Id: "B", Distance: 1}, "C": &Pq{Id: "C", Distance: 3}, "D": &Pq{Id: "D", Distance: 4}}
	if diff := deep.Equal(actual, expected); diff != nil {
		t.Error(diff)
	}
}

func TestDikjstra2(t *testing.T) {
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
	na.Connections = map[string]Node4{"B": {Id: "B", Distance: 1}, "C": {Id: "C", Distance: 4}}
	nb.Connections = map[string]Node4{"C": {Id: "C", Distance: 2}, "D": {Id: "D", Distance: 5}}
	nc.Connections = map[string]Node4{"D": {Id: "D", Distance: 1}}
	nd.Connections = map[string]Node4{}

	graph := map[string]Node4{"A": na, "B": nb, "C": nc, "D": nd}
	actual := DijkstraSearch(graph, "B")
	expected := map[string]*Pq{"A": &Pq{Id: "A", Distance: math.MaxInt64}, "B": &Pq{Id: "B", Distance: 0}, "C": &Pq{Id: "C", Distance: 2}, "D": &Pq{Id: "D", Distance: 3}}
	if diff := deep.Equal(actual, expected); diff != nil {
		t.Error(diff)
	}
}

func TestDikjstraFromBook(t *testing.T) {
	ns := Node4{
		Id:          "S",
		Connections: nil,
	}
	nu := Node4{
		Id:          "U",
		Connections: nil,
	}
	nv := Node4{
		Id:          "V",
		Connections: nil,
	}
	ny := Node4{
		Id:          "Y",
		Connections: nil,
	}
	nx := Node4{
		Id:          "X",
		Connections: nil,
	}
	nz := Node4{
		Id:          "Z",
		Connections: nil,
	}
	ns.Connections = map[string]Node4{"U": {Id: "U", Distance: 1}, "S": {Id: "S", Distance: 4}, "V": {Id: "V", Distance: 2}}
	nu.Connections = map[string]Node4{"Y": {Id: "Y", Distance: 3}, "X": {Id: "X", Distance: 1}}
	nv.Connections = map[string]Node4{"X": {Id: "X", Distance: 2}, "Z": {Id: "Z", Distance: 3}}
	nx.Connections = map[string]Node4{"Y": {Id: "Y", Distance: 1}, "Z": {Id: "Z", Distance: 2}}

	graph := map[string]Node4{"S": ns, "U": nu, "V": nv, "Y": ny, "X": nx, "Z": nz}
	actual := DijkstraSearch(graph, "S")
	fmt.Println(actual)
	expected := map[string]*Pq{"S": {Id: "S", Distance: 0}, "U": {Id: "U", Distance: 1}, "V": {Id: "V", Distance: 2}, "Y": {Id: "Y", Distance: 3}, "X": {Id: "X", Distance: 2}, "Z": {Id: "Z", Distance: 4}}
	if diff := deep.Equal(actual, expected); diff != nil {
		t.Error(diff)
	}
}
