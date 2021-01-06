package greensdk

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"net/http"
	"sort"
	"time"

	uuid "github.com/satori/go.uuid"
)

const (
	host                = "http://green.cn-shanghai.aliyuncs.com"
	method              = "POST"
	newline             = "\n"
	MIME                = "application/json"
	acsSignatureVersion = "1.0"
	acsSignatureMethod  = "HMAC-SHA1"
	acsVersion          = "2018-05-09"
)

func buildRequestHeader(requestBody string, path string, accessKeyId string, accessKeySecret string) map[string]string {
	md5Str := getContentMd5([]byte(requestBody))
	// 时间格式为RFC2616!
	gmtCreate := time.Now().UTC().Format(http.TimeFormat)

	acsHeader := map[string]string{
		"x-acs-signature-method":  acsSignatureMethod,
		"x-acs-signature-nonce":   uuid.NewV4().String(),
		"x-acs-signature-version": acsSignatureVersion,
		"x-acs-version":           acsVersion,
	}
	authorization := "acs" + " " + accessKeyId + ":" + signature(acsHeader, md5Str, path, accessKeySecret, gmtCreate)
	headers := map[string]string{
		"Content-Type":  MIME,
		"Accept":        MIME,
		"Date":          gmtCreate,
		"Content-MD5":   md5Str,
		"Authorization": authorization,
	}
	for i, v := range acsHeader {
		headers[i] = v
	}

	return headers
}

func getContentMd5(s []byte) string {
	md5Instance := md5.New()
	md5Instance.Write(s)
	return base64.StdEncoding.EncodeToString(md5Instance.Sum(nil))
}

func signature(acsHeader map[string]string, md5Str string, path string, accessKeySecret string, gmtCreate string) string {
	b := bytes.Buffer{}

	b.WriteString(method)
	b.WriteString(newline)

	b.WriteString(MIME)
	b.WriteString(newline)

	b.WriteString(md5Str)
	b.WriteString(newline)

	b.WriteString(MIME)
	b.WriteString(newline)

	b.WriteString(gmtCreate)
	b.WriteString(newline)

	// 对acsHeader的key排序
	acsKeys := make([]string, 0, len(acsHeader))
	for k := range acsHeader {
		acsKeys = append(acsKeys, k)
	}
	sort.Strings(acsKeys)

	for _, v := range acsKeys {
		b.WriteString(v)
		b.WriteString(":")
		b.WriteString(acsHeader[v])
		b.WriteString(newline)
	}

	b.WriteString(path)

	mac := hmac.New(sha1.New, []byte(accessKeySecret))
	mac.Write([]byte(b.String()))

	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}
