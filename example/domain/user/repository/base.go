package repository

import (
	"github.com/ravielze/oculi/example/model/dao"
	"github.com/ravielze/oculi/example/resources"
	"github.com/ravielze/oculi/request"
)

type (
	Repository interface {
		Create(req request.Context, user dao.User) (dao.User, error)
		GetByUsername(req request.Context, username string) (dao.User, error)
		GetByID(req request.Context, userId uint64) (dao.User, error)
		Update(req request.Context, userId uint64, request map[string]interface{}) error
	}

	repository struct {
		resource resources.Resource
	}
)

func New(r resources.Resource) Repository {
	return &repository{
		resource: r,
	}
}