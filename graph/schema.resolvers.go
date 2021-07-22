package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/Thiti-Dev/scraperor/graph/generated"
	"github.com/Thiti-Dev/scraperor/models"
)

func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	/*user := new(model.User)
	user.Email = "Thiti Mahawannakit"
	user.ID = "xxx64"

	var users []*model.User
	users = append(users, user)
	return users,nil*/
	users, err := r.UsersRepository.GetUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
type mutationResolver struct{ *Resolver }
