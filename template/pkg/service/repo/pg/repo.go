package pg

import (
	"{{Package}}/pkg/service/repo"
)

type pgRepo struct {
}

// NewRepo new pg repo implimentation
func NewRepo() repo.Repo {
	return &pgRepo{}
}
