package mapper

import (
	"github.com/methrilion/gourmet/pkg/svc-storage-writer/model"
	pbm "github.com/methrilion/gourmet/proto/model"
)

func LocationToPB(location *model.Location) *pbm.Location {
	return &pbm.Location{
		Id:          location.ID,
		Name:        location.Name,
		Description: location.Description,
		CurrencyId:  location.CurrencyID,
		Status:      location.Status,
	}
}

func PBToLocation(pbLocation *pbm.Location) *model.Location {
	return &model.Location{
		ID:          pbLocation.Id,
		Name:        pbLocation.Name,
		Description: pbLocation.Description,
		CurrencyID:  pbLocation.CurrencyId,
		Status:      pbLocation.Status,
	}
}
