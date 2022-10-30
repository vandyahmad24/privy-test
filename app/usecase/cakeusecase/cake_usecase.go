package cakeusecase

import (
	"context"
	"errors"
	"time"
	"vandyahmad24/privy/app/db/model"
	"vandyahmad24/privy/app/domain/repository"
	"vandyahmad24/privy/app/tracing"

	"github.com/mitchellh/mapstructure"
)

type CakeUsecase struct {
	repository repository.CakeRepository
}

func NewCakeUsecase(repository repository.CakeRepository) *CakeUsecase {
	return &CakeUsecase{repository: repository}
}

func (e *CakeUsecase) CreateCake(ctx context.Context, in interface{}) (interface{}, error) {
	sp := tracing.CreateChildSpan(ctx, string("Interactor"))
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	if in == nil {
		return nil, errors.New("request cannot be nil")
	}
	var inputCake *model.Cake
	err := mapstructure.Decode(in, &inputCake)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, errors.New("request cannot be nil")
	}

	inputCake.CreatedAt = time.Now()
	inputCake.UpdatedAt = time.Now()

	data, err := e.repository.InsertCake(sp, inputCake)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	tracing.LogResponse(sp, data)
	return data, nil
}

func (e *CakeUsecase) GetAllCake(ctx context.Context) (interface{}, error) {
	sp := tracing.CreateChildSpan(ctx, string("Interactor"))
	defer sp.Finish()

	data, err := e.repository.GetAll(sp)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}
	tracing.LogResponse(sp, data)

	return data, nil
}

func (e *CakeUsecase) GetCake(ctx context.Context, id int) (interface{}, error) {
	sp := tracing.CreateChildSpan(ctx, string("Interactor"))
	defer sp.Finish()

	data, err := e.repository.Get(sp, id)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}
	tracing.LogResponse(sp, data)

	return data, nil
}

func (e *CakeUsecase) DeleteCake(ctx context.Context, id int) error {
	sp := tracing.CreateChildSpan(ctx, string("Interactor"))
	defer sp.Finish()

	_, err := e.repository.Get(sp, id)
	if err != nil {
		tracing.LogError(sp, err)
		return errors.New("Cake Not Fund")
	}

	err = e.repository.Delete(sp, id)
	if err != nil {
		tracing.LogError(sp, err)
		return err
	}

	return nil
}

func (e *CakeUsecase) UpdateCake(ctx context.Context, id int, in interface{}) (interface{}, error) {
	sp := tracing.CreateChildSpan(ctx, string("Interactor"))
	defer sp.Finish()
	if in == nil {
		return nil, errors.New("request cannot be nil")
	}
	_, err := e.repository.Get(sp, id)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, errors.New("Cake Not Fund")
	}

	var inputCake *model.Cake
	err = mapstructure.Decode(in, &inputCake)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, errors.New("request cannot be nil")
	}

	inputCake.UpdatedAt = time.Now()

	data, err := e.repository.Update(sp, id, inputCake)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}
	tracing.LogResponse(sp, data)

	return data, nil
}
