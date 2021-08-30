package main

import (
	"errors"
	"fmt"
	"syscall"
)

func main() {
	// 7.8 The error interface
	// crear error
	er := errors.New("new error type")
	fmt.Printf("type: %T, value: %[1]v\n", er)
	// output: type: *errors.errorString, value: new error type

	// errors.New(): allocates a distinct error instance
	err1 := errors.New("EOF")
	err2 := errors.New("EOF")
	fmt.Println(err1 == err2) // flase

	// wrapper function (fmt.Errorf)
	er = fmt.Errorf("new error from Errorf")
	fmt.Println(er)

	// error in syscall package
	var err error = syscall.Errno(2)
	fmt.Println(err.Error())                    // "no such file or directory"
	fmt.Println(err)                            // "no such file or directory"
	fmt.Printf("type: %T, value: %[1]v\n", err) // "type: syscall.Errno, value: no such file or directory"
}
