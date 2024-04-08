package mapping

import (
	"ksi.com/gestion_app/cmd/entities"
)

type Mapper interface {
	MapEmployeToJSONObject(employe *entities.Employe) map[string]interface{}
	MapEmployeToJSON(employeeList []*entities.Employe) string
}
