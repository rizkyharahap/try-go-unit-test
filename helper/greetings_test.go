package helper

import (
	"fmt"
	"testing"
)

func TestGreetingsFail(t *testing.T) {
	result := Greetings("Rizki")

	if result != "Hello Rizki" {
		// error
		t.Fail()
	}

	fmt.Println("TestGreetings Done")
}

func TestGreetingsFailNow(t *testing.T) {
	result := Greetings("Rizki")

	if result != "Hello Rizki" {
		// error
		t.FailNow()
	}

	fmt.Println("TestGreetings Done")
}

func TestGreetingsError(t *testing.T) {
	result := Greetings("Rizki")

	if result != "Hello Rizki" {
		// error
		t.Error("Result must be 'Hello Rizki'")
	}

	fmt.Println("TestGreetings Done")
}

func TestGreetingsFatal(t *testing.T) {
	result := Greetings("Rizki")

	if result != "Hello Rizki" {
		// error
		t.Fatal("Result must be 'Hello Rizki'")
	}

	fmt.Println("TestGreetings Done")
}
