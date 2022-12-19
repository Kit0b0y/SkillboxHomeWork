package main

import (
	"fmt"

	"github.com/kit0b0y/skillboxhomeWork/newskillbox/28/pkg/storage"
	"github.com/kit0b0y/skillboxhomework/newskillbox/28/pkg/student"
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
