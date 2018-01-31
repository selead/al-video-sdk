package greensdk

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"net/http"
	"time"

	uuid "github.com/spiderorg/al-uuid"
)

const (
	host    string = "http://green.cn-shanghai.aliyuncs.com"
	method  string = "POST"
	MIME    string = "application/json"
	newline string = "\n"
)

func addRequestHeader(req *http.Request, path, userInfo, data, id, secret string) {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(data))
	cipherStr := md5Ctx.Sum(nil)
	base64Md5Str := base64.StdEncoding.EncodeToString(cipherStr)

	keySlice := []string{"x-acs-signature-method"}
	keySlice = append(keySlice, "x-acs-signature-nonce")
	keySlice = append(keySlice, "x-acs-signature-version")
	keySlice = append(keySlice, "x-acs-version")
	valueSlice := []string{"HMAC-SHA1", uuid.Rand().Hex(), "1.0", "2017-01-12"}

	now := time.Now().UTC()
	gmtDate := string([]byte(now.Weekday().String())[0:3]) + now.Format(", 02 Jan 2006 15:04:05 GMT")
	req.Header.Set("Accept", MIME)
	req.Header.Set("Content-Type", MIME)
	req.Header.Set("Content-Md5", base64Md5Str)
	req.Header.Set("Date", gmtDate)
	req.Header.Set(keySlice[0], valueSlice[0])
	req.Header.Set(keySlice[1], valueSlice[1])
	req.Header.Set(keySlice[2], valueSlice[2])
	req.Header.Set(keySlice[3], valueSlice[3])
	signature := singature(keySlice, valueSlice, path, gmtDate, base64Md5Str, userInfo, secret)
	req.Header.Set("Authorization", "acs"+" "+id+":"+signature)
}

func singature(keySlice, valueSlice []string, path, gmtDate, md5Str, userInfo, secret string) string {
	b := bytes.Buffer{}

	b.WriteString(method)
	b.WriteString(newline)

	b.WriteString(MIME)
	b.WriteString(newline)

	b.WriteString(md5Str)
	b.WriteString(newline)

	b.WriteString(MIME)
	b.WriteString(newline)

	b.WriteString(gmtDate)
	b.WriteString(newline)

	for i := 0; i < len(keySlice); i++ {
		b.WriteString(keySlice[i])
		b.WriteString(":")
		b.WriteString(valueSlice[i])
		b.WriteString(newline)
	}

	b.WriteString(path)
	b.WriteString("?clientInfo=")
	b.WriteString(userInfo)

	mac := hmac.New(sha1.New, []byte(secret))
	mac.Write([]byte(b.String()))

	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}
