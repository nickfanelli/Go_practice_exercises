// PinPan Game
// This program prints the numbers from 1 to 100, but for multiples of 3 it prints "Pin" instead of the number
// and for multiples of 5 it prints "Pan". For numbers which are multiples of both 3 and 5 it prints "PinPan".

package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("PinPan")
		} else if i%3 == 0 {
			fmt.Println("Pin")
		} else if i%5 == 0 {
			fmt.Println("Pan")
		} else {
			fmt.Println(i)
		}
	}
}
