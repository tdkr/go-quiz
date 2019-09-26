package main

type S struct {
}

func f(x interface{}) {
}

func g(x *interface{}) {
}

func main() {
	s := S{}
	p := &s
	f(s) //A correct
	g(s) //B incorrect
	f(p) //C correct
	g(p) //D incorrect

	var t interface{}
	g(&t)
}