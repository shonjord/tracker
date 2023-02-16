package action

import (
	"fmt"
	"strconv"

	"github.com/shonjord/tracker/internal/pkg/domain/entity"
	"github.com/shonjord/tracker/internal/pkg/domain/query"
	"github.com/shonjord/tracker/internal/pkg/presentation/rest"
)

type (
	computersGetter interface {
		GetAllComputers(*query.GetAllComputers) ([]*entity.Computer, error)
	}
	GetAllComputers struct {
		getter computersGetter
	}
)

// NewGetAllComputers returns a new instance of this action..
func NewGetAllComputers(g computersGetter) *GetAllComputers {
	return &GetAllComputers{
		getter: g,
	}
}

// Handle returns all computers using a limit in case provided.
func (a *GetAllComputers) Handle(req *rest.Request, res *rest.Response) error {
	var (
		cmd   = new(query.GetAllComputers)
		limit = 0
	)

	cmd.Limit = limit

	if req.HasQuery("limit") {
		strLimit := req.GetQuery("limit")
		limit, err := strconv.Atoi(strLimit)
		if err != nil {
			return res.BadRequest(fmt.Errorf("%s is not a valid integer for limit", strLimit))
		}

		cmd.Limit = limit
	}

	computers, err := a.getter.GetAllComputers(cmd)
	if err != nil {
		return nil
	}

	return res.WriteStruct(computers)
}
