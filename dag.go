package main

//图结构
type DAG struct {
	Vertexs []*Vertex
}

//顶点
type Vertex struct {
	Key      string
	Value    interface{}
	Parents  []*Vertex
	Children []*Vertex
}

//添加顶点
func (dag *DAG) AddVertex(v *Vertex) {
	dag.Vertexs = append(dag.Vertexs, v)
}

//添加边
func (dag *DAG) AddEdge(from, to *Vertex) {
	from.Children = append(from.Children, to)
	to.Parents = append(to.Parents, from)
}


