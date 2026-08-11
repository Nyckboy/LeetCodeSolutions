
// Time Complexity: O(1)
// Space Complexity: O(n)

package main

import (
	"fmt"
)


type Item struct {
	val int
	min int
}
type MinStack struct {
	stack []Item
}

func Constructor() MinStack {
	return MinStack{stack: make([]Item, 0)}
}

func (this *MinStack) Push(val int) {
	currMin := val
	if len(this.stack) > 0 {
		topElement := this.stack[len(this.stack) -1]
		currMin = min(currMin, topElement.min)
	}
	newVal := Item{val, currMin}
	this.stack = append(this.stack, newVal)
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack) - 1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack) - 1].val
}

func (this *MinStack) GetMin() int {
	return this.stack[len(this.stack) - 1].min
}

func main() {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	
	// Expected: -3
	fmt.Println("GetMin() ->", minStack.GetMin()) 
	
	minStack.Pop()
	
	// Expected: 0
	fmt.Println("Top() ->", minStack.Top())    
	
	// Expected: -2
	fmt.Println("GetMin() ->", minStack.GetMin()) 
}