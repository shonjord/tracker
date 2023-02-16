package handler

import (
	"github.com/google/uuid"
	"github.com/shonjord/tracker/internal/pkg/domain/entity"
)

type (
	computerFinderUpdater interface {
		computerFinder
		computerUpdater
	}
	computerFinder interface {
		FindOneByUUID(uuid.UUID) (*entity.Computer, error)
		FindMany(int) ([]*entity.Computer, error)
	}
	computerSaver interface {
		Save(*entity.Computer) error
	}
	computerUpdater interface {
		Update(*entity.Computer) error
	}
	employeeFinder interface {
		FindOneByUUID(uuid.UUID) (*entity.Employee, error)
	}
	adminNotifier interface {
		Notify(computer *entity.Computer) error
	}
)
