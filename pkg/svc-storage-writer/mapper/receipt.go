package mapper

import (
	"github.com/methrilion/gourmet/pkg/svc-storage-writer/model"
	pbm "github.com/methrilion/gourmet/proto/model"
)

func ReceiptToPB(receipt *model.Receipt) *pbm.Receipt {
	return &pbm.Receipt{
		Id:         receipt.ID,
		EmployeeId: receipt.EmployeeID,
		LocationId: receipt.LocationID,
		MethodId:   receipt.MethodID,
		Datetime:   receipt.Datetime,
	}
}
func PBToReceipt(pbReceipt *pbm.Receipt) *model.Receipt {
	return &model.Receipt{
		ID:         pbReceipt.Id,
		EmployeeID: pbReceipt.EmployeeId,
		LocationID: pbReceipt.LocationId,
		MethodID:   pbReceipt.MethodId,
		Datetime:   pbReceipt.Datetime,
	}
}
