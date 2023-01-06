package main

type ThreeInOneStack struct {
	array                []interface{}
	length               int
	arrayLastPosition    []int
	arraySartPosition    []int
	arrayCurrentPosition []int
}

func (threeInOneStack ThreeInOneStack) NewThreeInOneStack(length int) *ThreeInOneStack {
	array := []interface{}{}
	arrayLastPosition := []int{1, 2, 3}
	arraySartPosition := []int{1, 2, 3}
	arrayCurrentPosition := []int{1, 2, 3}
	tios := &ThreeInOneStack{
		array:                array,
		length:               length,
		arrayLastPosition:    arrayLastPosition,
		arraySartPosition:    arraySartPosition,
		arrayCurrentPosition: arrayCurrentPosition,
	}
	return tios
}
