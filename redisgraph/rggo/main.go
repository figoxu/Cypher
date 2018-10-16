package main

import (
	"github.com/gomodule/redigo/redis"
	rg "github.com/redislabs/redisgraph-go"
)

func main() {
	conn, _ := redis.Dial("tcp", "0.0.0.0:6379")
	defer conn.Close()

	graph := rg.Graph{}.New("social", conn)

	john := rg.Node{
		Label: "person",
		Properties: map[string]interface{}{
			"name":   "John Doe",
			"age":    33,
			"gender": "male",
			"status": "single",
		},
	}
	graph.AddNode(&john)

	japan := rg.Node{
		Label: "country",
		Properties: map[string]interface{}{
			"name": "Japan",
		},
	}
	graph.AddNode(&japan)

	edge := rg.Edge{
		Source:      &john,
		Relation:    "visited",
		Destination: &japan,
	}
	graph.AddEdge(&edge)

	graph.Commit()

	query := `MATCH (p:person)-[v:visited]->(c:country)
		   RETURN p.name, p.age, v.purpose, c.name`
	rs, _ := graph.Query(query)

	rs.PrettyPrint()
}