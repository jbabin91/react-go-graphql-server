package queries

import (
	"context"
	"github.com/graphql-go/graphql"
	"github.com/mongodb/mongo-go-driver/bson"

	"app/data"
	types "app/types"
)

type todoStruct struct {
	NAME string `json:"name"`
	DESCRIPTION string `json:"description"`
}

var GetNotTodos = &graphql.Field {
	Type: graphql.NewList(types.NotTodo),
	Description: "Get all not todos",
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		notTodoCollection := mongo.Client.Database("medium-app").Collection("Not_Todos")

		todos, err := notTodoCollection.Find(context.Background(), nil)

		if err != nil []todoStruct
	}
}