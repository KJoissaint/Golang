package models


type Employee struct {
	ID     int
	Name   string
	Salary float64
}


func (e *Employee) Raise(percent float64) {
	e.Salary += e.Salary * (percent / 100)
}
