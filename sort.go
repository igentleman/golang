package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
)

type student struct{
	name string
	age int
	vj int // 学生成绩
}

type studentSlice []student

func (s studentSlice) Len()int{
	return len(s)
}

func (s studentSlice) Less(i int,j int)bool{
	return s[i].age < s[j].age
}

func (s studentSlice) Swap(i int,j int){
	s[i],s[j] = s[j],s[i]
}

func main(){
	var students studentSlice
	//随机创建10个学生
	for i := 0;i < 10;i++{
		var student student
		student.name = "刘粤新" + strconv.Itoa(rand.Intn(100))
		student.age = randInt(9,20)
		student.vj = randInt(60,150)
		students = append(students,student)
	}
	sort.Sort(students)
	fmt.Println(students)
}

func randInt(min int , max int) int {
	return min + rand.Intn(max-min)
}