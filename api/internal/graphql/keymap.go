package graphql

import (
	"context"

	graphql "github.com/graph-gophers/graphql-go"
)

type keymap struct {
	ID graphql.ID
}

type KeymapResolver struct {
	k *keymap
}

func (r *KeymapResolver) ID() graphql.ID {
	return r.k.ID
}

func (_ *Resolver) Keymap(ctx context.Context, args struct{ ID graphql.ID }) *KeymapResolver {
	return &KeymapResolver{&keymap{ID: "666"}}
}
