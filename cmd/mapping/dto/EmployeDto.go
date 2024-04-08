package dto

import (
	"ksi.com/gestion_app/cmd/entities"
)

// EmployeDto représente un DTO pour un employé
type EmployeDto struct {
	ID          int64  `json:"id"`
	Nom         string `json:"nom"`
	Prenom      string `json:"prenom"`
	NbreAbsence int64  `json:"nbreAbsence"`
}

// NewEmployeDto crée un nouveau EmployeDto à partir d'un JSON
func NewEmployeDto(emp *entities.Employe) *EmployeDto {
	//employe := emp["employee"].(map[string]interface{})
	return &EmployeDto{
		ID:          emp.ID,
		Nom:         emp.Nom,
		Prenom:      emp.Prenom,
		NbreAbsence: emp.NbreAbsence,
	}
}
