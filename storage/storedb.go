package storage

import (
	"dummyapi/model"
	"dummyapi/util"
)

// StoreDb defines storage interface
type StoreDb interface {
	CreateUser(user *model.User) (*model.User, error)
	GetUser(id string) (*model.User, error)
	GetPostsByUser(*model.User, *util.Paginator) (*model.List[model.Post], error)
}
