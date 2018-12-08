package data

import "fmt"

func main() {
	var fun = func () *int {
		// int x = 0 <- invalid expression
		var x = 0
		fmt.Println(x)
		fmt.Println(&x)
		x = 1
		return &x
	}
	ret := fun()
	fmt.Println(ret)
	fmt.Println(*ret)
	fmt.Println(&ret)
	ret = call(ret)
	fmt.Println("after call function")
	fmt.Println(ret)
	fmt.Println(*ret)
	fmt.Println(&ret)

	ret = doublePointerCall(&ret)
	fmt.Println(ret)


}

func doublePointerCall(par **int) *int {
	fmt.Println(par)
	fmt.Println(*par)
	return *par
}


func call(par *int) *int {
	fmt.Println("call function")
	fmt.Println(par)
	fmt.Println(&par)
	var x = par
	fmt.Println(x)
	fmt.Println(&x)
	fmt.Println(*x)
	*x = *x + 1
	return x
}