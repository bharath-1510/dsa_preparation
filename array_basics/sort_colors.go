package arraybasics

func sortColors(arr []int) []int {
	l, mid, h := 0, 0, len(arr)-1
	for mid <= h {
		switch arr[mid] {
		case 0:
			temp := arr[mid]
			arr[mid] = arr[l]
			arr[l] = temp
			mid++
			l++
		case 1:
			mid++
		case 2:
			temp := arr[mid]
			arr[mid] = arr[h]
			arr[h] = temp
			h--
		}
	}
	return arr
}
