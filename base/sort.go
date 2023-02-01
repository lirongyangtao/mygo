package base

type sortI interface {
	Len() int
	Swap(i, j int)
	Less(i, j int) int32
}

func checkAndGetSliceLen(s sortI) *lessSwap {
	return &lessSwap{
		Swap:   s.Swap,
		Less:   s.Less,
		Length: s.Len(),
	}
}

type lessSwap struct {
	Less   func(i, j int) int32
	Swap   func(i, j int) //交换
	Length int            //数组长度
}

// 基础类型
type SortInt []int

func (s *SortInt) Len() int {
	return len(*s)
}
func (s *SortInt) Swap(i, j int) {
	data := *s
	data[i], data[j] = data[j], data[i]

}
func (s *SortInt) Less(i, j int) int32 {
	data := *s
	if data[i] < data[j] {
		return E1LessE2
	} else if data[i] > data[j] {
		return E1GenerateE2
	} else {
		return E1EqualE2
	}
}

// ================================冒泡====================================
func BubbleSort(s sortI) {
	lessSwap := checkAndGetSliceLen(s)
	bubbleSort2(lessSwap)
}

func bubbleSort(lessSwap *lessSwap) {
	for i := 0; i < lessSwap.Length; i++ {
		for j := 1; j < lessSwap.Length-i; j++ {
			if lessSwap.Less(j-1, j) == E1GenerateE2 {
				lessSwap.Swap(j-1, j)
			}
		}
	}
}

// 优化：完全有序就终止
func bubbleSort1(lessSwap *lessSwap) {
	for i := 0; i < lessSwap.Length; i++ {
		sorted := true
		for j := 1; j < lessSwap.Length-i; j++ {
			if lessSwap.Less(j-1, j) == E1GenerateE2 {
				lessSwap.Swap(j-1, j)
				sorted = false
			}
		}
		if sorted {
			break
		}
	}
}

// 优化：后面排序就停止
func bubbleSort2(lessSwap *lessSwap) {
	for i := 0; i < lessSwap.Length; i++ {
		sortedIndex := 1
		end := lessSwap.Length - i
		for j := 1; j < end; j++ {
			if lessSwap.Less(j-1, j) == E1GenerateE2 {
				lessSwap.Swap(j-1, j)
				sortedIndex = j - 1
			}
		}
		end = sortedIndex
	}
}

// ================================选择====================================
func SelectSort(s sortI) {
	lessSwap := checkAndGetSliceLen(s)
	selectSort(lessSwap)
}

// 找出最大值和最后一个交换
func selectSort(lessSwap *lessSwap) {
	for i := 0; i < lessSwap.Length; i++ {
		last := lessSwap.Length - i - 1
		for j := 0; j < last; j++ {
			if lessSwap.Less(j, last) == E1GenerateE2 {
				lessSwap.Swap(j, last)
			}
		}
	}
}

// ================================插入====================================
func InsertSort(s sortI) {
	lessSwap := checkAndGetSliceLen(s)
	insertSort(lessSwap)
}

func insertSort(lessSwap *lessSwap) {
	for i := 1; i < lessSwap.Length; i++ {
		for j := i; j > 0; j-- {
			if lessSwap.Less(j, j-1) == E1GenerateE2 {
				lessSwap.Swap(j, j-1)
			}
		}
	}
}

// 优化：i 往前挪动的时候进行二分搜索
func insertSort1(lessSwap *lessSwap) {
	for i := 1; i < lessSwap.Length; i++ {
		for j := i; j > 0; j-- {
			if lessSwap.Less(j, j-1) == E1GenerateE2 {
				lessSwap.Swap(j, j-1)
			}
		}
	}
}

// ================================堆排序====================================
func HeapSort(s sortI) {
	lessSwap := checkAndGetSliceLen(s)
	heapSort(lessSwap, 0, lessSwap.Length)
}

// 找出最大值和最后一个交换
func heapSort(lessSwap *lessSwap, a, b int) {
	first := a
	lo := 0
	hi := b - a
	for i := (hi - 1) >> 1; i >= 0; i-- {
		binaryShiftDown(lessSwap, i, hi, first)
	}

	for i := hi - 1; i >= 0; i-- {
		lessSwap.Swap(first, first+i)
		binaryShiftDown(lessSwap, lo, i, first)
	}
}

//================================归并排序 ====================================

//================================快速排序====================================

// ================================希尔排序====================================
