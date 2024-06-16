package whisky

import (
	"context"
	"github.com/Massad/gin-boilerplate/db"
	"github.com/Massad/gin-boilerplate/generated/models"
	"github.com/Massad/gin-boilerplate/models/whisky"
)

type Repositories interface {
	Create(context.Context, *whisky.Whisky) (*models.Whisky, error)
}

type repositories struct {
}

func NewRepositories() Repositories {
	return &repositories{}
}

func (r repositories) Create(ctx context.Context, whisky *whisky.Whisky) (*models.Whisky, error) {
	model := toBoiler(whisky)
	err := model.Insert(ctx, db.GetDB(), model)
	if err != nil {
		return nil, err
	}
	err = model.Reload(ctx, db.GetDB())
	if err != nil {
		return nil, err
	}
	return model, nil
}

func toBoiler(whisky *whisky.Whisky) *models.Whisky {
	return &models.Whisky{
		Strength:  whisky.Strength,
		Size:      whisky.Size,
		CreatedAt: whisky.CreatedAt,
		UpdatedAt: whisky.UpdatedAt,
	}
}
