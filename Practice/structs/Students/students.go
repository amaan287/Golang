package students

type Student struct {
	studentName string
	marks       int64
}

func CreateStudent(name string, mark int64) Student {
	return Student{
		studentName: name,
		marks:       mark,
	}
}

func (s *Student) AddTen() {
	s.marks = s.marks + 10
}
