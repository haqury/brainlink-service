package service

import (
	"context"
	"github.com/lihao1988/php2go/array"
	"v1/entity"
	"v1/repository"
)

type IBuilderMap interface {
	Get(ctx context.Context) (*entity.Map, error)
}
type BuilderMap struct {
	r                     repository.IRepository
	systemMouseRepository repository.ISystemMouseRepository
}

func NewBuilderMap(r repository.IRepository, systemMouseRepository repository.ISystemMouseRepository) IBuilderMap {
	return &BuilderMap{r, systemMouseRepository}
}

//func (s *BrainlinkService) Get(ctx context.Context, id int) (*entity.Test, error) {
//	return s.r.Get(ctx, id)
//}

func (s *BuilderMap) Get(ctx context.Context) (*entity.Map, error) {
	events, err := s.getEvents(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetMap(events, &entity.BuildConfig{})
}

func (s *BuilderMap) GetMap(e []*entity.Event, config *entity.BuildConfig) (*entity.Map, error) {
	m := entity.Map{}
	for _, event := range e {
		m.AddEvent(*event)
	}
	return m, nil
}

func (s *BuilderMap) getEvents(ctx context.Context) ([]*entity.Event, error) {
	events := []*entity.Event{}
	event := entity.Event{}
	m, err := s.r.List(ctx)
	ids := []int64{}
	array.Column(ids, m, "SystemMouseId", "")
	sm, err := s.systemMouseRepository.ListByIds(ctx, ids)
	for _, model := range m {
		dto := &entity.EegDto{
			Id:     0,
			Input:  *model,
			System: *sm[*model.SystemMouseId],
		}
		if s.setEvent(dto, &event) && event.EegDto != nil {
			events = append(events, &event)
		}
	}
	if err != nil {
		return nil, err
	}

	return events, nil
}

func (s *BuilderMap) setEvent(model *entity.EegDto, event *entity.Event) bool {
	if event == nil || event.EegDto == nil {
		event = &entity.Event{
			Id:     0,
			EegDto: []*entity.EegDto{model},
		}
	}

	e := event.EegDto[len(event.EegDto)]
	if model.System.ToX-e.System.ToX > 0 ||
		model.System.ToY-e.System.ToY > 0 {
		event.EegDto = append(event.EegDto, model)
		return true
	}
	return false
}

//
//func (s *BrainlinkService) DeleteTest(ctx context.Context, id int) (bool, error) {
//	return s.r.Delete(ctx, id)
//}
