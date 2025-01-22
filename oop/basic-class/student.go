package oop

type Student struct {
	FirstName string
	LastName  string
	Weight    int
	Height    int
	Grade     string
}

func (s Student) FullName() string {
	return s.FirstName + " " + s.LastName
}
