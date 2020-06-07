package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/MrFuku/kanban-go-nuxt-graphql/server/graph/generated"
	"github.com/MrFuku/kanban-go-nuxt-graphql/server/graph/model"
	"github.com/MrFuku/kanban-go-nuxt-graphql/server/models"
	"github.com/MrFuku/kanban-go-nuxt-graphql/server/util"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*models.Todo, error) {
	todo := &models.Todo{
		ID:     util.CreateUniqueID(),
		Text:   input.Text,
		Done:   false,
		UserID: input.UserID,
	}
	res := r.DB.Create(todo)
	if err := res.Error; err != nil {
		return nil, err
	}
	return todo, nil
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, input model.EditTodo) (*models.Todo, error) {
	todo := &models.Todo{ID: input.ID}
	res := r.DB.First(todo)
	if err := res.Error; err != nil {
		return nil, err
	}
	params := map[string]interface{}{}
	params["text"] = input.Text
	params["done"] = input.Done
	params["user_id"] = input.UserID
	todo.Text = input.Text
	todo.Done = input.Done
	todo.UserID = input.UserID
	res = r.DB.Model(&todo).Update(params)
	if err := res.Error; err != nil {
		return nil, err
	}
	return todo, nil
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, input string) (*models.Todo, error) {
	todo := &models.Todo{ID: input}
	res := r.DB.First(todo)
	if err := res.Error; err != nil {
		return nil, err
	}
	res = r.DB.Delete(todo)
	if err := res.Error; err != nil {
		return nil, err
	}
	return todo, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*models.User, error) {
	user := &models.User{
		ID:   util.CreateUniqueID(),
		Name: input.Name,
	}
	res := r.DB.Create(user)
	if err := res.Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*models.Todo, error) {
	var todos []*models.Todo
	res := r.DB.Find(&todos)
	if err := res.Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	var users []*models.User
	res := r.DB.Find(&users)
	if err := res.Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *queryResolver) Todo(ctx context.Context, id string) (*models.Todo, error) {
	todo := &models.Todo{ID: id}
	res := r.DB.First(&todo)
	if err := res.Error; err != nil {
		return nil, err
	}
	return todo, nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
	user := &models.User{ID: id}
	res := r.DB.First(user)
	if err := res.Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *todoResolver) User(ctx context.Context, obj *models.Todo) (*models.User, error) {
	user := &models.User{ID: obj.UserID}
	res := r.DB.First(user)
	if err := res.Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userResolver) Todos(ctx context.Context, obj *models.User) ([]*models.Todo, error) {
	var todos []*models.Todo
	res := r.DB.Where("user_id = ?", obj.ID).Find(&todos)
	if err := res.Error; err != nil {
		return nil, err
	}
	return todos, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
