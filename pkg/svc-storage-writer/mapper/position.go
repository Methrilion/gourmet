package mapper

import (
	"github.com/methrilion/gourmet/pkg/svc-storage-writer/model"
	pbm "github.com/methrilion/gourmet/proto/model"
)

func PositionToPB(position *model.Position) *pbm.Position {
	return &pbm.Position{
		Id:          position.ID,
		Name:        position.Name,
		Description: position.Description,
	}
}
func PBToPosition(pbPosition *pbm.Position) *model.Position {
	return &model.Position{
		ID:          pbPosition.Id,
		Name:        pbPosition.Name,
		Description: pbPosition.Description,
	}
}
