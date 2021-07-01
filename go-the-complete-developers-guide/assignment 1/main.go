package main

func main() {

	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for a := range arr {
		if a%2 == 0 {
			println(a, "is even")
		} else {
			println(a, "is odd")
		}
	}
}
