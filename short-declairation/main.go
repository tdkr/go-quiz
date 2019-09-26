package main

import "fmt"

func f() (interface{}, interface{}) {
	return nil, nil
}

func main()  {
	var x interface{} = nil
	//x, _ := f()
	x, _ = f()
	x, y := f()
	x, y = f()
	fmt.Println(x, y)
}
