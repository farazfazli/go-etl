package main

import (
	"fmt"
	"testing"
)

var addTable = []struct {
	in  []int
	out int
}{
	{[]int{1, 2}, 3},
	{[]int{1, 2, 4, 5}, 12},
}

func TestAdd(T *testing.T) {
	for _, entry := range addTable {
		result := Add(entry.in...)
		if result != entry.out {
			T.Log(entry.in)
			T.Log(result)
			T.Log(entry.out)
			T.Fatal("ERROR!")
		} else {
			T.Log("SUCCESS!")
		}
	}
}

func BenchmarkMain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main()
	}
}

func ExampleMain() {
	fmt.Println("Hello")

	// Output:
	// Hello
}
