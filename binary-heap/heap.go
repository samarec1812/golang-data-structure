package binary_heap

type BinaryHeap struct {
		arr [] int
}

func (c *BinaryHeap) HeapSize() int {
	return len(c.arr)
}

func (c* BinaryHeap) Push(value int) {
	c.arr = append(c.arr, value)
	i := c.HeapSize() - 1
	parent := (i - 1)/2

	for i > 0 && c.arr[parent] < c.arr[i] {
		temp := c.arr[i]
		c.arr[i] = c.arr[parent]
		c.arr[parent] = temp

		i = parent
		parent = (i - 1) / 2
	}
}
func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func (c *BinaryHeap) GetMax() int{
	result := c.arr[0]
	c.arr[0] = c.arr[c.HeapSize() - 1]
	c.arr = c.arr[:c.HeapSize()-1]
	c.Heapify(0)
	return result
}

func (c *BinaryHeap)Heapify(value int) {
	var (
		leftChild int
		rightChild int
		largestChild int
	)
	for {
		leftChild = 2*value + 1
		rightChild = 2*value + 2
		largestChild = value

		if leftChild < c.HeapSize() && c.arr[leftChild] > c.arr[largestChild] {
			largestChild = leftChild
		}
		if rightChild < c.HeapSize() && c.arr[rightChild] > c.arr[largestChild] {
			largestChild = rightChild
		}
		if largestChild == value {
			break
		}
		temp := c.arr[value]
		c.arr[value] = c.arr[largestChild]
		c.arr[largestChild] = temp
		value = largestChild
	}
}

func (c *BinaryHeap)BuildHeap(values []int) {
	c.arr = values
	// fmt.Println(c.arr)
	for i := c.HeapSize() / 2; i >= 0; i-- {
		c.Heapify(i)
	}
}
//
//func (c *Bina)