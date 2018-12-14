package speech

import (
	"encoding/json"
	"httpserv"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
	"util"
)

type info struct {
	Create int64  `json:"create"`
	Modify int64  `json:"modify"`
	Unique string `json:"unique"`
}

func scan() {
	go func() {
		for {
			time.Sleep(time.Minute)
			timeout := time.Now().Unix() - cfg.nTimeout
			overdue := timeout - cfg.nTimeout
			if infos, err := ioutil.ReadDir(root); err == nil {
				for _, info := range infos {
					auth := info.Name()
					time := modifyTime(auth)
					if time > timeout {
						setOutline(auth)
					} else if time > overdue {
						finish(auth)
					}
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
	setOutline(auth)
}

func setOutline(auth string) {
	if data, err := json.Marshal(info{createTime(auth), modifyTime(auth), tail(auth)}); err == nil {
		ioutil.WriteFile(getOutline(auth), data, 0644)
	}
}

func tail(auth string) string {
	if data, err := util.Tail(getUniques(auth), 256); err == nil {
		str := string(data)
		start := strings.LastIndex(str, ":")
		if start == -1 {
			return ""
		}

		end := strings.LastIndex(str, "\n")
		if end == -1 {
			end = len(str)
		}

		return str[start+1 : end]
	}

	return ""
}

func outline(writer http.ResponseWriter, request *http.Request) int {
	auth := httpserv.GetParam(request, "auth", "")
	if auth == "" {
		return httpserv.Send404(writer)
	}

	return httpserv.ServeFile(writer, request, nil, getOutline(auth))
}
