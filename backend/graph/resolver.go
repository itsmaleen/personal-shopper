package graph

import (
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/itsmaleen/personal-shopper/backend/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB *pg.DB
}

func (r *mutationResolver) GetImageByField(field, value string) (*model.ImageData, error) {
	imageData := &model.ImageData{}
	err := r.DB.Model(imageData).Where(fmt.Sprintf("%v = ?", field), value).First()
	if err != nil {
		return nil, fmt.Errorf("error getting imageData: %v", err)
	}

	return imageData, nil
}

func (r *mutationResolver) UpdateImageData(imageData *model.ImageData) (*model.ImageData, error) {
	_, err := r.DB.Model(imageData).Where("id = ?", imageData.ID).Update()
	if err != nil {
		return nil, fmt.Errorf("error updating imageData: %v", err)
	}

	return imageData, nil
}
