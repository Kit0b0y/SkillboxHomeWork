package main

import (
	"fmt"

	"github.com/Kit0b0y/SkillboxHomeWork/NewSkillbox/28/pkg/storage"
	"github.com/Kit0b0y/SkillboxHomeWork/NewSkillbox/28/pkg/student"
)

func main() {
	a := student.NewStudent("Вася", 24, 1)
	b := student.NewStudent("Семен", 32, 2)
	ss := storage.NewStudentStorage()
	ss.Put(a)
	ss.Put(b)
	fmt.Println(ss.Get(a.Name))
	fmt.Println(ss.Get(b.Name))
}
