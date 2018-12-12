package speech

import "time"

var produceTimes = make(map[string]int64)

func auto() {
	go func() {
		for {
			time.Sleep(time.Minute)
			timeout := time.Now().Unix() - cfg.nTimeout
			for auth := range consumers {
				if produceTimes[auth] < timeout {
					finish(auth)
				}
			}
		}
	}()
}

func finish(auth string) {
	for _, conn := range consumers[auth] {
		delete(consumerChans, conn)
	}
	delete(consumers, auth)
	delete(produceTimes, auth)
}
