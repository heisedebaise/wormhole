package memory

// Get 获取缓存数据。
func Get(unique string) []byte {
	if data, ok := bytes[unique]; ok {
		update(unique, true)

		return data
	}

	return nil
}

// GetString 获取缓存数据。
func GetString(unique string) string {
	if data := Get(unique); data != nil {
		return string(data)
	}

	return ""
}
