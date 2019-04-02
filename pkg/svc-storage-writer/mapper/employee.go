package mapper

import (
	"github.com/methrilion/gourmet/pkg/svc-storage-writer/model"
	pbm "github.com/methrilion/gourmet/proto/model"
)

func EmployeeToPB(employee *model.Employee) *pbm.Employee {
	return &pbm.Employee{
		Id:         employee.ID,
		FirstName:  employee.FirstName,
		LastName:   employee.LastName,
		PositionId: employee.PositionID,
		LocationId: employee.LocationID,
	}
}
func PBToEmployee(pbEmployee *pbm.Employee) *model.Employee {
	return &model.Employee{
		ID:         pbEmployee.Id,
		FirstName:  pbEmployee.FirstName,
		LastName:   pbEmployee.LastName,
		PositionID: pbEmployee.PositionId,
		LocationID: pbEmployee.LocationId,
	}
}
