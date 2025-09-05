package main

import "fmt"

func main() {
	arr := [50]int{}

	foundPrime := 0
	num := 0

	for foundPrime < 50 {
		if isPrime(num) {
			arr[foundPrime] = num
			foundPrime++
		}
		num++
	}
	fmt.Println(arr)

	fmt.Println("result : ")
	for i := 0; i < 50; i += 2 {
		fmt.Print(arr[i], "  ")
	}
}

func isPrime(num int) bool {
	if num == 0 || num == 1 {
		return false
	}
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
