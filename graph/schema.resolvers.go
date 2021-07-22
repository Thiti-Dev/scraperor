package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/Thiti-Dev/scraperor/graph/generated"
	"github.com/Thiti-Dev/scraperor/graph/model"
)

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	user := new(model.User)
	user.Email = "Thiti Mahawannakit"
	user.ID = "xxx64"

	var users []*model.User
	users = append(users, user)
	return users,nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

type mutationResolver struct{ *Resolver }
