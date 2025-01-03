package lecture3

type ArrayCollection struct {
	Items []int
}

func (ac *ArrayCollection) BinarySearch(item int) (int, int) {
	left := 0
	right := len(ac.Items) - 1
	count := 0
	//mid := (left+right)/2
	for left <= right {
		count++
		mid := (right + left) / 2

		if ac.Items[mid] == item {
			return mid, count
		} else if ac.Items[mid] < item {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1, count
}
