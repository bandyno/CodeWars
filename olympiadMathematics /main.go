package main

import (
	"fmt"
	"sort"
)

type student struct {
	name   string
	school *school
	points int
}

type school struct {
	name          string
	points        int
	countStudents int
}

func main() {
	schools := []school{
		{name: "Школа №48", points: 0, countStudents: 0},
		{name: "Школа №88", points: 0, countStudents: 0},
		{name: "Школа №89", points: 0, countStudents: 0},
	}

	students := []student{
		{"Александр Котлеткин", &schools[0], 32},
		{"Тимофей Белкин", &schools[0], 100},
		{"Нурхултан Динамщиков", &schools[0], 10},
		{"Сергей Сладенький", &schools[1], 76},
		{"Валентин Блогеров", &schools[2], 56},
	}

	sort.Slice(students, func(i, j int) bool {
		return students[i].points > students[j].points
	})

	fmt.Println("Таблица победителей:")

	for i := 0; i < 3; i++ {
		fmt.Printf("Место №%d: %s - %d бал., %v\n", i+1, students[i].name, students[i].points, students[i].school.name)
	}

	for _, v := range students {
		v.school.points += v.points
		v.school.countStudents++
	}

	sort.Slice(schools, func(i, j int) bool {
		return schools[i].points/schools[i].countStudents > schools[j].points/schools[j].countStudents
	})

	for _, v := range schools {
		fmt.Printf("%s - средний балл: %v\n", v.name, v.points/v.countStudents)
	}
}
