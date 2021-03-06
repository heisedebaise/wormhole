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
	Finish bool         `json:"finish"`
}

type typeStruct struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

func scan() {
	go func() {
		for {
			time.Sleep(time.Second)
			timeout := time.Now().Unix() - cfg.nTimeout
			overdue := timeout - cfg.nTimeout
			if infos, err := ioutil.ReadDir(root); err == nil {
				for _, info := range infos {
					auth := info.Name()
					mTime := modifyTime(auth)
					if mTime > timeout {
						setOutline(auth, false)
					} else if mTime > overdue {
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
	setOutline(auth, true)
}

func setOutline(auth string, finish bool) {
	file, err := os.Open(getUniques(auth))
	if err != nil {
		return
	}

	defer file.Close()
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

	outline := getOutline(auth)
	if data, err := json.Marshal(outlineStruct{Create: createTime(auth), Modify: modifyTime(auth), Unique: unique, Types: ts, Finish: finish || finished(outline)}); err == nil {
		ioutil.WriteFile(outline, data, 0644)
	}
}

func finished(outline string) bool {
	if data, err := ioutil.ReadFile(outline); err == nil {
		var ol outlineStruct
		if json.Unmarshal(data, &ol) == nil {
			return ol.Finish
		}
	}

	return false
}
