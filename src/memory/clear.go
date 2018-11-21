package memory

import "time"

func clear() {
	now := time.Now().Unix()
	for unique, time := range times {
		if now-time < cfg.Deadline {
			continue
		}

		delete(times, unique)
		delete(bytes, unique)
	}
}
