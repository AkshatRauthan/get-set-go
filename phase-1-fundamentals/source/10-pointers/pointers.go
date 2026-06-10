package main

func doSomething(x *int){
	*x += 2
	println("Changed value in called function: ", *x)
}

func pointers(){
	// Declaration of Pointers....
	var x *int = new(int)
	*x = 67

	println("Initial Value: ", *x)
	doSomething(x) // OR simply use &x after normal int declration
	println("Value after function call: ", *x)
}
