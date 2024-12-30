package service

import (
	"context"
	"database/sql"

	"github.com/fajarherdian22/topo-api/exception"
	"github.com/fajarherdian22/topo-api/repository"
	"github.com/fajarherdian22/topo-api/web"
)

type RanServiceImpl struct {
	q *repository.Queries
}

func NewRanService(q *repository.Queries) *RanServiceImpl {
	return &RanServiceImpl{
		q: q,
	}
}

func (service *RanServiceImpl) ListLevel(ctx context.Context, arg string) ([]string, error) {
	payload, err := service.q.ListDistinctNamesByLevel(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			return payload, exception.NewBadRequestError(err.Error())
		}
		return payload, exception.NewInternalError(err.Error())
	}
	return payload, nil
}

func (service *RanServiceImpl) GetAllData(ctx context.Context, arg string) ([]web.RanResponse, error) {
	var resp []web.RanResponse
	payload, err := service.q.GetAllByLevel(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			return resp, exception.NewNotFoundError(err.Error())
		}
		return resp, err
	}
	return web.NewRanResponses(payload), nil
}

func (service *RanServiceImpl) GetByLevel(ctx context.Context, arg repository.GetByLevelAndNameParams) ([]web.RanResponse, error) {
	var resp []web.RanResponse
	payload, err := service.q.GetByLevelAndName(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			return resp, exception.NewNotFoundError(err.Error())
		}
		return resp, err
	}
	return web.NewRanResponses(payload), nil
}

func (service *RanServiceImpl) GetByReference(ctx context.Context, arg repository.GetByLevelAndReferenceParams) ([]web.RanResponse, error) {
	var resp []web.RanResponse
	payload, err := service.q.GetByLevelAndReference(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			return resp, exception.NewNotFoundError(err.Error())
		}
		return resp, err
	}
	return web.NewRanResponses(payload), nil
}
