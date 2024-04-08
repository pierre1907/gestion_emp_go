package main

import (
	"fmt"

	"ksi.com/gestion_app/cmd/entities"
	"ksi.com/gestion_app/cmd/mapping"
	"ksi.com/gestion_app/cmd/services"
)

func main() {
	// Création d'une instance du service Employe avec un MapperImpl
	employeService := services.GetInstance(mapping.MapperImpl{})

	// Création d'un employé en utilisant le design pattern Builder
	emp := &entities.Employe{
		ID:           1,
		Nom:          "KSI",
		Prenom:       "SPE",
		NbreAbsence:  10,
		ModePaiement: "OM",
	}
	fmt.Println(emp)

	// Enregistrement de l'employé
	employeService.Save(emp, entities.JSON)

	// Affichage des données JSON des employés
	fmt.Println(employeService.GetEmployeJSON())
}
