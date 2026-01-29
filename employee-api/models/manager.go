package models

// Manager représente un manager qui est aussi un employé avec une équipe
type Manager struct {
	Employee         
	TeamSize int     
}
