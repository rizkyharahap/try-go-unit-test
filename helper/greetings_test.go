package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreetingsFail(t *testing.T) {
	result := Greetings("Rizki")

	if result != "Hello Rizki" {
		// error
		t.Fail()
	}

	fmt.Println("TestGreetingsFail Done")
}

func TestGreetingsFailNow(t *testing.T) {
	result := Greetings("Rizki")

	if result != "Hello Rizki" {
		// error
		t.FailNow()
	}

	fmt.Println("TestGreetingsFailNow Done")
}

func TestGreetingsError(t *testing.T) {
	result := Greetings("Rizki")

	if result != "Hello Rizki" {
		// error
		t.Error("Result must be 'Hello Rizki'")
	}

	fmt.Println("TestGreetingsError Done")
}

func TestGreetingsFatal(t *testing.T) {
	result := Greetings("Rizki")

	if result != "Hello Rizki" {
		// error
		t.Fatal("Result must be 'Hello Rizki'")
	}

	fmt.Println("TestGreetingsFatal Done")
}

func TestGreetingsAssert(t *testing.T) {
	result := Greetings("Rizki")
	assert.Equal(t, "Hello Rizki", result, "Result must be 'Hello Rizki'")
	fmt.Println("TestGreetingsAssert Done")
}
