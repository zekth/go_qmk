package graphql

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/zekth/go_qmk/api/internal/dependencies"
	"github.com/zekth/go_qmk/api/internal/utils"
)

type Resolver struct{}

var keyboards = []*keyboard{
	{ID: "1", Name: "unikorn"},
	{ID: "2", Name: "Jane"},
}
var keyboardData = make(map[graphql.ID]*keyboard)

type Handler struct {
	Schema       *graphql.Schema
	Handler      *relay.Handler
	Dependencies dependencies.Dependencies
}

func NewHandler(filename string, dependencies dependencies.Dependencies) (*Handler, error) {
	// mocks
	for _, s := range keyboards {
		keyboardData[s.ID] = s
	}

	h := Handler{}
	h.Dependencies = dependencies
	buff, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	s := string(buff)
	schema, err := graphql.ParseSchema(s, &Resolver{})
	if err != nil {
		return nil, err
	}
	h.Schema = schema
	h.Handler = &relay.Handler{Schema: schema}
	return &h, nil
}

// Handler is the HTTP handler for graphql
func (h *Handler) GetHTTPHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := utils.InjectEnvInContext(c.Request.Context(), h.Dependencies)
		h.Handler.ServeHTTP(c.Writer, c.Request.WithContext(ctx))
	}
}
