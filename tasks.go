package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/eapache/queue"
)

func BFS(root *Vertex) []*Vertex {
	q := queue.New()
	q.Add(root)
	visited := make(map[string]*Vertex)
	all := make([]*Vertex, 0)
	for q.Length() > 0 {
		qSize := q.Length()
		for i := 0; i < qSize; i++ {
			//pop vertex
			currVert := q.Remove().(*Vertex)
			if _, ok := visited[currVert.Key]; ok {
				continue
			}
			visited[currVert.Key] = currVert
			all = append([]*Vertex{currVert}, all...)
			for _, val := range currVert.Children {
				if _, ok := visited[val.Key]; !ok {
					q.Add(val) //add child
				}
			}
		}
	}
	return all
}

func BFSNew(root *Vertex) [][]*Vertex {
	q := queue.New()
	q.Add(root)
	visited := make(map[string]*Vertex)
	all := make([][]*Vertex, 0)
	for q.Length() > 0 {
		qSize := q.Length()
		tmp := make([]*Vertex, 0)
		for i := 0; i < qSize; i++ {
			//pop vertex
			currVert := q.Remove().(*Vertex)
			if _, ok := visited[currVert.Key]; ok {
				continue
			}
			visited[currVert.Key] = currVert
			tmp = append(tmp, currVert)
			for _, val := range currVert.Children {
				if _, ok := visited[val.Key]; !ok {
					q.Add(val) //add child
				}
			}
		}
		all = append([][]*Vertex{tmp}, all...)
	}
	return all
}

// 同层任务 —— 串行执行
func doTasksSerially(vertexs []*Vertex) {
	for _, v := range vertexs {
		time.Sleep(1 * time.Second)
		fmt.Printf("do %v, result is %v \n", v.Key, v.Value)
	}
}

// 同层任务 —— 并发执行
func doTasksParallel(vertexs []*Vertex) {
	var wg sync.WaitGroup
	for _, v := range vertexs {
		wg.Add(1)
		go func(v *Vertex) {
			defer wg.Done()
			time.Sleep(1 * time.Second)
			fmt.Printf("do %v, result is %v \n", v.Key, v.Value)
		}(v) //notice
	}
	wg.Wait()
}
