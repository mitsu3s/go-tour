package functions

func SliceAppend() {
	var s []int
	printSlice(s)

	s = append(s, 0)
	printSlice(s)

	s = append(s, 1, 2, 3, 4)
	printSlice(s)
}
