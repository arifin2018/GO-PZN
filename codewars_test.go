package gopzn

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)


func LoveFuncMySelf(flower1, flower2 int) bool {
	if (flower1 % 2 == 0 && flower2 % 2 == 1) || (flower2 % 2 == 0 && flower1 % 2 == 1) {
		return true
	}else{
		return false
	}
}

func LoveFunc(flower1, flower2 int) bool {
	fmt.Println("flower1", flower1)
	fmt.Println("flower2", flower2)
	fmt.Println("flower1+flower2", flower1+flower2)
	fmt.Println("(flower1+flower2) % 2 == 1 => ", (flower1+flower2) % 2 == 1)
	return !((flower1+flower2) % 2 == 0)
}

func TestLoveFunc(t *testing.T)  {
	assert.Equal(t, true, LoveFunc(1,4))
	assert.Equal(t, false, LoveFunc(2,2))
	assert.Equal(t, true, LoveFunc(0,1))
	assert.Equal(t, false, LoveFunc(0,0))
}

type Node struct {
    data int
    next *Node
}

type List struct {
    head *Node
}

func MoveZeros(arr []int) []int {
	var zero int
	var data []int
	list := &List{}
	for _, v := range arr {
		if v == 0 {
			zero++
		}else{
			list.add(v)
			data = append(data, v)
		}
	}
	for i := 0; i < zero; i++ {
		data = append(data, 0)
	}

	return data
}

func (l *List) add(value int) {
	newNode := &Node{data: value}
	if l.head == nil {
        l.head = newNode
        return
    }

    curr := l.head
    for curr.next != nil {
        curr = curr.next
    }

    curr.next = newNode
}

func TestMovingZerosToTheEnd(t *testing.T)  {
	fmt.Println(MoveZeros([]int{1,2,0,1,0,1,0,3,0,1}))
	// assert.Equal(t, false, MoveZeros([]int{ 1, 2, 1, 1, 3, 1, 0, 0, 0, 0 }))
}