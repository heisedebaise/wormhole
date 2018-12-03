package memory

import (
	"rsync"
	"util"
)

// Put 放入缓存区。
func Put(unique string, data []byte, deadline int64) {
	bytes[unique] = data
	update(unique, false)
	if deadline > 0 {
		deadlines[unique] = deadline
	}

	rsync.SendBytes(rsync.MemoryFlag, unique, util.Merge(8+len(data), util.Int64ToBytes(deadline), data))
}

// PutString 放入缓存区。
func PutString(unique string, data string, deadline int64) {
	Put(unique, []byte(data), deadline)
}

func sync(unique string, message []byte) {
	if len(message) == 8 {
		updates[unique] = util.BytesToInt64(message[:8])

		return
	}

	bytes[unique] = message[8:]
	update(unique, false)
	if deadline := util.BytesToInt64(message[:8]); deadline > 0 {
		deadlines[unique] = deadline
	}
}
