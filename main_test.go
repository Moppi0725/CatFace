package main

import "testing"

func Example_Main() { g
	oMain([]string{})
	// Output:
  	// Hello World
}
func Test_Main(t *testing.T) {
	if status := goMain([]string{}); status != 0 { 
		t.Error("Expected 0, got ", status)
	}
}