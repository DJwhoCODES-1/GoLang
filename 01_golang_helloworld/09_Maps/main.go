package main

import "fmt"

func main() {
	fmt.Println("Maps : A map is a data structure that provides an unordered collection of key-value pairs, where each is unique.")
	fmt.Println("If a key does not exists and we try to get its data, it will give us 0")

	studentGrades := make(map[string]int)

	studentGrades["Devanshu"] = 95
	studentGrades["Arush"] = 49
	studentGrades["Nadeem"] = 78

	fmt.Println("Marks of Devanshu:", studentGrades["Devanshu"])

	studentGrades["Arush"] = 94
	arushMarks, arushDoeExists := studentGrades["Arush"]
	fmt.Println(arushMarks, arushDoeExists)
	fmt.Println("Marks of Arush:", studentGrades["Arush"])

	delete(studentGrades, "Arush")
	fmt.Println("Marks of Arush:", studentGrades["Arush"])

	marks, doeExists := studentGrades["Arush"]
	fmt.Println(marks, doeExists)

	for index, value := range studentGrades {
		fmt.Printf("Key is %s and marks are %d\n", index, value)
	}

	person := map[string]int{
		"alice":   64,
		"david":   87,
		"charlie": 34,
	}

	for index, value := range person {
		fmt.Printf("Key is %s and marks are %d\n", index, value)
	}
}
