package employee

type Employee struct {
	Name   string
	Salary int64
}

func CreateEmployee(name string, salary int64) Employee {
	return Employee{
		Name:   name,
		Salary: salary,
	}
}
func (e *Employee) Raise() {
	if e.Salary < 20000 {
		e.Salary = e.Salary + (e.Salary / 10)
	}
}
