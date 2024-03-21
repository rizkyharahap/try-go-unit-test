package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

/**
* Before and After Unit Test
* Can be add some action before or after run test with testing.M
* This will be run per module test not each of test
**/
func TestMain(m *testing.M) {
	// before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	// after
	fmt.Println("AFTER UNIT TEST")
}

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

/**
* Sub Test can be write test function inside test function
**/
func TestSubTest(t *testing.T) {
	t.Run("Rizki", func(t *testing.T) {
		result := Greetings("Rizki")
		require.Equal(t, "Hello Rizki", result, "Result must be 'Hello Rizki'")
	})
	t.Run("Harahap", func(t *testing.T) {
		result := Greetings("Harahap")
		require.Equal(t, "Hello Harahap", result, "Result must be 'Hello Harahap'")
	})
}

/**
* Test table used for define data and test each of struct
**/
func TestTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{name: "Greetings(Rizki)", request: "Rizki", expected: "Hello Rizki"},
		{name: "Greetings(Mahfuddin)", request: "Mahfuddin", expected: "Hello Mahfuddin"},
		{name: "Greetings(Harahap)", request: "Harahap", expected: "Hello Harahap"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Greetings(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

/**
* Benchmark for check code performance
**/
func BenchmarkGreetings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greetings("Rizki")
	}
}
