package main

import (
	"fmt"
	"runtime"
)

func main() {
	// Call the testPanic function to run the test
	if err := testPanic(); err != nil {
		fmt.Println("Error:", err)
	}
}

// testPanic simulates a function that encounters a panic to
// test our catchPanic function
func testPanic() (err error) {
	// Schedule the catchPanic function to be called when
	// the testPanic function returns.
	defer catchPanic(&err)

	fmt.Println("Start Test")

	// Mimic a traditional error from a function.
	err = mimicError("1")

	// Trying to dereference a nil pointer will cause the
	// runtime to panic.
	var p *int
	// 此行panic后，其后面的代码将不会被执行。
	// 但该函数仍然会有返回值。如果返回值没有被命名，则函数返回相应类型的默认值；
	// 如果返回值被命名(此处返回值命名为err)，则返回panic发生时返回值(err)；
	// 另外，defer可以在函数返回之前修改该返回值(err)。
	*p = 10

	// Mimic a traditional error from a function.
	// err = mimicError("1")

	fmt.Println("End Test")
	return err
}

// catchPanic catches panics and processes the error.
func catchPanic(err *error) {
	// Check if a panic occurred.
	if r := recover(); r != nil {
		fmt.Println("PANIC Deferred")

		// Capture the stack trace.
		buf := make([]byte, 10000)
		runtime.Stack(buf, false)
		fmt.Println("Stack Trace:", string(buf))

		// If the caller wants the error back provide it.
		if err != nil {
			*err = fmt.Errorf("%v", r)
		}
	}
}

func mimicError(key string) error {
	return fmt.Errorf("Mimic Error: %s", key)
}
