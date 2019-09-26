package main

//An interface variable has a dynamic type and a dynamic value.
//It is nil only if both the type and value are nil.

type S struct{}

func (s S) F() {}

type IF interface {
	F()
}

func InitType() S {
	var s S
	return s
}

func InitPointer() *S {
	var s *S
	return s
}

func InitEFace() interface{} {
	var s interface{}
	return s
}

func InitEfaceType() interface{} {
	var s S
	return s
}

func InitEfacePointer() interface{} {
	var s *S
	return s
}

func InitIfaceType() IF {
	var s S
	return s
}

func InitIfacePointer() IF {
	var s *S
	return s
}

func main() {
	//println(InitType() == nil)
	println(InitPointer() == nil)
	println(InitEFace() == nil)
	println(InitEfaceType() == nil)
	println(InitEfacePointer() == nil)
	println(InitIfaceType() == nil)
	println(InitIfacePointer() == nil)
}
