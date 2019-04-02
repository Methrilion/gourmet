package mapper

import (
	"github.com/methrilion/gourmet/pkg/svc-storage-writer/model"
	pbm "github.com/methrilion/gourmet/proto/model"
)

func RateOfExchangeToPB(rateOfExchange *model.RateOfExchange) *pbm.RateOfExchange {
	return &pbm.RateOfExchange{
		Id:     rateOfExchange.ID,
		FromId: rateOfExchange.FromID,
		ToId:   rateOfExchange.ToID,
		Price:  rateOfExchange.Price,
	}
}

func PBToRateOfExchange(pbRateOfExchange *pbm.RateOfExchange) *model.RateOfExchange {
	return &model.RateOfExchange{
		ID:     pbRateOfExchange.Id,
		FromID: pbRateOfExchange.FromId,
		ToID:   pbRateOfExchange.ToId,
		Price:  pbRateOfExchange.Price,
	}
}
