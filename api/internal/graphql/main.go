package graphql

import (
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

type Resolver struct{}

var keyboards = []*keyboard{
	{ID: "1", Name: "unikorn"},
	{ID: "2", Name: "Jane"},
}
var keyboardData = make(map[graphql.ID]*keyboard)

// Handler is the HTTP handler for graphql
func Handler() gin.HandlerFunc {

	// mocks
	for _, s := range keyboards {
		keyboardData[s.ID] = s
	}

	buff, err := ioutil.ReadFile("schema.graphql")
	if err != nil {
		log.Fatal(err)
	}
	s := string(buff)
	schema := graphql.MustParseSchema(s, &Resolver{})
	r := &relay.Handler{Schema: schema}
	return func(c *gin.Context) {
		r.ServeHTTP(c.Writer, c.Request)
	}
}
