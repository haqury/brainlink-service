package service

import (
	"context"
	"v1/entity"
	"v1/repository"
)

type IMetaGeneretor interface {
	Sync() (error)
}

type MetaGeneretor struct {
	r                     repository.IRepository
	systemMouseRepository repository.ISystemMouseRepository
}

//
//func NewMetaGeneretor(r repository.IRepository, systemMouseRepository repository.ISystemMouseRepository) IMetaGeneretor {
//	return &MetaGeneretor{r, systemMouseRepository}
//}

//func (s *BrainlinkService) Get(ctx context.Context, id int) (*entity.Test, error) {
//	return s.r.Get(ctx, id)
//}
//
//func (s *MetaGeneretor) Sync() (error) {
//	events, _ := s.getEvents(context.Background())
//	for _, e := range events {
//
//	}
//}

func (s *MetaGeneretor) getEvents(ctx context.Context) ([]*entity.Event, error) {
	events := []*entity.Event{}
	event := &entity.Event{}
	m, err := s.r.List(ctx)
	ids := []int64{}
	for _, mm := range m {
		ids = append(ids, *mm.SystemMouseId)
	}
	sm, err := s.systemMouseRepository.ListByIds(ctx, ids)
	for _, model := range m {
		dto := &entity.EegDto{
			Id:     0,
			Input:  *model,
			System: *sm[*model.SystemMouseId],
		}
		event = s.setEvent(*dto, event)
		if event != nil {
			events = append(events, event)
		}
	}
	if err != nil {
		return nil, err
	}

	return events, nil
}

func (s *MetaGeneretor) setEvent(model entity.EegDto, event *entity.Event) *entity.Event {
	if event == nil || event.EegDto == nil {
		event = &entity.Event{
			Id:     0,
			EegDto: []entity.EegDto{model},
		}
	}

	e := event.EegDto[len(event.EegDto)-1]
	if model.System.ToX*e.System.ToX == 0 ||
		model.System.ToY*e.System.ToY == 0 {
		return nil
	}

	if model.System.ToX*e.System.ToX > 0 ||
		model.System.ToY*e.System.ToY > 0 {
		event.EegDto = append(event.EegDto, model)
		return event
	}

	event = &entity.Event{
		Id:     0,
		EegDto: []entity.EegDto{model},
	}
	event.EegDto = append(event.EegDto, model)
	return event
}

//
//func (s *BrainlinkService) DeleteTest(ctx context.Context, id int) (bool, error) {
//	return s.r.Delete(ctx, id)
//}
