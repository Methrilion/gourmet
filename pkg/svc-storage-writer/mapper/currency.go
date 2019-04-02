package mapper

import (
	"github.com/methrilion/gourmet/pkg/svc-storage-writer/model"
	pbm "github.com/methrilion/gourmet/proto/model"
)

func CurrencyToPB(currency *model.Currency) *pbm.Currency {
	return &pbm.Currency{
		Id:   currency.ID,
		Name: currency.Name,
		Code: currency.Code,
	}
}

func PBToCurrency(pbCurrency *pbm.Currency) *model.Currency {
	return &model.Currency{
		ID:   pbCurrency.Id,
		Name: pbCurrency.Name,
		Code: pbCurrency.Code,
	}
}
