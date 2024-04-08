package mapping

import (
	"encoding/json"

	"ksi.com/gestion_app/cmd/entities"
	"ksi.com/gestion_app/cmd/mapping/dto"
)

type MapperImpl struct{}

func (m MapperImpl) MapEmployeToJSONObject(employe *entities.Employe) map[string]interface{} {
	employeeDetails := make(map[string]interface{})
	employeeDetails["id"] = employe.ID
	employeeDetails["nom"] = employe.Nom
	employeeDetails["prenom"] = employe.Prenom
	employeeDetails["salaire"] = employe.Salaire
	employeeDetails["nbreAbsence"] = employe.NbreAbsence
	employeeDetails["modePaiement"] = employe.ModePaiement
	employeeObject := make(map[string]interface{})
	employeeObject["employee"] = employeeDetails
	return employeeObject
}

func (m MapperImpl) MapEmployeToJSON(employeeList []*entities.Employe) string {
	var employeesDto []*dto.EmployeDto
	for _, emp := range employeeList {
		empDto := dto.NewEmployeDto(emp)
		employeesDto = append(employeesDto, empDto)
	}

	jsonData, _ := json.MarshalIndent(employeesDto, "", "    ")
	return string(jsonData)
}
