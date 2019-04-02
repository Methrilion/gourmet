package mapper

import (
	"github.com/methrilion/gourmet/pkg/svc-storage-writer/model"
	pbm "github.com/methrilion/gourmet/proto/model"
)

func PriceToPB(price *model.Price) *pbm.Price {
	return &pbm.Price{
		Id:         price.ID,
		ProductId:  price.ProductID,
		LocationId: price.LocationID,
		Price:      price.Price,
	}
}
func PBToPrice(pbPrice *pbm.Price) *model.Price {
	return &model.Price{
		ID:         pbPrice.Id,
		ProductID:  pbPrice.ProductId,
		LocationID: pbPrice.LocationId,
		Price:      pbPrice.Price,
	}
}
