package repository

import (
	"github.com/LKarlon/test_task/pkg/models"
	"github.com/jmoiron/sqlx"
)

type Storage interface {
	GetInfo(serverID int) (models.Servers, error)
	DeleteInfo(serverID int) error
}

type storage struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Storage {
	return &storage{
		db: db,
	}
}

func (s *storage) GetInfo(serverID int) (res models.Servers, err error) {
	res = models.Servers{}
	tmp := models.Interface{}
	rows, err := s.db.Query(qGetInfo, serverID)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&tmp.Name, &tmp.ServerName, &tmp.Values)
		if err != nil {
			return
		}
		res.Interfaces = append(res.Interfaces, tmp)
	}
	return
}

func (s *storage) DeleteInfo(serverID int) (err error) {
	_, err = s.db.Exec(qDeleteInfo, serverID)
	return
}
