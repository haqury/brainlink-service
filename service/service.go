package service

import (
	"context"
	"v1/entity"
	"v1/repository"
)

type IBrainlinkService interface {
	//Get(ctx context.Context, id int) (*entity.EegDto, error)
	List(ctx context.Context) ([]*entity.EegDto, error)
	Add(ctx context.Context, message *entity.EegDto) (*entity.EegDto, error)
	//DeleteTest(ctx context.Context, id int) (bool, error)
}
type BrainlinkService struct {
	r                     repository.IRepository
	systemMouseRepository repository.ISystemMouseRepository
}

func NewBrainlinkService(r repository.IRepository, systemMouseRepository repository.ISystemMouseRepository) *BrainlinkService {
	return &BrainlinkService{r, systemMouseRepository}
}

//func (s *BrainlinkService) Get(ctx context.Context, id int) (*entity.Test, error) {
//	return s.r.Get(ctx, id)
//}

func (s *BrainlinkService) List(ctx context.Context) ([]*entity.EegHistoryModel, error) {
	return s.r.List(ctx)
}

func (s *BrainlinkService) Add(ctx context.Context, m *entity.EegDto) (*entity.EegDto, error) {
	var err error
	m.System, err = s.systemMouseRepository.Add(ctx, m.System)
	if err != nil {
		return nil, err
	}
	m.Input.SystemMouseId = &m.System.Id
	m.Input, err = s.r.Add(ctx, m.Input)
	if err != nil {
		return nil, err
	}
	return m, nil
}

//
//func (s *BrainlinkService) DeleteTest(ctx context.Context, id int) (bool, error) {
//	return s.r.Delete(ctx, id)
//}
