package base

type Heap interface {
	PeekTop() interface{}                      //查看堆顶元素
	Pop() interface{}                          //弹出堆顶元素
	Replace(ele interface{}) (pre interface{}) //替换堆顶元素
	Remove(index int)
	Len() int
	Add(ele ...interface{})
}

type binaryHeap struct {
	elements []interface{}
	cmp      CompareFunc //比较
}

func NewBinaryHeap(cmp CompareFunc) Heap {
	if cmp == nil {
		panic(any("cmp should not nil"))
	}
	return &binaryHeap{
		cmp: cmp,
	}
}

func NewFourHeap(cmp CompareFunc) Heap {
	if cmp == nil {
		panic(any("cmp should not nil"))
	}
	return &fourHeap{
		cmp: cmp,
	}
}
func (h *binaryHeap) Pop() interface{} { //弹出堆顶元素
	ele := h.PeekTop()
	if ele != nil {
		h.Remove(0)
	}
	return ele
}

func (h *binaryHeap) Add(eles ...interface{}) {
	for _, ele := range eles {
		h.elements = append(h.elements, ele)
		h.shiftUp(h.Len() - 1)
	}
}

func (h *binaryHeap) PeekTop() interface{} {
	if h.Len() <= 0 {
		return nil
	}
	return h.elements[0]
}

// 替换堆顶元素
func (h *binaryHeap) Replace(ele interface{}) (pre interface{}) {
	if h.Len() == 0 {
		h.elements[0] = ele
	} else {
		pre = h.elements[0]
		h.elements[0] = ele
		h.shiftDown(0)
	}
	return ele
}

func (h *binaryHeap) shiftDown(index int) {
	n := h.Len()
	if index < 0 || index >= n {
		return
	}
	ele := h.elements[index]
	for { //叶子节点数量
		left := index<<1 + 1
		if left >= n {
			break
		}
		w := h.elements[left]
		if left+1 < n && h.cmp(w, h.elements[left+1]) == E1GenerateE2 {
			w = h.elements[left+1]
			left++
		}
		if h.cmp(w, ele) == E1GenerateE2 {
			break
		}
		h.elements[index] = h.elements[left]
		index = left
	}
	h.elements[index] = ele

}
func (h *binaryHeap) shiftUp(index int) {
	if index < 0 || index >= h.Len() {
		return
	}
	ele := h.elements[index]
	for index > 0 {
		par := (index - 1) >> 1
		if h.cmp(ele, h.elements[par]) == E1GenerateE2 {
			break
		}
		h.elements[index] = h.elements[par]
		index = par
	}
	h.elements[index] = ele

}
func (h *binaryHeap) Remove(index int) {
	if index < 0 || index >= h.Len() {
		return
	}
	h.elements[index] = h.elements[h.Len()-1]
	h.elements[index] = -1
	h.shiftUp(index)

}
func (h *binaryHeap) Len() int {
	return len(h.elements)
}

type fourHeap struct {
	cmp      CompareFunc //比较
	elements []any
}

func (h *fourHeap) PeekTop() interface{} {
	if h.Len() <= 0 {
		return nil
	}
	return h.elements[0]
}
func (h *fourHeap) Replace(ele interface{}) (pre interface{}) {
	if h.Len() == 0 {
		h.elements[0] = ele
	} else {
		pre = h.elements[0]
		h.elements[0] = ele
		h.shiftDown(0)
	}
	return ele
}

// 删除某个元素
func (h *fourHeap) Remove(index int) {
	if index < 0 || index >= h.Len() {
		return
	}
	h.elements[index] = h.elements[h.Len()-1]
	h.elements = h.elements[:h.Len()-1]
	h.shiftDown(index)
}

func (h *fourHeap) Len() int {
	return len(h.elements)
}

func (h *fourHeap) Add(eles ...interface{}) {
	for _, ele := range eles {
		h.elements = append(h.elements, ele)
		h.shiftUp(len(h.elements) - 1)
	}
}
func (h *fourHeap) shiftDown(index int) {
	n := h.Len()
	if index < 0 || index >= n {
		return
	}
	ele := h.elements[index]
	for { //叶子节点数量
		left := index<<2 + 1
		mid := left + 2
		if left >= n {
			break
		}
		w := h.elements[left]

		if left+1 < n && h.cmp(h.elements[left+1], w) == E1LessE2 {
			w = h.elements[left+1]
			left++
		}

		if mid < n {
			w3 := h.elements[mid]

			if mid+1 < n && h.cmp(h.elements[mid+1], w3) == E1LessE2 {
				w3 = h.elements[mid+1]
				mid++
			}

			if h.cmp(w3, w) == E1LessE2 {
				left = mid
				w = w3
			}
		}

		if h.cmp(w, ele) == E1GenerateE2 {
			break
		}
		h.elements[index] = h.elements[left]
		index = left

	}
	h.elements[index] = ele
}

func (h *fourHeap) Pop() interface{} { //弹出堆顶元素
	ele := h.PeekTop()
	if ele != nil {
		h.Remove(0)
	}
	return ele
}
func (h *fourHeap) shiftUp(index int) {
	if index < 0 || index >= h.Len() {
		return
	}
	ele := h.elements[index]

	for index > 0 {
		par := (index - 1) >> 2
		if h.cmp(ele, h.elements[par]) == E1GenerateE2 {
			break
		}
		h.elements[index] = h.elements[par]
		index = par
	}
	h.elements[index] = ele
}
