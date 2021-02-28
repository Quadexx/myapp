package main

import "fmt"

func main() {
	var s1 []int
	fmt.Println("Empty Slice", s1)
	s1 = append(s1, 200)
	fmt.Println("Now full slice", s1)

	fmt.Println("Length slice", len(s1))
	fmt.Println("Lenght inside array", s1, cap(s1))
	s1 = append(s1, 220)
	fmt.Println("Lenght inside array is upper", s1, cap(s1))
	s1 = append(s1, 230)
	s1 = append(s1, 240)
	fmt.Println("Lenght inside array is upper", s1, cap(s1))
	s1 = append(s1, 250)
	fmt.Println("Lenght inside array is upper", s1, cap(s1))

	s2 := []int{
		30,
		40,
		50,
		70,
		80,
	}
	fmt.Println(s2, cap(s2), len(s2))

	slice3 := make([]int, 10)
	slice3 = append(slice3, 1)
	fmt.Println(slice3, len(slice3), cap(slice3))

	return
}
