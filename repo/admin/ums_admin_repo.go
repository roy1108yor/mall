package adminrepo

import "github.com/kalougata/mall/pkg/data"

type umsAdminRepo struct {
	*data.Data
}

type UmsAdminRepo interface {
}

func NewUmsAdminRepo(data *data.Data) UmsAdminRepo {
	return &umsAdminRepo{Data: data}
}
