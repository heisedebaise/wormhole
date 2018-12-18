package httpserv

import (
	"net/http"
	"util"
)

// Auth IP白名单或签名认证。
func Auth(writer http.ResponseWriter, request *http.Request) int {
	if util.InWhiteList(GetIP(request)) || util.CheckSign(request.Form) {
		return 0
	}

	return SendFailure(writer, Failure{Code: 9995, Message: "IP白名单及签名认证失败！"})
}
