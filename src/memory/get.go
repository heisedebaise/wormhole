package memory

// Get 获取缓存数据。
func Get(unique string) []byte {
	data, ok := bytes[unique]
	if !ok {
		return nil
	}

	update(unique)

	return data
}
