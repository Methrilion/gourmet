package mapper

import (
	"github.com/methrilion/gourmet/pkg/svc-storage-writer/model"
	pbm "github.com/methrilion/gourmet/proto/model"
)

func ProductToPB(product *model.Product) *pbm.Product {
	return &pbm.Product{
		Id:          product.ID,
		Name:        product.Name,
		Description: product.Description,
	}
}
func PBToProduct(pbProduct *pbm.Product) *model.Product {
	return &model.Product{
		ID:          pbProduct.Id,
		Name:        pbProduct.Name,
		Description: pbProduct.Description,
	}
}
