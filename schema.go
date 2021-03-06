package main

import (
	"github.com/dukfaar/classBackend/class"
	"github.com/dukfaar/goUtils/relay"
)

var Schema string = `
		schema {
			query: Query
			mutation: Mutation
		}

		type Query {
			classes(first: Int, last: Int, before: String, after: String): ClassConnection!
			class(id: ID!): Class!
			classByName(name: String!, namespaceId: ID): Class!
			classByNameOrSynonym(name: String!, namespaceId: ID): Class!
		}

		input ClassMutationInput {
			name: String
			namespaceId: ID
			synonyms: [String]
			maxLevel: Int
		}

		type Mutation {
			createClass(input: ClassMutationInput): Class!
			updateClass(id: ID!, input: ClassMutationInput): Class!
			deleteClass(id: ID!): ID
		}` +
	relay.PageInfoGraphQLString +
	class.GraphQLType
