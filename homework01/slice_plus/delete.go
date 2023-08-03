package slicePlus

func SliceRemove[T any](slice *[]T, index int) (removedElement any) {
	// remove the index-th element of slice, return the removed element.
	subSlice1 := (*slice)[:index]
	subSlice2 := (*slice)[index+1:]
	removedElement = (*slice)[index]

	sliceTmp := make([]T, 0, cap(subSlice1)+cap(subSlice2))

	for _, i := range subSlice1 {
		sliceTmp = append(sliceTmp, i)
	}

	for _, i := range subSlice2 {
		sliceTmp = append(sliceTmp, i)
	}

	*slice = sliceTmp

	return
}
