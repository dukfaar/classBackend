package class

import (
	"github.com/dukfaar/goUtils/relay"
	"github.com/globalsign/mgo/bson"
)

type Model struct {
	ID          bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string        `json:"name,omitempty" bson:"name,omitempty"`
	NamespaceID bson.ObjectId `json:"namespaceId,omitempty" bson:"namespaceId,omitempty"`
	Synonyms    []string      `json:"synonyms,omitempty" bson:"synonyms,omitempty"`
	MaxLevel    int32         `json:"maxLevel,omitempty" bson:"maxLevel,omitempty"`
}

type MutationInput struct {
	Name        *string
	NamespaceID *string
	Synonyms    *[]*string
	MaxLevel    *int32
}

var GraphQLType = `
type Class {
	_id: ID
	name: String
	namespaceId: ID
	synonyms: [String]
	maxLevel: Int
}
` +
	relay.GenerateConnectionTypes("Class")
