package Hello

import "fmt"

func Hello(name string) {
	fmt.Println("Hello ", name)
}

func Hello1(name string) string {
	return "Hello " + name
}

func Hello2(name string){
	fmt.Println("Hello2, ", name)
}