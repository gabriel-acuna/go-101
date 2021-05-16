package main

import "fmt"

func main() {

	days := make(map[int]string)

	days[0] = "Monday"
	days[1] = "Tuesday"
	days[2] = "Wednesday"
	days[3] = "Thursday"
	days[4] = "Friday"
	days[5] = "Saturday"
	days[6] = "Sunday"

	fmt.Println(days)
	delete(days, 6)
	fmt.Println(days)

	scores := make(map[string][]float32)

	scores["Gabo"] = []float32{9, 7.86, 10}
	scores["Stefano"] = []float32{8.58, 9.36, 10}

	fmt.Println(scores)

	for user, qualifications := range scores {
		fmt.Println(user, qualifications)
	}
}
