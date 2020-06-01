package mock

import (
	"net/http"
	"fmt"
)

//解析参数 返回eorderNo
func DecodeMd5Request(request *http.Request) string {
	eorderId := request.PostFormValue("eorderId")
	fmt.Print(eorderId)
	return eorderId
}
