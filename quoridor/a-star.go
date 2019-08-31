package quoridor

import (
	"boardbots-server/util"
	"container/heap"
	"math"
)

/**
 * A struct for the priority queue, holds the Position and priority of the Node in the Board graph
 */
type PQNode struct {
	position util.Position
	prev     *PQNode

	distance, priority int
}

// Calculates the priority of the node as sum of distance to goal + path so far
func (node *PQNode) setPriority(goal util.Position) {
	if goal.Row < 0 {
		node.priority = util.AbsInt(goal.Col-node.position.Col) + node.distance
	} else if goal.Col < 0 {
		node.priority = util.AbsInt(goal.Row-node.position.Row) + node.distance
	}
}

/**
 * Priority Queue methods for the heap
 */
type PriorityQueue []*PQNode

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(left, right int) bool {
	return pq[left].priority < pq[right].priority
}
func (pq PriorityQueue) Swap(left, right int) {
	pq[left], pq[right] = pq[right], pq[left]
}

func (pq *PriorityQueue) Push(item interface{}) {
	*pq = append(*pq, item.(*PQNode))
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

/**
 * Simple A* Algorithm, implementing a best first search by calculating distance to goal plus
 * the length of path from the pawn.
 */
func (game *Game) FindPath(start, goal util.Position) []util.Position {
	pq := &PriorityQueue{}
	heap.Init(pq)
	//breath-first / best-first
	pq.Push(newPQNode(start, goal))
	visited := make(map[util.Position]bool)
	for pq.Len() != 0 {
		node := heap.Pop(pq).(*PQNode)
		if reachedGoal(node.position, goal) {
			return buildPath(node)
		}
		if _, seen := visited[node.position]; seen {
			continue
		}
		visited[node.position] = true
		neighbors := getReachableNeighbors(node, goal, game)
		for _, neighbor := range neighbors {
			heap.Push(pq, neighbor)
		}
	}
	return nil
}
func reachedGoal(current util.Position, goal util.Position) bool {
	return current.Row == goal.Row || current.Col == goal.Col
}

func getReachableNeighbors(node *PQNode, goal util.Position, game *Game) []*PQNode {
	neighbors := make([]*PQNode, 0, 4)
	for _, dir := range directions {
		neighborPositions := game.Board.getValidMoveByDirection(node.position, dir)
		if neighborPositions != nil {
			for _, neighborPos := range neighborPositions {
				node := &PQNode{position: neighborPos, prev: node, priority: math.MaxInt32, distance: node.distance + 1}
				node.setPriority(goal)
				neighbors = append(neighbors, node)
			}
		}
	}
	return neighbors
}

func buildPath(node *PQNode) []util.Position {
	path := make([]util.Position, 0)
	cursor := node
	for cursor.prev != nil {
		path = append(path, cursor.position)
		cursor = cursor.prev
	}
	//reverse
	for idx := len(path)/2 - 1; idx >= 0; idx-- {
		opp := len(path) - 1 - idx
		path[idx], path[opp] = path[opp], path[idx]
	}
	return path
}

func newPQNode(position util.Position, goal util.Position) *PQNode {
	node := &PQNode{
		position: position,
	}
	node.setPriority(goal)
	return node
}
