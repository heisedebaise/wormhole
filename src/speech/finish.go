package speech

import (
	"httpserv"
	"net/http"
	"time"
)

var produceTimes = make(map[string]int64)

func scan() {
	go func() {
		for {
			time.Sleep(time.Minute)
			timeout := time.Now().Unix() - cfg.nTimeout
			for auth := range consumers {
				if time, ok := produceTimes[auth]; ok && time < timeout {
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

func finishTime(writer http.ResponseWriter, request *http.Request) int {
	auth := httpserv.GetParam(request, "auth", "")
	if auth == "" {
		return httpserv.Send404(writer)
	}

	if _, ok := produceTimes[auth]; ok {
		writer.Write([]byte("-1"))
	}

	return 200
}
