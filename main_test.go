package main

import "testing"

func TestHelloWorld(t *testing.T) {
	expected := "Hello, World!"
	if HelloWorld() != expected {
		t.Errorf("Expected %s but got %s", expected, HelloWorld())
	}
}