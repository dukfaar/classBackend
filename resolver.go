package main

import (
	"context"

	"github.com/dukfaar/classBackend/class"
	"github.com/dukfaar/goUtils/permission"
	"github.com/dukfaar/goUtils/relay"
	"github.com/globalsign/mgo/bson"
	graphql "github.com/graph-gophers/graphql-go"
)

type Resolver struct {
}

func (r *Resolver) Classes(ctx context.Context, args struct {
	First  *int32
	Last   *int32
	Before *string
	After  *string
}) (*class.ConnectionResolver, error) {
	classService := ctx.Value("classService").(class.Service)

	var totalChannel = make(chan int)
	go func() {
		var total, _ = classService.Count()
		totalChannel <- total
	}()

	var classesChannel = make(chan []class.Model)
	go func() {
		result, _ := classService.List(args.First, args.Last, args.Before, args.After)
		classesChannel <- result
	}()

	var (
		start string
		end   string
	)

	var classes = <-classesChannel

	if len(classes) == 0 {
		start, end = "", ""
	} else {
		start, end = classes[0].ID.Hex(), classes[len(classes)-1].ID.Hex()
	}

	hasPreviousPageChannel, hasNextPageChannel := relay.GetHasPreviousAndNextPage(len(classes), start, end, classService)

	return &class.ConnectionResolver{
		Models: classes,
		ConnectionResolver: relay.ConnectionResolver{
			relay.Connection{
				Total:           int32(<-totalChannel),
				From:            start,
				To:              end,
				HasNextPage:     <-hasNextPageChannel,
				HasPreviousPage: <-hasPreviousPageChannel,
			},
		},
	}, nil
}

func setDataOnModel(model *class.Model, input *class.MutationInput) {
	if input.Name != nil {
		model.Name = *input.Name
	}

	if input.NamespaceID != nil {
		model.NamespaceID = bson.ObjectIdHex(*input.NamespaceID)
	}

	if input.Synonyms != nil {
		model.Synonyms = make([]string, len(*input.Synonyms))
		for i := range *input.Synonyms {
			model.Synonyms[i] = *(*input.Synonyms)[i]
		}
	}
}

func (r *Resolver) CreateClass(ctx context.Context, args struct {
	Input *class.MutationInput
}) (*class.Resolver, error) {
	err := permission.Check(ctx, "mutation.createClass")
	if err != nil {
		return nil, err
	}

	classService := ctx.Value("classService").(class.Service)

	inputModel := class.Model{}
	setDataOnModel(&inputModel, args.Input)

	newModel, err := classService.Create(&inputModel)

	if err == nil {
		return &class.Resolver{
			Model: newModel,
		}, nil
	}

	return nil, err
}

func (r *Resolver) UpdateClass(ctx context.Context, args struct {
	Id    string
	Input *class.MutationInput
}) (*class.Resolver, error) {
	err := permission.Check(ctx, "mutation.updateClass")
	if err != nil {
		return nil, err
	}

	classService := ctx.Value("classService").(class.Service)

	model, err := classService.FindByID(args.Id)
	setDataOnModel(model, args.Input)

	newModel, err := classService.Update(args.Id, model)

	if err == nil {
		return &class.Resolver{
			Model: newModel,
		}, nil
	}

	return nil, err
}

func (r *Resolver) DeleteClass(ctx context.Context, args struct {
	Id string
}) (*graphql.ID, error) {
	err := permission.Check(ctx, "mutation.deleteClass")
	if err != nil {
		return nil, err
	}

	classService := ctx.Value("classService").(class.Service)

	deletedID, err := classService.DeleteByID(args.Id)
	result := graphql.ID(deletedID)

	if err == nil {
		return &result, nil
	}

	return nil, err
}

func (r *Resolver) Class(ctx context.Context, args struct {
	Id string
}) (*class.Resolver, error) {
	classService := ctx.Value("classService").(class.Service)

	queryNamespace, err := classService.FindByID(args.Id)

	if err == nil {
		return &class.Resolver{
			Model: queryNamespace,
		}, nil
	}

	return nil, err
}

func (r *Resolver) ClassByName(ctx context.Context, args struct {
	Name        string
	NamespaceId *string
}) (*class.Resolver, error) {
	classService := ctx.Value("classService").(class.Service)

	queryNamespace, err := classService.FindByName(args.Name, args.NamespaceId)

	if err == nil {
		return &class.Resolver{
			Model: queryNamespace,
		}, nil
	}

	return nil, err
}

func (r *Resolver) ClassByNameOrSynonym(ctx context.Context, args struct {
	Name        string
	NamespaceId *string
}) (*class.Resolver, error) {
	classService := ctx.Value("classService").(class.Service)

	queryNamespace, err := classService.FindByNameOrSynonym(args.Name, args.NamespaceId)

	if err == nil {
		return &class.Resolver{
			Model: queryNamespace,
		}, nil
	}

	return nil, err
}
