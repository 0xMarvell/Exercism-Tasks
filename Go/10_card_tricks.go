/* Instructions
As a magician-to-be, Elyse needs to practice some basics. She has a stack of cards that she wants to manipulate.

To make things a bit easier she only uses the cards 1 to 10.

1. Retrieve a card from a stack
Return the card at position index from the given stack.

card, ok := GetItem([]int{1, 2, 4, 1}, 2)
fmt.Println(card)
// Output: 4
fmt.Println(ok)
// Output: true
If the index is out of bounds (ie. if it is negative or after the end of the stack), we want to return 0 and false:

card, ok := GetItem([]int{1, 2, 4, 1}, 10)
fmt.Println(card)
// Output: 0
fmt.Println(ok)
// Output: false
2. Exchange a card in the stack
Exchange the card at position index with the new card provided and return the adjusted stack. Note that this will also change the input slice which is ok.

index := 2
newCard := 6
SetItem([]int{1, 2, 4, 1}, index, newCard)
// Output: []int{1, 2, 6, 1}
If the index is out of bounds (ie. if it is negative or after the end of the stack), we want to append the new card to the end of the stack:

index := -1
newCard := 6
SetItem([]int{1, 2, 4, 1}, index, newCard)
// Output: []int{1, 2, 4, 1, 6}
3. Create a stack of cards
Create a stack of given length and fill it with cards of the given value. If the given length is negative or zero, return an empty stack.

PrefilledSlice(8, 3)
// Output: []int{8, 8, 8}
4. Remove a card from the stack
Remove the card at position index from the stack and return the stack.

RemoveItem([]int{3, 2, 6, 4, 8}, 2)
// Output: []int{3, 2, 4, 8}
If the index is out of bounds (ie. if it is negative or after the end of the stack), we want to leave the stack unchanged:

RemoveItem([]int{3, 2, 6, 4, 8}, 11)
// Output: []int{3, 2, 6, 4, 8} */

// package cards
package main

import "fmt"

// GetItem retrieves an item from a slice at given position. The second return value indicates whether
// the given index exists in the slice or not.
func GetItem(slice []int, index int) (int, bool) {
	//panic("Please implement the GetItem function")
	if len(slice) > index && index >= 0 {
		return slice[index], index < len(slice)-1
	}
	return 0, false
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	//panic("Please implement the SetItem function")
	if len(slice) > index && index >= 0 {
		slice[index] = value
	} else {
		slice = append(slice, value)
	}
	return slice
}

// PrefilledSlice creates a slice of given length and prefills it with the given value.
func PrefilledSlice(value, length int) []int {
	panic("Please implement the PrefilledSlice function")
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	panic("Please implement the RemoveItem function")
}

func main() {
	card, ok := GetItem([]int{1, 2, 4, 1}, 2)
	fmt.Println(card, ok)
	card, ok = GetItem([]int{1, 2, 4, 1}, 10)
	fmt.Println(card, ok)
	index := 2
	newCard := 6
	fmt.Println(SetItem([]int{1, 2, 4, 1}, index, newCard))
	index = -1
	newCard = 6
	fmt.Println(SetItem([]int{1, 2, 4, 1}, index, newCard))
}
