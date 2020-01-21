package graphql

import graphql "github.com/graph-gophers/graphql-go"

type KeyboardResolver struct {
	k *keyboard
}

type keyboard struct {
	ID   graphql.ID
	Name string
}

func (r *KeyboardResolver) Name() string {
	return r.k.Name
}

func (r *KeyboardResolver) ID() graphql.ID {
	return r.k.ID
}

func (_ *Resolver) Keyboard(args struct{ ID graphql.ID }) *KeyboardResolver {
	if s := keyboardData[args.ID]; s != nil {
		return &KeyboardResolver{s}
	}
	return nil
}

func (_ *Resolver) Keyboards() *[]*KeyboardResolver {
	var keebs []*KeyboardResolver
	keebs = append(keebs, &KeyboardResolver{&keyboard{ID: "3", Name: "unikorn"}})
	return &keebs
}
