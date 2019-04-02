package mapper

import (
	"github.com/methrilion/gourmet/pkg/svc-storage-writer/model"
	pbm "github.com/methrilion/gourmet/proto/model"
)

func MethodToPB(method *model.Method) *pbm.Method {
	return &pbm.Method{
		Id:   method.ID,
		Name: method.Name,
	}
}
func PBToMethod(pbMethod *pbm.Method) *model.Method {
	return &model.Method{
		ID:   pbMethod.Id,
		Name: pbMethod.Name,
	}
}
