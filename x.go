package testgomod

import "fmt"

// Hi prints greetings for some name
func Hi(name string) {
	fmt.Printf("Hey, %s!!\n", name)
}

// Bye prints bye message
func Bye(name string) {
	fmt.Printf("Bye, %s :(\n", name)
}
