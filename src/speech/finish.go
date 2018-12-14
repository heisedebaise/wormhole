package speech

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type outlineStruct struct {
	Create int64        `json:"create"`
	Modify int64        `json:"modify"`
	Unique string       `json:"unique"`
	Types  []typeStruct `json:"types"`
}

type typeStruct struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
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
	file, err := os.Open(getUniques(auth))
	if err != nil {
		return
	}

	types := make(map[string]int)
	var unique string
	scanner := bufio.NewScanner(bufio.NewReader(file))
	for scanner.Scan() {
		line := scanner.Text()
		indexOf := strings.Index(line, ":")
		if indexOf == -1 {
			continue
		}

		types[line[:indexOf]]++
		unique = line[indexOf+1:]
	}
	var ts []typeStruct
	for name, count := range types {
		ts = append(ts, typeStruct{name, count})
	}

	if data, err := json.Marshal(outlineStruct{createTime(auth), modifyTime(auth), unique, ts}); err == nil {
		ioutil.WriteFile(getOutline(auth), data, 0644)
	}
}
