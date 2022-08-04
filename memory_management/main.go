package main

import "fmt"

func main() {

	//1. Unassigned pointer
	var unAssignedPointer *int
	fmt.Println("Value of unAssignedPointer:", unAssignedPointer)       //<nil>
	//fmt.Println("Invalid pointer:", *unAssignedPointer) //Crashes the application

	//2. Assigned pointer
	someInteger := 67
	var assignedPointer = &someInteger
	fmt.Println("Value of variable being pointed:", *assignedPointer) //67
	fmt.Println("Value of memory address of variable being pointed:", assignedPointer) // The memory address of someInteger (0xc0000b2008)

	//3. Changing a variable from a pointer
	value1 := 4
	pointer1 := &value1
	fmt.Println("Value1 from pointer:", *pointer1)

	*pointer1 = *pointer1/2 //
	fmt.Println("Pointer 1:", *pointer1) //2
	fmt.Println("Value 1:", value1) //2

}
