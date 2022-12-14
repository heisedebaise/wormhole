package wormhole

import (
    "sync/atomic"
    "time"
)

var units = []string{"B", "KB", "MB", "GB", "TB"}
var count, request, response int64

func stat() {
	for {
		time.Sleep(time.Minute)
		n1, u1 := flow(request)
		n2, u2 := flow(response)
		atomic.StoreInt64(&request, 0)
		atomic.StoreInt64(&response, 0)
		Log("count=%d;request=%d%s/s;response=%d%s/s", count, n1, units[u1], n2, units[u2])
	}
}

func flow(n int64) (int64, int) {
	n /= 60
	unit := 0
	for n > 1024 {
		n >>= 10
		unit += 1
	}

	return n, unit
}
