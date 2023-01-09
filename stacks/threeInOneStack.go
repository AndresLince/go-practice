/**
 * 3.1
 * Three in One: Describe how you could use a single array to implement
 * three stacks.
 */
package stacks

import (
	"errors"
	"math"
)

type ThreeInOneStack struct {
	array                []interface{}
	length               int
	arrayLastPosition    []int
	arrayStartPosition   []int
	arrayCurrentPosition []int
}

func NewThreeInOneStack(length int) *ThreeInOneStack {
	array := make([]interface{}, length)
	firstLength := int(math.Floor(float64(length / 3)))
	arrayLastPosition := []int{
		firstLength - 1,
		firstLength*2 - 1,
		length - 1,
	}
	arrayStartPosition := []int{
		0,
		firstLength,
		firstLength * 2,
	}
	arrayCurrentPosition := []int{
		0,
		firstLength,
		firstLength * 2,
	}
	tios := &ThreeInOneStack{
		array:                array,
		length:               length,
		arrayLastPosition:    arrayLastPosition,
		arrayStartPosition:   arrayStartPosition,
		arrayCurrentPosition: arrayCurrentPosition,
	}
	return tios
}

func (threeInOneStack *ThreeInOneStack) Push(stack int, data interface{}) error {
	if threeInOneStack.arrayCurrentPosition[stack] > threeInOneStack.arrayLastPosition[stack] {
		return errors.New("full_stack")
	}
	currentIndex := threeInOneStack.arrayCurrentPosition[stack]
	threeInOneStack.array = append(threeInOneStack.array[:currentIndex], threeInOneStack.array[currentIndex:]...) // currentIndex < len(a)
	threeInOneStack.array[currentIndex] = data
	threeInOneStack.arrayCurrentPosition[stack]++
	return nil
}

func (threeInOneStack *ThreeInOneStack) Pop(stack int) (interface{}, error) {
	if threeInOneStack.arrayCurrentPosition[stack] == threeInOneStack.arrayStartPosition[stack] {
		return nil, errors.New("empty_stack")
	}
	threeInOneStack.arrayCurrentPosition[stack]--
	return threeInOneStack.array[threeInOneStack.arrayCurrentPosition[stack]], nil
}

func (threeInOneStack ThreeInOneStack) IsEmpty(stack int) bool {
	return threeInOneStack.arrayCurrentPosition[stack] == threeInOneStack.arrayStartPosition[stack]
}

func (threeInOneStack ThreeInOneStack) Peak(stack int) interface{} {
	currentIndex := threeInOneStack.arrayCurrentPosition[stack] - 1
	return threeInOneStack.array[currentIndex]
}
