package util

// Merge 合并byte数组。
func Merge(length int, bytes ...[]byte) []byte {
	if length <= 0 {
		length = 0
		for _, array := range bytes {
			length += len(array)
		}
	}

	array := make([]byte, length)
	var i int
	for _, arr := range bytes {
		for _, b := range arr {
			array[i] = b
			i++
		}
	}

	return array
}
