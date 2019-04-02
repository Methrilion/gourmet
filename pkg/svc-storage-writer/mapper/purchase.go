package mapper

import (
	"github.com/methrilion/gourmet/pkg/svc-storage-writer/model"
	pbm "github.com/methrilion/gourmet/proto/model"
)

func PurchaseToPB(purchase *model.Purchase) *pbm.Purchase {
	return &pbm.Purchase{
		Id:        purchase.ID,
		ReceiptId: purchase.ReceiptID,
		PriceId:   purchase.PriceID,
		Amount:    purchase.Amount,
		Price:     purchase.Price,
		Result:    purchase.Result,
	}
}
func PBToPurchase(pbPurchase *pbm.Purchase) *model.Purchase {
	return &model.Purchase{
		ID:        pbPurchase.Id,
		ReceiptID: pbPurchase.ReceiptId,
		PriceID:   pbPurchase.PriceId,
		Amount:    pbPurchase.Amount,
		Price:     pbPurchase.Price,
		Result:    pbPurchase.Result,
	}
}
