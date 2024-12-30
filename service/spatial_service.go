package service

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/fajarherdian22/topo-api/exception"
	"github.com/fajarherdian22/topo-api/repository"
)

type KabKotaServiceImpl struct {
	repo *repository.KabKotaRepository
}

func NewKabKotaService(repo *repository.KabKotaRepository) *KabKotaServiceImpl {
	return &KabKotaServiceImpl{
		repo: repo,
	}
}

func (service *KabKotaServiceImpl) GetAllSpatial(ctx context.Context) ([]repository.KabKota, error) {
	payload, err := service.repo.GetAllSpatial(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return payload, exception.NewBadRequestError(err.Error())
		}
		return payload, exception.NewInternalError(err.Error())
	}
	return payload, err
}

func (service *KabKotaServiceImpl) GetSpatialByFilter(ctx context.Context, arg repository.GetSpatialLv) ([]repository.KabKota, error) {
	payload, err := service.repo.GetSpatialByFilter(ctx, arg)
	fmt.Println(err)
	if err != nil {
		if err == sql.ErrNoRows {
			return payload, exception.NewNotFoundError(err.Error())
		}
		return payload, exception.NewNotFoundError(err.Error())
	}
	return payload, err
}
