package base

//二分搜索

func BinarySearch(data []int, value int) (index int, ok bool) {
	return binarySearch(data, 0, len(data), value)
}

func binarySearch(data []int, begin, end int, value int) (index int, ok bool) {
	if begin == end {
		return begin, false
	}
	mid := begin + (end-begin)/2
	if data[mid] == value {
		return mid, true
	} else if data[mid] < value {
		index, ok = binarySearch(data, begin, mid, value)
	} else if data[mid] > value {
		index, ok = binarySearch(data, mid+1, end, value)
	}
	return
}

// 插入排序优化的二分搜索，找到第一个大于value的值
func binarySearchForInsertMerge(data []int, begin, end int, value int) (index int) {
	for begin < end {
		mid := begin + (end-begin)/2
		if data[mid] <= value {
			begin = mid + 1
		} else if data[mid] > value {
			end = mid
		}
	}
	return begin
}
