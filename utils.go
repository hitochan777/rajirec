package rajirec

func AppendAllIfMissing(slice []int, i []int) []int {
	for _, ele := range i {
		slice = AppendIfMissing(slice, ele)
	}
	return slice
}

func AppendIfMissing(slice []int, i int) []int {
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
}
