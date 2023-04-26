package service

import (
	"context"
	"v1/entity"
	"v1/repository"
)

type IBrainlinkService interface {
	//Get(ctx context.Context, id int) (*entity.EegDto, error)
	List(ctx context.Context) ([]*entity.EegHistoryModel, error)
	Add(ctx context.Context, m *entity.EegDto) (*entity.EegDto, error)
	//DeleteTest(ctx context.Context, id int) (bool, error)
}
type BrainlinkService struct {
	r                     repository.IRepository
	systemMouseRepository repository.ISystemMouseRepository
}

func NewBrainlinkService(r repository.IRepository, systemMouseRepository repository.ISystemMouseRepository) IBrainlinkService {
	return &BrainlinkService{r, systemMouseRepository}
}

//func (s *BrainlinkService) Get(ctx context.Context, id int) (*entity.Test, error) {
//	return s.r.Get(ctx, id)
//}

func (s *BrainlinkService) List(ctx context.Context) ([]*entity.EegHistoryModel, error) {
	return s.r.List(ctx)
}

func (s *BrainlinkService) Add(ctx context.Context, m *entity.EegDto) (*entity.EegDto, error) {
	system, err := s.systemMouseRepository.Add(ctx, &m.System)
	if system != nil {
		m.System = *system
	}
	if err != nil {
		return nil, err
	}
	m.Input.SystemMouseId = &m.System.Id
	input, err := s.r.Add(ctx, &m.Input)
	if input != nil {
		m.Input = *input
	}
	if err != nil {
		return nil, err
	}
	return m, nil
}

//
//func (s *BrainlinkService) DeleteTest(ctx context.Context, id int) (bool, error) {
//	return s.r.Delete(ctx, id)
//}
