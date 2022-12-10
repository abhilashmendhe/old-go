package main

import (
	"math/rand"
	"time"
)

const (
	FOLLOWER  = "follower"
	CANDIDATE = "candidate"
	LEADER    = "leader"
)

// Node represents a single node in the Raft cluster.
type Node struct {
	ID        int
	State     string
	Leader    int
	Votes     int
	NextIndex int
	Log       []string
}

// NewNode returns a new Node with the given ID.
func NewNode(id int) *Node {
	return &Node{
		ID:    id,
		State: FOLLOWER,
	}
}

// RequestVote asks other nodes for their vote in the election.
func (n *Node) RequestVote() {
	n.State = CANDIDATE
	n.Votes = 1
	n.Leader = n.ID
	// Send RequestVote RPCs to other nodes.
}

// AppendEntries appends entries to the node's log.
func (n *Node) AppendEntries(entries []string) {
	n.State = LEADER
	n.NextIndex = len(n.Log)
	n.Log = append(n.Log, entries...)
	// Send AppendEntries RPCs to other nodes.
}

func main() {
	// Initialize the cluster with three nodes.
	nodes := []*Node{
		NewNode(1),
		NewNode(2),
		NewNode(3),
	}

	// Set a random election timeout for each node.
	for _, node := range nodes {
		node.ElectionTimeout = time.Duration(rand.Intn(150) + 150)
	}

	// Start the simulation.
	for {
		// Loop through each node and update its state.
		for _, node := range nodes {
			switch node.State {
			case FOLLOWER:
				node.ElectionTimeout--
				if node.ElectionTimeout <= 0 {
					node.RequestVote()
				}

			}
		}
	}
}
