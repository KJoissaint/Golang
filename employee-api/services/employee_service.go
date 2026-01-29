package services

import "employee-api/models"

//ici on définit le contrat pour la gestion des employés
type EmployeeService interface {
	GetAll() []models.Employee
}


type EmployeeServiceImpl struct {
	employees []models.Employee
}


func NewEmployeeService() EmployeeService {
	return &EmployeeServiceImpl{
		employees: []models.Employee{
			{ID: 1, Name: "Alice", Salary: 5000},
			{ID: 2, Name: "Bob", Salary: 7000},
			{ID: 3, Name: "Charlie", Salary: 6000},
		},
	}
}


func (s *EmployeeServiceImpl) GetAll() []models.Employee {
	return s.employees
}
