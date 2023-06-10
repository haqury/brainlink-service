package service

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"v1/entity"
	"v1/repository"
)

type IBuilderMap interface {
	Get(ctx context.Context) ([]*entity.Event, error)
	Handle(w http.ResponseWriter, r *http.Request)
	HandleEvent(w http.ResponseWriter, r *http.Request)
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

func (s BuilderMap) Handle(w http.ResponseWriter, r *http.Request) {
	// Парсим тело запроса
	_, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body",
			http.StatusInternalServerError)
	}
	defer r.Body.Close()

	// Выводим тело запроса в консоль
	m, _ := s.Get(context.Background())
	d, _ := json.Marshal(m)
	w.Write(d)
}

func (s BuilderMap) HandleEvent(w http.ResponseWriter, r *http.Request) {
	// Парсим тело запроса
	_, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body",
			http.StatusInternalServerError)
	}
	defer r.Body.Close()

	// Выводим тело запроса в консоль
	m, _ := s.GetEvents(context.Background())
	d, _ := json.Marshal(m)
	w.Write(d)
}

func (s *BuilderMap) GetEvents(ctx context.Context) ([]*entity.Event, error) {
	events := []*entity.Event{}
	event := &entity.Event{}
	m, err := s.r.ListUseEvent(ctx)
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
		dto.EventName = *model.EventName
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

func (s *BuilderMap) GetMap(e []*entity.Event, config *entity.BuildConfig) (*entity.Map, error) {
	m := entity.Map{}
	for _, event := range e {
		m.AddEvent(*event)
	}
	return &m, nil
}

func (s *BuilderMap) Get(ctx context.Context) ([]*entity.Event, error) {
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
		dto.EventName = *model.EventName
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

func (s *BuilderMap) setEvent(model entity.EegDto, event *entity.Event) *entity.Event {
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
