package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.47

import (
	"context"
	"pet-pamper/api"
	"pet-pamper/model"
)

// Auth is the resolver for the auth field.
func (r *mutationResolver) Auth(ctx context.Context) (*model.AuthMutation, error) {
	return &model.AuthMutation{}, nil
}

// Product is the resolver for the product field.
func (r *mutationResolver) Product(ctx context.Context) (*model.ProductMutation, error) {
	return &model.ProductMutation{}, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context) (*model.UserQuery, error) {
	return &model.UserQuery{}, nil
}

// Product is the resolver for the product field.
func (r *queryResolver) Product(ctx context.Context) (*model.ProductQuery, error) {
	return &model.ProductQuery{}, nil
}

// Mutation returns api.MutationResolver implementation.
func (r *Resolver) Mutation() api.MutationResolver { return &mutationResolver{r} }

// Query returns api.QueryResolver implementation.
func (r *Resolver) Query() api.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }