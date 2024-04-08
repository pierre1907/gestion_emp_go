package services

import (
	"fmt"
	"ksi.com/gestion_app/cmd/entities"
	"ksi.com/gestion_app/cmd/mapping"
	"ksi.com/gestion_app/cmd/repository"
	//"gestion_app/entities"
	//"gestion_app/mapping"
	//"gestion_app/repository"
)

type EmployeService struct {
	mapper          mapping.Mapper
	employeJSONRepo *repository.EmployeJSONRepository
	employeDBRepo   *repository.EmployeDBRepository
}

var instance *EmployeService

func GetInstance(mapper mapping.Mapper) *EmployeService {
	if instance == nil {
		instance = &EmployeService{
			mapper:          mapper,
			employeJSONRepo: repository.NewEmployeJSONRepository("/data/employe.json"),
			employeDBRepo:   repository.NewEmployeDBRepository(),
		}
	}
	return instance
}

func (s *EmployeService) VirerSalaire(employe *entities.Employe) {
	fmt.Printf("Salaire Virement fait pour %s %s\n", employe.Nom, employe.Prenom)
	fmt.Printf("Salaire Virement fait en %s\n", employe.ModePaiement)
}

func (s *EmployeService) Save(employe *entities.Employe, modeSave entities.ModeSave) {
	if modeSave == entities.JSON {
		maJsonObject := s.mapper.MapEmployeToJSONObject(employe)

		// Convertir maJsonObject en employé, si nécessaire
		employeFromJSON := convertJSONObjectToEmploye(maJsonObject)
		s.employeJSONRepo.SaveJSON(employeFromJSON)
	} else {
		s.employeDBRepo.SaveDB(employe)
	}
}

// Fonction pour convertir un objet JSON en un objet Employe
func convertJSONObjectToEmploye(jsonObj map[string]interface{}) *entities.Employe {
	employe := &entities.Employe{}

	if id, ok := jsonObj["id"].(int64); ok {
		employe.ID = id
	}
	if nom, ok := jsonObj["nom"].(string); ok {
		employe.Nom = nom
	}
	if prenom, ok := jsonObj["prenom"].(string); ok {
		employe.Prenom = prenom
	}
	if salaire, ok := jsonObj["salaire"].(float64); ok {
		employe.Salaire = salaire
	}
	if nbreAbsence, ok := jsonObj["nbreAbsence"].(int64); ok {
		employe.NbreAbsence = nbreAbsence
	}
	if modePaiement, ok := jsonObj["modePaiement"].(string); ok {
		employe.ModePaiement = modePaiement
	}
	if retenu, ok := jsonObj["retenu"].(float64); ok {
		employe.Retenu = retenu
	}
	if prime, ok := jsonObj["prime"].(float64); ok {
		employe.Prime = prime
	}
	if salaireBrut, ok := jsonObj["salaireBrut"].(float64); ok {
		employe.SalaireBrut = salaireBrut
	}

	return employe
}

func (s *EmployeService) GetEmployeJSON() string {
	// Modification pour gérer les deux valeurs retournées par GetEmployeJSON
	employeJson, _ := s.employeJSONRepo.GetEmployeJSON()
	return s.mapper.MapEmployeToJSON(employeJson)
}
