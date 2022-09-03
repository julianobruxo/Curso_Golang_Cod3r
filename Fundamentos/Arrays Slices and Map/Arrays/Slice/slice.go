package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3}
	s1 := []int{1, 2, 3}
	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(a2)

	s2 := a2[0:4]
	fmt.Println(s2)

	s3 := s2[1]
	fmt.Println(s3)
}
