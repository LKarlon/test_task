package service

import (
	"github.com/LKarlon/test_task/pkg/models"
)

type storage interface {
	GetInfo(serverID int) (models.Servers, error)
	DeleteInfo(serverID int) error
}

type Service interface {
	GetInfo(serverID int) (res models.Servers, err error)
	DeleteInfo(serverID int) (err error)
}

type service struct {
	storage
}

func NewService(storage storage) Service {
	return &service{storage: storage}
}

func (s *service) GetInfo(serverID int) (res models.Servers, err error) {
	return s.storage.GetInfo(serverID)
}

func (s *service) DeleteInfo(serverID int) (err error) {
	return s.storage.DeleteInfo(serverID)
}
