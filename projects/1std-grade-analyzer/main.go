package main

import (
	"fmt"
)

// -------- Helper Functions --------

// Calculate average score
func average(nums []int) float64 {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return float64(sum) / float64(len(nums))
}

// Find min and max
func minMax(nums []int) (int, int) {
	min, max := nums[0], nums[0]
	for _, v := range nums {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}

// Safe division with error handling
func safeDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero is not allowed")
	}
	return a / b, nil
}

// -------- Main Program --------

func main() {
	var n int
	fmt.Print("Enter number of students: ")
	fmt.Scanln(&n)

	// store data
	students := make([]string, 0, n)
	marks := make(map[string][]int)

	for i := 0; i < n; i++ {
		var name string
		fmt.Printf("\nEnter name of student %d: ", i+1)
		fmt.Scanln(&name)

		students = append(students, name)

		var m int
		fmt.Printf("Enter number of subjects for %s: ", name)
		fmt.Scanln(&m)

		scores := make([]int, m)
		fmt.Printf("Enter %d scores separated by space: ", m)
		for j := 0; j < m; j++ {
			fmt.Scan(&scores[j])
		}
		marks[name] = scores
	}

	// process each student
	fmt.Println("\n--- Results ---")
	for _, name := range students {
		scores := marks[name]
		avg := average(scores)
		min, max := minMax(scores)

		fmt.Printf("\nStudent: %s\n", name)
		fmt.Printf("Scores: %v | Average: %.2f | Min: %d | Max: %d\n", scores, avg, min, max)

		switch {
		case avg >= 90:
			fmt.Println("Grade: A")
		case avg >= 75:
			fmt.Println("Grade: B")
		case avg >= 60:
			fmt.Println("Grade: C")
		default:
			fmt.Println("Grade: D")
		}
	}

	// demonstrate safe division
	fmt.Println("\n--- Safe Division Demo ---")
	result, err := safeDivide(100, 5)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("100 / 5 =", result)
	}
}

/*
Sample output:

Enter number of students: 2

Enter name of student 1: Alice
Enter number of subjects for Alice: 3
Enter 3 scores separated by space: 85 90 78

Enter name of student 2: Bob
Enter number of subjects for Bob: 3
Enter 3 scores separated by space: 70 65 80

--- Results ---

Student: Alice
Scores: [85 90 78] | Average: 84.33 | Min: 78 | Max: 90
Grade: B

Student: Bob
Scores: [70 65 80] | Average: 71.67 | Min: 65 | Max: 80
Grade: C

--- Safe Division Demo ---
100 / 5 = 20

*/
