package utils

import "fmt"

type Stack []interface{}

func (stack *Stack) Push(value interface{}) {
	(*stack) = append((*stack), value)
}

func (stack *Stack) Size() int {
	return len(*stack)
}

func (stack *Stack) IsEmpty() bool {
	return stack.Size() == 0
}

func (stack *Stack) Pop() interface{} {
	if !stack.IsEmpty() {
		value := (*stack)[stack.Size()-1]
		(*stack) = (*stack)[:stack.Size()-1]
		return value
	}
	return nil
}

func (stack *Stack) Peek() interface{} {
	if !stack.IsEmpty() {
		return (*stack)[stack.Size()-1]
	}
	return nil
}

type ArrayList []interface{}

func (list *ArrayList) Add(a interface{}, x ...int) {
	if len(x) == 0 {
		(*list) = append((*list), a)
	} else {
		index := x[0]
		(*list) = append(append((*list)[:index], a), (*list)[index:]...)
	}
}

func (list *ArrayList) Remove(index int) {
	(*list) = append((*list)[:index], (*list)[index+1:]...)
}

func (list *ArrayList) Size() int {
	return len(*list)
}

func (list *ArrayList) Get(index int) interface{} {
	return (*list)[index]
}

func (list *ArrayList) Set(index int, a interface{}) {
	(*list)[index] = a
}

func (list *ArrayList) Contains(a interface{}) bool {
	for _, value := range *list {
		if a == value {
			return true
		}
	}
	return false
}

func (list *ArrayList) Sort() {
	Array := true
	for i := 0; i < list.Size(); i++ {
		value := list.Get(i)
		switch Type := value.(type) {
		case int:
			Array = true
		default:
			_ = Type
			Array = false
		}
		if !Array {
			break
		}
	}
	nums := make([]int, list.Size())
	for index := 0; index < list.Size(); index++ {
		nums[index] = list.Get(index).(int)
	}
	sort(nums)
	for i := 0; i < len(nums); i++ {
		list.Set(i, nums[i])
	}
}

func Merge(A []int, B []int) []int {
	var result = make([]int, len(A)+len(B))
	i, j := 0, 0
	index := 0
	for i < len(A) && j < len(B) {
		if A[i] < B[j] {
			result[index] = A[i]
			i++
		} else if B[j] < A[i] {
			result[index] = B[j]
			j++
		} else {
			result[index] = A[i]
			i++
			index++
			result[index] = B[j]
			j++
		}
		index++
	}
	for i < len(A) {
		result[index] = A[i]
		index++
		i++
	}
	for j < len(B) {
		result[index] = B[j]
		index++
		j++
	}
	return result
}

func sort(nums []int) {
	if len(nums) < 2 {
		return
	}
	mid := len(nums) / 2
	A := nums[:mid]
	B := nums[mid:]
	sort(A)
	sort(B)
	copy(nums, Merge(A, B))
}

type Node struct {
	val  interface{}
	next *Node
}

type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func (list *LinkedList) Add(x interface{}, useless ...int) {
	node := new(Node)
	node.val = x
	if list.length == 0 {
		list.head = node
		list.tail = node
	} else {
		list.tail.next = node
		list.tail = list.tail.next
	}
	list.length++
}

func (list *LinkedList) Remove(index int) {

	if index < 0 || index >= list.length {
		return
	}
	if index == 0 {
		list.head = list.head.next
		return
	}
	pointer := list.head
	for i := 1; i < index; i++ {
		pointer = pointer.next
	}
	if index == list.length-1 {
		list.tail = pointer
		list.tail.next = nil
	} else {
		pointer.next = pointer.next.next
	}
	list.length--
}

func (list *LinkedList) Size() int {
	return list.length
}

func (list *LinkedList) Print() {
	pointer := list.head
	fmt.Print("[")
	for pointer != nil {
		fmt.Print(pointer.val)
		if pointer != list.tail {
			fmt.Print("->")
		}
		pointer = pointer.next
	}
	fmt.Println("]")
}

func (list *LinkedList) Filter(function func(interface{}) bool) LinkedList {
	pointer := list.head
	var linkedList LinkedList
	for pointer != nil {
		if function(pointer.val.(int)) {
			linkedList.Add(pointer.val)
		}
		pointer = pointer.next
	}
	return linkedList
}

func (list *LinkedList) Get(index int) interface{} {
	pointer := list.head
	for i := 1; i < index; i++ {
		pointer = pointer.next
	}
	return pointer.val
}

func (list *LinkedList) Set(index int, value interface{}) {
	pointer := list.head
	for i := 1; i < index; i++ {
		pointer = pointer.next
	}
	pointer.val = value
}

type List interface {
	Add(interface{}, ...int)
	Remove(int)
	Size() int
	Get(int) interface{}
	Set(int, interface{})
}

type Queue struct {
	Contents LinkedList
}

func (q *Queue) Add(a interface{}) {
	q.Contents.Add(a)
}

func (q *Queue) Remove() interface{} {
	value := (*q).Contents.Get(0)
	q.Contents.Remove(0)
	return value
}

func (q *Queue) Size() int {
	return q.Contents.Size()
}

type Set map[interface{}]bool

func (set *Set) Add(a interface{}) {
	if len(*set) == 0 {
		(*set) = make(map[interface{}]bool)
	}
	(*set)[a] = true
}

func (set *Set) Remove(a interface{}) {
	if len(*set) == 0 {
		return
	}
	delete((*set), a)
}

func (set *Set) Contains(a interface{}) bool {
	return (*set)[a]
}

func (set *Set) Size() int {
	return len(*set)
}

type Map map[interface{}]interface{}

func (m *Map) Put(a interface{}, x interface{}) {
	(*m)[a] = x
}

func (m *Map) Remove(a interface{}) {
	delete(*m, a)
}

func (m *Map) Get(a interface{}) interface{} {
	return (*m)[a]
}

func (m *Map) ContainsKey(a interface{}) bool {
	_, ok := (*m)[a]
	return ok
}

func (m *Map) GetOrDefault(a interface{}, x interface{}) interface{} {
	if m.ContainsKey(a) {
		return m.Get(a)
	}
	return x
}

func (m *Map) Size() int {
	return len(*m)
}

func (m *Map) KeySet() Set {
	var keys Set
	for key := range *m {
		keys.Add(key)
	}
	return keys
}

func (m *Map) Values() List {
	var values *ArrayList
	for _, value := range *m {
		values.Add(value)
	}
	return values
}
