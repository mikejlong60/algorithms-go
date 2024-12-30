package chapter5

import (
	"fmt"
	"github.com/greymatter-io/golangz/sets"
	log "github.com/sirupsen/logrus"
	"testing"
)

var nodeEq = func(x, y int) bool {
	if x == y {
		return true
	} else {
		return false
	}
}

func printLocalMinimums(r []*Node) []int {
	var mins []int
	fmt.Print(len(r))
	for _, j := range r {
		log.Infof("Value:%v", j.Value)
		mins = append(mins, j.Value)
	}
	return mins
}

var a, b, c, d, e, f, g, h, i, j, k, l, m, n, o Node

func buildNodes() {
	totalSteps = 0
	a = Node{
		Value: 140,
	}

	b = Node{
		Value: 130,
	}
	c = Node{
		Value: 120,
	}
	d = Node{
		Value: 1,
	}
	e = Node{
		Value: 2,
	}
	f = Node{
		Value: 90,
	}
	g = Node{
		Value: 80,
	}
	h = Node{
		Value: 70,
	}
	i = Node{
		Value: 60,
	}
	j = Node{
		Value: 50,
	}
	k = Node{
		Value: 40,
	}
	l = Node{
		Value: 3,
	}
	m = Node{
		Value: 20,
	}
	n = Node{
		Value: 10,
	}
	o = Node{
		Value: 1800,
	}
	a.Left = &b
	a.Right = &c
	b.Parent = &a
	c.Parent = &a
	d.Parent = &b
	e.Parent = &b
	f.Parent = &c
	g.Parent = &c

	h.Parent = &d
	i.Parent = &d
	j.Parent = &e
	k.Parent = &e
	l.Parent = &f
	m.Parent = &f
	n.Parent = &g
	o.Parent = &g

	b.Left = &d
	b.Right = &e
	c.Left = &f
	c.Right = &g
	d.Left = &h
	d.Right = &i
	e.Left = &j
	e.Right = &k
	f.Left = &l
	f.Right = &m
	g.Left = &n
	g.Right = &o
}

func TestLocalMinimumsFromB(t *testing.T) {
	buildNodes()
	r := LocalMinimums(&b)
	mins := printLocalMinimums(r)
	var expected = []int{1, 2}
	if !sets.SetEquality(mins, expected, nodeEq) {
		t.Errorf("Actual:%v Expected:%v", mins, expected)
	}
	log.Infof("Total Steps:%v", totalSteps)
}

func TestLocalMinimumsFromC(t *testing.T) {
	buildNodes()
	r := LocalMinimums(&c)
	mins := printLocalMinimums(r)
	var expected = []int{3, 20, 10}
	if !sets.SetEquality(mins, expected, nodeEq) {
		t.Errorf("Actual:%v Expected:%v", mins, expected)
	}
	log.Infof("Total Steps:%v", totalSteps)
}

func TestLocalMinimumsFromD(t *testing.T) {
	buildNodes()
	r := LocalMinimums(&d)
	mins := printLocalMinimums(r)
	var expected = []int{1}
	if !sets.SetEquality(mins, expected, nodeEq) {
		t.Errorf("Actual:%v Expected:%v", mins, expected)
	}
	log.Infof("Total Steps:%v", totalSteps)
}

func TestLocalMinimumsFromO(t *testing.T) {
	buildNodes()
	r := LocalMinimums(&o)
	mins := printLocalMinimums(r)
	var expected = []int{}
	if !sets.SetEquality(mins, expected, nodeEq) {
		t.Errorf("Actual:%v Expected:%v", mins, expected)
	}
	log.Infof("Total Steps:%v", totalSteps)
}

func TestLocalMinimumsFromG(t *testing.T) {
	buildNodes()
	r := LocalMinimums(&g)
	mins := printLocalMinimums(r)
	var expected = []int{10}
	if !sets.SetEquality(mins, expected, nodeEq) {
		t.Errorf("Actual:%v Expected:%v", mins, expected)
	}
	log.Infof("Total Steps:%v", totalSteps)
}

func TestLocalMinimumsFromA(t *testing.T) {
	buildNodes()
	r := LocalMinimums(&a)
	mins := printLocalMinimums(r)
	var expected = []int{1, 2, 3, 20, 10}
	if !sets.SetEquality(mins, expected, nodeEq) {
		t.Errorf("Actual:%v Expected:%v", mins, expected)
	}
	log.Infof("Total Steps:%v", totalSteps)
}
