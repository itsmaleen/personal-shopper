package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/gofrs/uuid"
	"github.com/itsmaleen/personal-shopper/backend/graph/generated"
	"github.com/itsmaleen/personal-shopper/backend/graph/model"
)

func (r *mutationResolver) CreateImageData(ctx context.Context, input model.NewImageData) (*model.ImageData, error) {
	imageData := &model.ImageData{
		ID:   uuid.Must(uuid.NewV4()).String(),
		Name: input.Name,
		URI:  "https://bit.ly/3mCSn2i",
	}

	_, err := r.DB.Model(imageData).Insert()
	if err != nil {
		return nil, fmt.Errorf("error inserting imageData: %v", err)
	}

	return imageData, nil
}

func (r *mutationResolver) UploadImage(ctx context.Context, input model.Image) (bool, error) {
	userFileName := fmt.Sprintf("%v-%v", input.ImageDataID, input.File.Filename)

	stream, readErr := io.ReadAll(input.File.File)
	if readErr != nil {
		fmt.Printf("error from file %v", readErr)
	}

	imageURI := fmt.Sprintf("../images/%v", userFileName)
	fileErr := os.WriteFile(imageURI, stream, 0644)
	if fileErr != nil {
		return false, fmt.Errorf("error saving file: %v", fileErr)
	}

	imageData, imageDataErr := r.GetImageByField("id", *input.ImageDataID)
	if imageDataErr != nil {
		return false, fmt.Errorf("error getting imageData: %v", imageDataErr)
	}

	imageData.URI = imageURI
	_, updateErr := r.UpdateImageData(imageData)
	if updateErr != nil {
		return false, fmt.Errorf("error updating imageData: %v", updateErr)
	}

	return true, nil
}

func (r *mutationResolver) CreateTag(ctx context.Context, input model.NewTag) (*model.Tag, error) {
	tag := &model.Tag{
		ID:    uuid.Must(uuid.NewV4()).String(),
		Name:  input.Name,
		Color: input.Color,
	}

	_, err := r.DB.Model(tag).Insert()
	if err != nil {
		return nil, fmt.Errorf("error inserting tag: %v", err)
	}

	return tag, nil
}

func (r *mutationResolver) AssignTag(ctx context.Context, input model.AddTag) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Tags(ctx context.Context) ([]*model.Tag, error) {
	var tags []*model.Tag
	err := r.DB.Model(&tags).Select()
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func (r *queryResolver) Images(ctx context.Context) ([]*model.ImageData, error) {
	var imageDatas []*model.ImageData
	err := r.DB.Model(&imageDatas).Select()
	if err != nil {
		return nil, err
	}
	return imageDatas, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
