package mock

import (
	"net/http"
	"io/ioutil"
	"bytes"
	"fmt"
	url2 "net/url"
	"strings"
)

func EdbProxy(req *http.Request,  writer http.ResponseWriter) []byte  {
	// we need to buffer the body if we want to read it here and send it
	// in the request.
	body, err := ioutil.ReadAll(req.Body)
	println("body", string(body))


	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}

	// you can reassign the body if you need to parse it as multipart
	req.Body = ioutil.NopCloser(bytes.NewReader(body))

	// create a new url from the raw RequestURI sent by the client
	url := "http://vip3013.edb09.net/rest/index.aspx"


	form := make(url2.Values)
	var sign []string
	for h, val := range  req.PostForm {
		fmt.Println(h, val)
		if h != "sign" {
			form[h] = val
		}else {
			sign = val
		}
	}
	form["sign"] = sign

	for h, val := range form {
		fmt.Println("proxy", h, val)
	}

	proxyReq, err := http.NewRequest(req.Method, url, strings.NewReader(form.Encode()))

	// We may want to filter some headers, otherwise we could just use a shallow copy
	// proxyReq.Header = req.Header
	proxyReq.Header = make(http.Header)
	for h, val := range req.Header {
		fmt.Println(h, val)
		proxyReq.Header[h] = val
	}

	httpClient := http.Client{}
	resp, err := httpClient.Do(proxyReq)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadGateway)
	}
	defer resp.Body.Close()
	buf, _ := ioutil.ReadAll(resp.Body)
	return buf
	// legacy code
}
