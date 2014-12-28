package main

import (
	"fmt"
)

func main() {
	var a [5]int
	a[4] = 100
	fmt.Println(a)

	x := [5]float64{98, 93, 77, 82, 83}

	var total float64 = 0
	for _, value := range x {
		total += value
	}
	fmt.Println(total / float64(len(x)))

	var s = make([]float64, 5, 10)
	s[0] = 12
	fmt.Println(s)

	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)

	slice3 := []int{1, 2, 3}
	slice4 := make([]int, 2)
	copy(slice4, slice3)
	fmt.Println(slice3, slice4)

	m1 := make(map[string]int)
	m1["key"] = 10
	fmt.Println(m1["key"])

	m2 := make(map[int]int)
	m2[1] = 10
	fmt.Println(m2[1])

	elements := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
		"N":  "Nitrogen",
		"O":  "Oxygen",
		"F":  "Fluorine",
		"Ne": "Neon",
	}
	fmt.Println(elements["Li"])
	fmt.Println(elements)

	delete(elements, "H")
	fmt.Println(elements)

	if name, ok := elements["Un"]; ok {
		fmt.Println(name, ok)
	}
	if name, ok := elements["Be"]; ok {
		fmt.Println(name, ok)
	}

	elements2 := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
		"Be": map[string]string{
			"name":  "Beryllium",
			"state": "solid",
		},
		"B": map[string]string{
			"name":  "Boron",
			"state": "solid",
		},
		"C": map[string]string{
			"name":  "Carbon",
			"state": "solid",
		},
		"N": map[string]string{
			"name":  "Nitrogen",
			"state": "gas",
		},
		"O": map[string]string{
			"name":  "Oxygen",
			"state": "gas",
		},
		"F": map[string]string{
			"name":  "Fluorine",
			"state": "gas",
		},
		"Ne": map[string]string{
			"name":  "Neon",
			"state": "gas",
		},
	}

	if el, ok := elements2["He"]; ok {
		fmt.Println(el["name"], el["state"])
	}
}
