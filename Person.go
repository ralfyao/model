package model

type Student struct {
	name  string
	score float32
}

func (s *Student) GetName() string {
	return s.name
}

func (s *Student) SetName(name string) {
	s.name = name
}
