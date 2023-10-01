package main

import (
	"fmt"

	"github.com/bashery/kv"
)

func main() {

	intmap := kv.New[int, int]()
	intmap.Set(1, 123)

	fmt.Println(intmap.Get(1))
	test, ok := intmap.Get(234)
	if !ok {
		fmt.Println(234, "not found")
	}

	fmt.Println(test)

	intval, _ := intmap.Get(1)
	if intval != 123 {
		fmt.Printf("key must be %d", 123)
	}

	if ok := intmap.HasKey(1); ok != true {
		fmt.Printf("'1' key must be exist")
	}

	//
	strmap := kv.New[string, string]()
	strmap.Set("hi", "hello")
	sval, _ := strmap.Get("hi")
	if sval != "hello" {
		fmt.Printf("value must be %s", "hello")
	}

	if ok := strmap.HasKey("hi"); ok != true {
		fmt.Printf("'hi' key must be exist")
	}

	fmt.Println("All Functions pass")
	fmt.Println("=============== concurrent mod ==============")
}
