package graph

import (
	"github.com/feedloop/syde/internal/storage"
)

// Node represents an entity in the graph.
type Node struct {
	ID   string
	Kind string
	Name string
	File string
}

// Edge represents a relationship in the graph.
type Edge struct {
	Source string
	Target string
	Type   string
	Label  string
}

// SubgraphResult is the result of a graph query.
type SubgraphResult struct {
	Nodes []Node
	Edges []Edge
}

// Neighbors returns all directly connected entities.
func Neighbors(idx *storage.Index, entityID string) (*SubgraphResult, error) {
	result := &SubgraphResult{}
	seen := make(map[string]bool)

	outbound, err := idx.GetOutbound(entityID)
	if err != nil {
		return nil, err
	}
	for _, rel := range outbound {
		if !seen[rel.Target] {
			result.Nodes = append(result.Nodes, Node{ID: rel.Target})
			seen[rel.Target] = true
		}
		result.Edges = append(result.Edges, Edge{
			Source: entityID,
			Target: rel.Target,
			Type:   rel.Type,
			Label:  rel.Rel.Label,
		})
	}

	inbound, err := idx.GetInbound(entityID)
	if err != nil {
		return nil, err
	}
	for _, rel := range inbound {
		if !seen[rel.Source] {
			result.Nodes = append(result.Nodes, Node{ID: rel.Source})
			seen[rel.Source] = true
		}
		result.Edges = append(result.Edges, Edge{
			Source: rel.Source,
			Target: entityID,
			Type:   rel.Type,
			Label:  rel.Rel.Label,
		})
	}

	return result, nil
}

// ImpactAnalysis does BFS to find all transitively connected entities.
func ImpactAnalysis(idx *storage.Index, entityID string, maxDepth int) (map[int][]Node, error) {
	if maxDepth <= 0 {
		maxDepth = 3
	}

	result := make(map[int][]Node)
	visited := map[string]bool{entityID: true}
	queue := []struct {
		id    string
		depth int
	}{{entityID, 0}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.depth >= maxDepth {
			continue
		}

		// Get outbound relationships
		outbound, _ := idx.GetOutbound(current.id)
		for _, rel := range outbound {
			if !visited[rel.Target] {
				visited[rel.Target] = true
				node := Node{ID: rel.Target}
				result[current.depth+1] = append(result[current.depth+1], node)
				queue = append(queue, struct {
					id    string
					depth int
				}{rel.Target, current.depth + 1})
			}
		}

		// Get inbound relationships
		inbound, _ := idx.GetInbound(current.id)
		for _, rel := range inbound {
			if !visited[rel.Source] {
				visited[rel.Source] = true
				node := Node{ID: rel.Source}
				result[current.depth+1] = append(result[current.depth+1], node)
				queue = append(queue, struct {
					id    string
					depth int
				}{rel.Source, current.depth + 1})
			}
		}
	}

	return result, nil
}
