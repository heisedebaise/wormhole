package httpserv

import (
	"encoding/json"
	"log"
	"net/http"
)

// GetParam 获取请求参数。
func GetParam(request *http.Request, name string, defaultValue string) string {
	if _, has := request.Form[name]; has {
		return request.Form[name][0]
	}

	return defaultValue
}

// GetJSON 获取JSON参数。
func GetJSON(request *http.Request, parameter interface{}) error {
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(parameter); err != nil {
		log.Printf("decode http/s body as json failure %q\n", err)

		return err
	}

	return nil
}
