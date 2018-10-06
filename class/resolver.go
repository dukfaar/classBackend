package class

import graphql "github.com/graph-gophers/graphql-go"

type Resolver struct {
	Model *Model
}

func (r *Resolver) ID() *graphql.ID {
	id := graphql.ID(r.Model.ID.Hex())
	return &id
}

func (r *Resolver) Name() *string {
	return &r.Model.Name
}

func (r *Resolver) NamespaceID() *graphql.ID {
	id := graphql.ID(r.Model.NamespaceID.Hex())
	return &id
}

func (r *Resolver) MaxLevel() *int32 {
	return &r.Model.MaxLevel
}

func (r *Resolver) Synonyms() *[]*string {
	l := make([]*string, len(r.Model.Synonyms))
	for i, input := range r.Model.Synonyms {
		l[i] = &input
	}

	return &l
}
