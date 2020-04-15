package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

var primes = map[int]bool{}

func isPrimeSqrt(value int) bool {

	if value == 1 {
		return false
	} else if value == 2 || value == 3 {
		return true
	} else if value%2 == 0 || value%3 == 0 {
		return false
	}

	if isPrime, exists := primes[value]; exists {
		return isPrime
	}

	for i := 2; i <= int(math.Floor(math.Sqrt(float64(value)))); i++ {
		if value%i == 0 {
			primes[value] = false
			return false
		}
	}

	isPrime := value > 1
	primes[value] = isPrime
	return isPrime
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for scanner.Scan() {
		num, _ := strconv.ParseInt(scanner.Text(), 10, 32)
		if isPrimeSqrt(int(num)) {
			fmt.Fprintln(writer, "1")
		} else {
			fmt.Fprintln(writer, "0")
		}
	}
}
