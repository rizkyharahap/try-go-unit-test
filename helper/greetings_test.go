package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

/**
* Testing with Fail
* If fail the test still running until end of code
**/
func TestGreetingsFail(t *testing.T) {
	result := Greetings("Rizki")

	if result != "Hello Rizki" {
		// error
		t.Fail()
	}

	fmt.Println("TestGreetingsFail Done")
}

/**
* Testing with FailNow
* If fail the test function is end
* line of code below failNow will not executed
**/
func TestGreetingsFailNow(t *testing.T) {
	result := Greetings("Rizki")

	if result != "Hello Rizki" {
		// error
		t.FailNow()
	}

	fmt.Println("TestGreetingsFailNow Done")
}

/**
* Testing with Error
* If fail will be call t.Fail() with some agrs (optional)
* If fail the test still running until end of code
**/
func TestGreetingsError(t *testing.T) {
	result := Greetings("Rizki")

	if result != "Hello Rizki" {
		// error
		t.Error("Result must be 'Hello Rizki'")
	}

	fmt.Println("TestGreetingsError Done")
}

/**
* Testing with Fatal
* If fail will be call t.FailNow() with some agrs (optional)
* If fail the test function is end
* line of code below failNow will not executed
**/
func TestGreetingsFatal(t *testing.T) {
	result := Greetings("Rizki")

	if result != "Hello Rizki" {
		// error
		t.Fatal("Result must be 'Hello Rizki'")
	}

	fmt.Println("TestGreetingsFatal Done")
}

/**
* Testing with Assert
* If fail will be call t.Fail() with some agrs (optional) and give some detail of errors
* If fail the test still running until end of code
**/
func TestGreetingsAssert(t *testing.T) {
	result := Greetings("Rizki")
	assert.Equal(t, "Hello Rizki", result, "Result must be 'Hello Rizki'")
	fmt.Println("TestGreetingsAssert Done")
}

/**
* Testing with Assert
* If fail will be call t.FailNow() with some agrs (optional) and give some detail of errors
* If fail the test function is end
* line of code below failNow will not executed
**/
func TestGreetingsRequire(t *testing.T) {
	result := Greetings("Rizki")
	require.Equal(t, "Hello Rizki", result, "Result must be 'Hello Rizki'")
	fmt.Println("TestGreetingsRequire Done")
}

/**
* Simulate skip testing
**/
func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Cannot run on Mac OS")
	}

	result := Greetings("Rizki")
	require.Equal(t, "Hello Rizki", result, "Result must be 'Hello Rizki'")
}
