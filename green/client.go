package greensdk

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

type DefaultClient struct {
	Profile
}

func (d *DefaultClient) SetAK(a, k string) {
	d.AccessKeyId = a
	d.AccessKeySecret = k
}

func (d *DefaultClient) ToString(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		logrus.Errorln(err)
		return ""
	}
	return string(b)
}

func (d *DefaultClient) Request(path string, c *ClientInfo, data interface{}) (res []byte, ok bool) {
	userInfo := d.ToString(c)
	dataInfo := d.ToString(data)

	if userInfo == "" || dataInfo == "" {
		logrus.Errorfp("", c, data)
		return []byte("You request is error."), false
	}

	client := &http.Client{}
	URL := host + path + "?clientInfo=" + url.QueryEscape(userInfo)
	req, err := http.NewRequest(method, URL, strings.NewReader(dataInfo))
	if err != nil {
		logrus.Errorfp("", c, data, err)
		return []byte(err.Error()), false
	}

	addRequestHeader(req, path, userInfo, dataInfo, d.AccessKeyId, d.AccessKeySecret)

	response, err := client.Do(req)
	if err != nil {
		logrus.Errorfp("", c, data, err)
		return []byte(err.Error()), false
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logrus.Errorfp("", c, data, err)
		return []byte(err.Error()), false
	}
	return body, true
}

// APIStatusMessage represents an API status message
type APIStatusMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func apiOk(msg string) *APIStatusMessage {
	return &APIStatusMessage{0, msg}
}

func apiError(msg string) *APIStatusMessage {
	return &APIStatusMessage{-1, msg}
}

func doJSONWrite(w http.ResponseWriter, code int, obj interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	start := time.Now()
	if err := json.NewEncoder(w).Encode(obj); err != nil {
		elapsed := time.Since(start)
		logrus.Errorfp("", elapsed, err.Error(), obj)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AllowMethods(next http.HandlerFunc, methods ...string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, method := range methods {
			if r.Method == method {
				next(w, r)
				return
			}
		}
		doJSONWrite(w, 405, apiError("Method not supported"))
	}
}
