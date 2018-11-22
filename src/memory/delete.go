package memory

import "time"

func clear() {
	now := time.Now().Unix()
	scan(deadlines, now)
	scan(updates, now)
}

func scan(m map[string]int64, now int64) {
	for unique, time := range m {
		if now > time {
			Delete(unique)
		}
	}
}

// Delete 删除缓存数据。
func Delete(unique string) {
	delete(bytes, unique)
	delete(updates, unique)
	delete(deadlines, unique)
}
