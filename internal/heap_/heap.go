package heap_

import (
	"container/heap"
	"fmt"
)

type Node struct {
	v  string
	w  int
	id int
}

type PriorityQueue []*Node

func (q PriorityQueue) Len() int {
	return len(q)
}
func (q PriorityQueue) Less(i, j int) bool {
	return q[i].w < q[j].w
}

func (q PriorityQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].id = i
	q[j].id = j
}
func (q *PriorityQueue) Push(n interface{}) {
	N := n.(*Node)
	N.id = q.Len()
	*q = append(*q, N)
}
func (q *PriorityQueue) Pop() interface{} {
	l := q.Len()
	last := (*q)[l-1]
	*q = (*q)[:l-1]
	return last
}
func Heap() {
	nodes := []*Node{
		{v: "apple", w: 10},
		{v: "banana", w: 3},
		{v: "pear", w: 4},
	}
	pq := make(PriorityQueue, len(nodes))
	for i, n := range nodes {
		pq[i] = &Node{
			v:  n.v,
			w:  n.w,
			id: i,
		}
	}
	for pq.Len() > 0 {
		n := heap.Pop(&pq).(*Node)
		fmt.Printf("%.2d:%s", n.w, n.v)
	}
}
