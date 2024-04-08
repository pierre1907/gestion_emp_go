package entities

// Employe représente un employé
type Employe struct {
    ID           int64
    Nom          string
    Prenom       string
    Salaire      float64
    NbreAbsence  int64
    ModePaiement string
    Retenu       float64
    Prime        float64
    SalaireBrut  float64
}