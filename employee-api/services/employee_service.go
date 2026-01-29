package services

import (
	"employee-api/models"
	"errors"
)

// Interface avec les nouvelles méthodes
type EmployeeService interface {
	GetAll() []models.Employee
	Add(employee models.Employee) models.Employee
	RaiseSalary(id int, percent float64) error
}

type EmployeeServiceImpl struct {
	employees []models.Employee
	nextID    int // Pour générer les IDs automatiquement
}

func NewEmployeeService() EmployeeService {
	return &EmployeeServiceImpl{
		employees: []models.Employee{
			{ID: 1, Name: "Alice", Salary: 5000},
			{ID: 2, Name: "Bob", Salary: 7000},
			{ID: 3, Name: "Charlie", Salary: 6000},
		},
		nextID: 4, // Prochain ID disponible
	}
}

func (s *EmployeeServiceImpl) GetAll() []models.Employee {
	return s.employees
}

// Add ajoute un employé avec un ID généré automatiquement
func (s *EmployeeServiceImpl) Add(employee models.Employee) models.Employee {
	employee.ID = s.nextID
	s.nextID++
	s.employees = append(s.employees, employee)
	return employee
}

// RaiseSalary augmente le salaire d'un employé via pointeur
func (s *EmployeeServiceImpl) RaiseSalary(id int, percent float64) error {
	for i := range s.employees {
		if s.employees[i].ID == id {
			// Utilisation du pointeur pour modifier l'employé
			(&s.employees[i]).Raise(percent)
			return nil
		}
	}
	return errors.New("employé non trouvé")
}
