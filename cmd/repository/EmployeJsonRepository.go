package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"ksi.com/gestion_app/cmd/entities"
)

// EmployeJSONRepository est responsable de l'accès aux données JSON (sauvegarde et récupération).
type EmployeJSONRepository struct {
	filename string
}

// NewEmployeJSONRepository crée une nouvelle instance de EmployeJSONRepository.
func NewEmployeJSONRepository(filename string) *EmployeJSONRepository {
	return &EmployeJSONRepository{
		filename: filename,
	}
}

// GetEmployeJSON récupère la liste des employés depuis le fichier JSON.
func (repo *EmployeJSONRepository) GetEmployeJSON() ([]*entities.Employe, error) {
	var employeeList []*entities.Employe

	// Ouvre le fichier JSON
	file, err := os.Open(repo.filename)
	if err != nil {
		return nil, fmt.Errorf("impossible d'ouvrir le fichier JSON: %w", err)
	}
	defer file.Close()

	// Lit le contenu du fichier JSON
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("impossible de lire le contenu du fichier JSON: %w", err)
	}

	// Décode le contenu JSON dans la liste des employés
	err = json.Unmarshal(byteValue, &employeeList)
	if err != nil {
		return nil, fmt.Errorf("impossible de décoder le contenu JSON: %w", err)
	}

	return employeeList, nil
}

// SaveJSON sauvegarde un nouvel employé dans le fichier JSON.
func (repo *EmployeJSONRepository) SaveJSON(employe *entities.Employe) error {
	// Récupère la liste actuelle des employés
	employeeList, err := repo.GetEmployeJSON()
	if err != nil {
		return fmt.Errorf("impossible de récupérer la liste des employés: %w", err)
	}

	// Ajoute le nouvel employé à la liste
	employeeList = append(employeeList, employe)

	// Convertit la liste des employés en format JSON
	jsonData, err := json.MarshalIndent(employeeList, "", "    ")
	if err != nil {
		return fmt.Errorf("impossible de convertir la liste des employés en format JSON: %w", err)
	}

	// Écrit les données JSON dans le fichier
	err = ioutil.WriteFile(repo.filename, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("impossible d'écrire les données JSON dans le fichier: %w", err)
	}

	return nil
}
