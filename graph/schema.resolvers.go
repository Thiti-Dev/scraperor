package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/Thiti-Dev/scraperor/core"
	"github.com/Thiti-Dev/scraperor/graph/generated"
	"github.com/Thiti-Dev/scraperor/graph/model"
	"github.com/Thiti-Dev/scraperor/models"
)

func (r *mutationResolver) ScrapeNow(ctx context.Context, input *model.ScrapeInput) (*model.ScrapeOutput, error) {
	
	engineEntity := core.CreateEngineEntity()
	engineEntity.AddTask(input.TargetURL,input.Pointer)
	
	elements := engineEntity.Scrape()
	return &model.ScrapeOutput{
		Elements: elements,
	},nil
}

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

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
