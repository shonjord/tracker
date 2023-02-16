package action

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/shonjord/tracker/internal/pkg/domain/entity"
	"github.com/shonjord/tracker/internal/pkg/domain/query"
	"github.com/shonjord/tracker/internal/pkg/presentation/rest"
)

type (
	computerGetter interface {
		GetComputer(*query.GetComputerByUUID) (*entity.Computer, error)
	}
	GetComputerByUUID struct {
		getter computerGetter
	}
)

// NewGetComputerByUUID returns a new instance of this action.
func NewGetComputerByUUID(f computerGetter) *GetComputerByUUID {
	return &GetComputerByUUID{
		getter: f,
	}
}

// Handle returns a computer by its UUID.
func (a *GetComputerByUUID) Handle(req *rest.Request, res *rest.Response) error {
	var (
		cmd = new(query.GetComputerByUUID)
	)

	uuidParam := req.GetParam(computerUUIDParam)
	computerUUID, err := uuid.Parse(uuidParam)
	if err != nil {
		return res.BadRequest(&rest.Error{
			Message:               fmt.Sprintf("invalid UUID provided: %s", uuidParam),
			HTTPStatusCounterpart: http.StatusBadRequest,
		})
	}

	cmd.ComputerUUID = computerUUID

	computer, err := a.getter.GetComputer(cmd)
	if err != nil {
		return err
	}

	return res.WriteStruct(computer)
}
