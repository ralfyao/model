package model

func Min(arr []int) (min int) {
	min = arr[0]
	for _, v := range arr {
		if v < min {
			v = min
		}
	}
	return
}
