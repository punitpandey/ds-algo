package graph

type graph struct {
	Length       int
	AdjacentList map[interface{}][]interface{}
}

func New() *graph {
	return &graph{
		Length:       0,
		AdjacentList: make(map[interface{}][]interface{}),
	}
}

func (graph *graph) AddNode(name interface{}) {
	graph.AdjacentList[name] = nil
	graph.Length++
}

func (graph *graph) RemoveNode(name interface{}) {
	if graph.Length == 0 {
		return
	}
	//clear all connections with the node to be deleted
	for _, node := range graph.AdjacentList[name] {
		for i, adjacentNode := range graph.AdjacentList[node] {
			if adjacentNode == name {
				graph.AdjacentList[node][i] = graph.AdjacentList[node][len(graph.AdjacentList[node])-1]
				graph.AdjacentList[node] = graph.AdjacentList[node][:len(graph.AdjacentList[node])-1]
			}
		}
	}
	delete(graph.AdjacentList, name)
	graph.Length--
}

func (graph *graph) AddEdge(nodeOne, nodeTwo interface{}) {
	graph.AdjacentList[nodeOne] = append(graph.AdjacentList[nodeOne], nodeTwo)
	graph.AdjacentList[nodeTwo] = append(graph.AdjacentList[nodeTwo], nodeOne)
}

func (graph *graph) RemoveEdge(nodeOne, nodeTwo interface{}) {
	//removes connection from nodeOne -> nodeTwo
	for i, node := range graph.AdjacentList[nodeOne] {
		if node == nodeTwo {
			graph.AdjacentList[nodeOne][i] = graph.AdjacentList[nodeOne][len(graph.AdjacentList[nodeOne])-1]
			graph.AdjacentList[nodeOne] = graph.AdjacentList[nodeOne][:len(graph.AdjacentList[nodeOne])-1]
		}
	}
	//removes connection from nodeTwo -> nodeOne
	for i, node := range graph.AdjacentList[nodeTwo] {
		if node == nodeOne {
			graph.AdjacentList[nodeTwo][i] = graph.AdjacentList[nodeTwo][len(graph.AdjacentList[nodeTwo])-1]
			graph.AdjacentList[nodeTwo] = graph.AdjacentList[nodeTwo][:len(graph.AdjacentList[nodeTwo])-1]
		}
	}
}

func (graph *graph) Show() map[interface{}][]interface{} {
	return graph.AdjacentList
}
