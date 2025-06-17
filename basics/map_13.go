// Program to create a map and print its keys and values

package main

import (
	"fmt"
)

func main() {

	// creating a map
	subjectMarks := map[string]float32{"Golang": 85, "Java": 80, "Python": 81}

	fmt.Println(subjectMarks)

	// create a map
	flowerColor := map[string]string{"Sunflower": "Yellow", "Jasmine": "White", "Hibiscus": "Red"}

	// access value for key Sunflower
	fmt.Println(flowerColor["Sunflower"]) // Yellow

	// access value for key HIbiscus
	fmt.Println(flowerColor["Hibiscus"]) // Red

	// create a map
	capital := map[string]string{"Nepal": "Kathmandu", "US": "New York"}
	fmt.Println("Initial Map: ", capital)

	// change value of US to Washington DC
	capital["US"] = "Washington DC"

	fmt.Println("Updated Map: ", capital)

	// create a map
	students := map[int]string{1: "John", 2: "Lily"}
	fmt.Println("Initial Map: ", students)

	// add element with key 3
	students[3] = "Robin"

	// add element with key 5
	students[5] = "Julie"

	fmt.Println("Updated Map: ", students)

	// create a map
	personAge := map[string]int{"Hermione": 21, "Harry": 20, "John": 25}
	fmt.Println("Initial Map: ", personAge)

	// remove element of map with key John
	delete(personAge, "John")

	fmt.Println("Updated Map: ", personAge)

	// create a map
	squaredNumber := map[int]int{2: 4, 3: 9, 4: 16, 5: 25}

	// for-range loop to iterate through each key-value of map
	for number, squared := range squaredNumber {
		fmt.Printf("Square of %d is %d\n", number, squared)
	}

	// create a map using make()
	student := make(map[int]string)

	// add elements to the map
	student[1] = "Harry"
	student[2] = "Lilly"
	student[5] = "Harmonie"

	fmt.Println(student)

	// create a map
	places := map[string]string{"Nepal": "Kathmandu", "US": "Washington DC", "Norway": "Oslo"}

	// access only the keys of the map
	for country := range places {
		fmt.Println(country)
	}

	for _, capital := range places {
		fmt.Println(capital)
	}
}
