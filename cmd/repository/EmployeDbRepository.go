package repository

import (
	"fmt"

	"ksi.com/gestion_app/cmd/entities"
)

type EmployeDBRepository struct{}

func NewEmployeDBRepository() *EmployeDBRepository {
	return &EmployeDBRepository{}
}

func (repo *EmployeDBRepository) SaveDB(employe *entities.Employe) {
	fmt.Println("Ajout dans une Base de Donnee")
}
