package memory

// Put 放入缓存区。
func Put(unique string, data []byte) {
	bytes[unique] = data
	update(unique)
}
