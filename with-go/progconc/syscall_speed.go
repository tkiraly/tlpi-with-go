package main

import (
	"fmt"
	"os"
	"strconv"
	"syscall"
)

func myFunc() int { return 1 }

/* syscall_speed.go

   By repeatedly invoking a simple system call (getppid()), we can get some
   idea of the cost of making system calls.

   Usage: time syscall_speed numcalls
                           Def=10000000

   Exporting environment variable NOSYSCALL causes a call to a simple function
   returning an integer, which can be used to compare the overhead
   of a simple function call against that of a system call.
*/

func main() {
	_, NOSYSCALL := os.LookupEnv("NOSYSCALL")
	numCalls := 10000000
	if len(os.Args) > 1 {
		var err error
		numCalls, err = strconv.Atoi(os.Args[1])
		if err != nil {
			panic(err)
		}
	}
	// Environment variable simulates the compile time define
	// and for loop was cut into two pieces to mimic more precisely
	// the original C code
	if NOSYSCALL {
		fmt.Printf("Calling normal function()\n")
		for i := 0; i < numCalls; i++ {
			myFunc()
		}
	} else {
		fmt.Printf("Calling getppid()\n")
		for i := 0; i < numCalls; i++ {
			syscall.Getpid()
		}
	}
}
