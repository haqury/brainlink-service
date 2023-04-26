package repository

import (
	"context"
	"v1/db"
	"v1/entity"
)

type ISystemMouseRepository interface {
	//Get(ctx context.Context, id int) (*entity.Test, error)
	//List(ctx context.Context) ([]*entity.Test, error)
	ListByIds(ctx context.Context, ids []int64) ([]*entity.SystemInfo, error)
	Add(ctx context.Context, e *entity.SystemInfo) (*entity.SystemInfo, error)
	//Delete(ctx context.Context, id int) (bool, error)
}
type SystemMouseRepository struct {
	db *db.DB
}

func NewSystemMouseRepository(d *db.DB) *SystemMouseRepository {
	return &SystemMouseRepository{db: d}
}

//
//func (r *Repository) Get(ctx context.Context, id int) (*entity.Test, error) {
//	e := entity.Test{}
//
//	//err := r.db.Client.
//	//	QueryRow(ctx, "SELECT id FROM test WHERE id = $1 AND deleted_at IS NULL", id).
//	//	Scan(&e.Id)
//
//	//if err != nil {
//	//	return nil, err
//	//}
//
//	return &e, nil
//}

func (r *SystemMouseRepository) ListByIds(ctx context.Context, ids []int64) ([]*entity.SystemInfo, error) {
	var d []*entity.SystemInfo

	idss := entity.NewIds(ids)
	rows, err := r.db.Client.Query(ctx, "SELECT id, x, y, tox, toy, endx, endy FROM brainlink.system_mouse"+
		idss.WhereIn("id"))

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var e entity.SystemInfo
		err := rows.Scan(&e.Id, &e.X, &e.Y, &e.ToX, &e.ToY, &e.EndX, &e.EndY)

		if err != nil {
			return nil, err
		}

		d[e.Id] = &e
	}

	return d, nil
}

func (r *SystemMouseRepository) Add(ctx context.Context, e *entity.SystemInfo) (*entity.SystemInfo, error) {
	err := r.db.Client.
		QueryRow(ctx, "INSERT INTO brainlink.system_mouse(x, y, tox, toy, endx, endy) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
			e.X,
			e.Y,
			e.ToX,
			e.ToY,
			e.EndX,
			e.EndY,
		).Scan(&e.Id)

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
