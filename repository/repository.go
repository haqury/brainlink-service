package repository

import (
	"context"
	"v1/db"
	"v1/entity"
)

type IRepository interface {
	//Get(ctx context.Context, id int) (*entity.Test, error)
	List(ctx context.Context) ([]*entity.EegHistoryModel, error)
	ListUseEvent(ctx context.Context) ([]*entity.EegHistoryModel, error)
	Add(ctx context.Context, t *entity.EegHistoryModel) (*entity.EegHistoryModel, error)
	//Delete(ctx context.Context, id int) (bool, error)
}
type Repository struct {
	db *db.DB
}

func NewRepository(d *db.DB) *Repository {
	return &Repository{db: d}
}

//func (r *Repository) Get(ctx context.Context, id int) (*entity.Test, error) {
//	e := entity.Test{}
//
//	err := r.db.Client.
//		QueryRow(ctx, "SELECT id FROM test WHERE id = $1 AND deleted_at IS NULL", id).
//		Scan(&e.Id)
//
//	if err != nil {
//		return nil, err
//	}
//
//	return &e, nil
//}

func (r *Repository) List(ctx context.Context) ([]*entity.EegHistoryModel, error) {
	var d []*entity.EegHistoryModel

	rows, err := r.db.Client.Query(ctx, "SELECT attention, meditation, signal, delta, theta, lowalpha, highalpha, lowbeta, highbeta, lowgamma, highgamma, system_mouse_id, event_name FROM brainlink.eeg_history WHERE event_name = '' order by id desc limit 300 ")

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var e entity.EegHistoryModel
		err := rows.Scan(
			&e.Attention,
			&e.Meditation,
			&e.Signal,
			&e.Delta,
			&e.Theta,
			&e.LowAlpha,
			&e.HighAlpha,
			&e.LowBeta,
			&e.HighBeta,
			&e.LowGamma,
			&e.HighGamma,
			&e.SystemMouseId,
			&e.EventName,
		)

		if err != nil {
			return nil, err
		}

		d = append(d, &e)
	}

	return d, nil
}

func (r *Repository) ListUseEvent(ctx context.Context) ([]*entity.EegHistoryModel, error) {
	var d []*entity.EegHistoryModel

	rows, err := r.db.Client.Query(ctx, "SELECT attention, meditation, signal, delta, theta, lowalpha, highalpha, lowbeta, highbeta, lowgamma, highgamma, system_mouse_id, event_name FROM brainlink.eeg_history "+
		"WHERE event_name != ''"+
		"order by id desc limit 300 ")

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var e entity.EegHistoryModel
		err := rows.Scan(
			&e.Attention,
			&e.Meditation,
			&e.Signal,
			&e.Delta,
			&e.Theta,
			&e.LowAlpha,
			&e.HighAlpha,
			&e.LowBeta,
			&e.HighBeta,
			&e.LowGamma,
			&e.HighGamma,
			&e.SystemMouseId,
			&e.EventName,
		)

		if err != nil {
			return nil, err
		}

		d = append(d, &e)
	}

	return d, nil
}

func (r *Repository) Add(ctx context.Context, e *entity.EegHistoryModel) (*entity.EegHistoryModel, error) {
	err := r.db.Client.
		QueryRow(ctx, "INSERT INTO brainlink.eeg_history(attention, meditation, signal, delta, theta, lowalpha, highalpha, lowbeta, highbeta, lowgamma, highgamma, system_mouse_id, event_name) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)",
			&e.Attention,
			&e.Meditation,
			&e.Signal,
			&e.Delta,
			&e.Theta,
			&e.LowAlpha,
			&e.HighAlpha,
			&e.LowBeta,
			&e.HighBeta,
			&e.LowGamma,
			&e.HighGamma,
			&e.SystemMouseId,
			e.EventName,
		).
		Scan(
			&e.Attention,
			&e.Meditation,
			&e.Signal,
			&e.Delta,
			&e.Theta,
			&e.LowAlpha,
			&e.HighAlpha,
			&e.LowBeta,
			&e.HighBeta,
			&e.LowGamma,
			&e.HighGamma,
			&e.SystemMouseId,
			e.EventName,
		)

	if err != nil {
		return nil, err
	}

	return e, nil
}

//
//func (r *Repository) Delete(ctx context.Context, id int) (bool, error) {
//	if _, err := r.Get(ctx, id); err != nil {
//		return false, err
//	}
//
//	//_, err := r.db.Client.
//	//	Query(ctx, "UPDATE test SET deleted_at = now() WHERE id = $1", id)
//
//	//if err != nil {
//	//	return false, err
//	//}
//
//	return true, nil
//}
