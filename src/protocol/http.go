package protocol

import (
	"fmt"
	"net/http"
)

func Http(host string, path string, handler func(writer http.ResponseWriter, request *http.Request)) {
	http.HandleFunc(path, handler)
	fmt.Printf("Listening on %s\n", host)
	err := http.ListenAndServe(host, nil)
	if err != nil {
		fmt.Println(err)
	}
}
