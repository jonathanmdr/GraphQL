package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/jonathanmdr/GraphQL/graph/generated"
	"github.com/jonathanmdr/GraphQL/graph/model"
)

// Courses is the resolver for the courses field.
func (r *categoryResolver) Courses(ctx context.Context, obj *model.Category) ([]*model.Course, error) {
	courses, err := r.CourseDB.FindByCategoryID(obj.ID)
	if err != nil {
		return nil, err
	}
	var coursesModel []*model.Course
	for _, course := range courses {
		var courseModel = &model.Course{
			ID:          course.ID,
			Name:        course.Name,
			Description: &course.Description,
		}
		coursesModel = append(coursesModel, courseModel)
	}
	return coursesModel, nil
}

// CreateCategory is the resolver for the createCategory field.
func (r *mutationResolver) CreateCategory(ctx context.Context, input model.NewCategory) (*model.Category, error) {
	category, err := r.CategoryDB.Create(input.Name, *input.Description)
	if err != nil {
		return nil, err
	}
	return &model.Category{
		ID:          category.ID,
		Name:        category.Name,
		Description: &category.Description,
	}, nil
}

// CreateCourse is the resolver for the createCourse field.
func (r *mutationResolver) CreateCourse(ctx context.Context, input model.NewCourse) (*model.Course, error) {
	course, err := r.CourseDB.Create(input.Name, *input.Description, input.CategoryID)
	if err != nil {
		return nil, err
	}
	return &model.Course{
		ID:          course.ID,
		Name:        course.Name,
		Description: &course.Description,
	}, nil
}

// Categories is the resolver for the categories field.
func (r *queryResolver) Categories(ctx context.Context) ([]*model.Category, error) {
	categories, err := r.CategoryDB.FindAll()
	if err != nil {
		return nil, err
	}
	var categoriesModel []*model.Category
	for _, category := range categories {
		var categoryModel = &model.Category{
			ID:          category.ID,
			Name:        category.Name,
			Description: &category.Description,
		}
		categoriesModel = append(categoriesModel, categoryModel)
	}
	return categoriesModel, nil
}

// Courses is the resolver for the courses field.
func (r *queryResolver) Courses(ctx context.Context) ([]*model.Course, error) {
	courses, err := r.CourseDB.FindAll()
	if err != nil {
		return nil, err
	}
	var coursesModel []*model.Course
	for _, course := range courses {
		var courseModel = &model.Course{
			ID:          course.ID,
			Name:        course.Name,
			Description: &course.Description,
		}
		coursesModel = append(coursesModel, courseModel)
	}
	return coursesModel, nil
}

// Category returns generated.CategoryResolver implementation.
func (r *Resolver) Category() generated.CategoryResolver { return &categoryResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type categoryResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
