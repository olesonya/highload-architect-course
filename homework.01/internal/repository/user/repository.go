package user

import (
	"sync"

	"github.com/olesonya/highload-architect-course/homework.01/internal/database"
	def "github.com/olesonya/highload-architect-course/homework.01/internal/repository"
)

var _ def.UserRepository = (*repository)(nil)

type repository struct {
	db *database.Database
	m  sync.RWMutex
}

func NewRepository(db *database.Database) *repository {
	return &repository{
		db: db,
	}
}
