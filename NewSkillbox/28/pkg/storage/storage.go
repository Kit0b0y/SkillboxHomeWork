package storage

import (
	"fmt"

	"github.com/Kit0b0y/SkillboxHomeWork/NewSkillbox/28/pkg/student"
)

type StudentStorage map[string]*student.Student

func NewStudentStorage() StudentStorage {
	return make(map[string]*student.Student)
}
func (ss StudentStorage) Put(s *student.Student) {
	ss[s.Name] = s
}
func (ss StudentStorage) Get(studentName string) (*student.Student, error) {
	s, ok := ss[studentName]
	if !ok {
		return nil, fmt.Errorf("Нет такого студента в списке")
	}
	return s, nil
}
